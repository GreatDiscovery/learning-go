package vaultclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	vault "github.com/hashicorp/vault/api"
	"github.com/stretchr/testify/require"
)

func TestClientPutAndGet(t *testing.T) {
	t.Parallel()

	const token = "test-token"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("X-Vault-Token"); got != token {
			t.Errorf("X-Vault-Token = %q, want %q", got, token)
		}
		if got := r.URL.Path; got != "/v1/secret/data/learning/demo" {
			t.Errorf("request path = %q, want %q", got, "/v1/secret/data/learning/demo")
		}

		switch r.Method {
		case http.MethodPut:
			var body struct {
				Data map[string]interface{} `json:"data"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("decode request body: %v", err)
				http.Error(w, "invalid JSON", http.StatusBadRequest)
				return
			}
			if got := body.Data["message"]; got != "hello vault" {
				t.Errorf("message = %#v, want %q", got, "hello vault")
			}
			writeJSON(t, w, map[string]interface{}{
				"data": map[string]interface{}{
					"created_time":  "2026-07-19T00:00:00Z",
					"deletion_time": "",
					"destroyed":     false,
					"version":       1,
				},
			})
		case http.MethodGet:
			writeJSON(t, w, map[string]interface{}{
				"data": map[string]interface{}{
					"data": map[string]interface{}{
						"message": "hello vault",
					},
					"metadata": map[string]interface{}{
						"created_time":  "2026-07-19T00:00:00Z",
						"deletion_time": "",
						"destroyed":     false,
						"version":       1,
					},
				},
			})
		default:
			http.Error(w, "unexpected method", http.StatusMethodNotAllowed)
		}
	}))
	t.Cleanup(server.Close)

	client, err := New(Config{Address: server.URL, Token: token})
	require.NoError(t, err)

	ctx := context.Background()
	require.NoError(t, client.Put(ctx, "learning/demo", map[string]interface{}{
		"message": "hello vault",
	}))

	data, err := client.Get(ctx, "learning/demo")
	require.NoError(t, err)
	require.Equal(t, "hello vault", data["message"])
}

func TestNewValidatesRequiredConfig(t *testing.T) {
	t.Parallel()

	_, err := New(Config{Token: "token"})
	require.EqualError(t, err, "vault address is required")

	_, err = New(Config{Address: "http://127.0.0.1:8200"})
	require.EqualError(t, err, "vault token is required")
}

func TestClientGetSoftDeletedSecret(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(t, w, map[string]interface{}{
			"data": map[string]interface{}{
				"data": nil,
				"metadata": map[string]interface{}{
					"created_time":  "2026-07-19T00:00:00Z",
					"deletion_time": "2026-07-19T00:01:00Z",
					"destroyed":     false,
					"version":       1,
				},
			},
		})
	}))
	t.Cleanup(server.Close)

	client, err := New(Config{Address: server.URL, Token: "test-token"})
	require.NoError(t, err)

	data, err := client.Get(context.Background(), "deleted")
	require.Nil(t, data)
	require.ErrorIs(t, err, vault.ErrSecretNotFound)
}

func writeJSON(t *testing.T, w http.ResponseWriter, data map[string]interface{}) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		t.Errorf("encode response: %v", err)
	}
}
