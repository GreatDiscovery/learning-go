package vaultclient

import (
	"context"
	"errors"
	"fmt"
	"strings"

	vault "github.com/hashicorp/vault/api"
)

const defaultMount = "secret"

// Config contains the minimum settings needed by a Vault KV v2 client.
// Address and Token should normally come from VAULT_ADDR and VAULT_TOKEN.
type Config struct {
	Address string
	Token   string
	Mount   string
}

// Client is a small application-facing wrapper around Vault's official Go
// client. It deliberately exposes domain operations instead of the raw client.
type Client struct {
	kv *vault.KVv2
}

// New creates a token-authenticated KV v2 client.
func New(cfg Config) (*Client, error) {
	cfg.Address = strings.TrimSpace(cfg.Address)
	cfg.Token = strings.TrimSpace(cfg.Token)
	cfg.Mount = strings.Trim(cfg.Mount, "/ ")

	if cfg.Address == "" {
		return nil, errors.New("vault address is required")
	}
	if cfg.Token == "" {
		return nil, errors.New("vault token is required")
	}
	if cfg.Mount == "" {
		cfg.Mount = defaultMount
	}

	apiConfig := vault.DefaultConfig()
	if err := apiConfig.ReadEnvironment(); err != nil {
		return nil, fmt.Errorf("read Vault environment configuration: %w", err)
	}
	apiConfig.Address = cfg.Address

	apiClient, err := vault.NewClient(apiConfig)
	if err != nil {
		return nil, fmt.Errorf("create Vault API client: %w", err)
	}
	apiClient.SetToken(cfg.Token)

	return &Client{kv: apiClient.KVv2(cfg.Mount)}, nil
}

// Put creates a secret or writes a new version of an existing secret.
func (c *Client) Put(ctx context.Context, path string, data map[string]interface{}) error {
	if err := validatePath(path); err != nil {
		return err
	}
	if len(data) == 0 {
		return errors.New("secret data is required")
	}

	if _, err := c.kv.Put(ctx, path, data); err != nil {
		return fmt.Errorf("write Vault secret %q: %w", path, err)
	}
	return nil
}

// Get returns the latest version of a secret.
func (c *Client) Get(ctx context.Context, path string) (map[string]interface{}, error) {
	if err := validatePath(path); err != nil {
		return nil, err
	}

	secret, err := c.kv.Get(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("read Vault secret %q: %w", path, err)
	}
	// The official KV v2 client returns a nil Data field without an error when
	// the latest version is soft-deleted. Normalize that response for callers
	// that only consume secret data and do not inspect version metadata.
	if secret.Data == nil {
		return nil, fmt.Errorf("read Vault secret %q: %w", path, vault.ErrSecretNotFound)
	}
	return secret.Data, nil
}

// Delete soft-deletes the latest version of a secret. KV v2 can restore it.
func (c *Client) Delete(ctx context.Context, path string) error {
	if err := validatePath(path); err != nil {
		return err
	}

	if err := c.kv.Delete(ctx, path); err != nil {
		return fmt.Errorf("delete Vault secret %q: %w", path, err)
	}
	return nil
}

func validatePath(path string) error {
	if strings.Trim(path, "/ ") == "" {
		return errors.New("secret path is required")
	}
	return nil
}
