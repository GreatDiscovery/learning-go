package vaultclient

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestIntegrationClientKVv2(t *testing.T) {
	if os.Getenv("VAULT_INTEGRATION") != "1" {
		t.Skip("set VAULT_INTEGRATION=1 to run against a real Vault server")
	}

	mount := os.Getenv("VAULT_MOUNT")
	client, err := New(Config{
		Address: os.Getenv("VAULT_ADDR"),
		Token:   os.Getenv("VAULT_TOKEN"),
		Mount:   mount,
	})
	require.NoError(t, err)

	runID := fmt.Sprintf("%d", time.Now().UnixNano())
	secretPath := "learning-go/integration/" + runID
	want := map[string]interface{}{
		"message": "hello from the Vault integration test",
		"run_id":  runID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// DeleteMetadata removes all versions and metadata, so repeated test runs do
	// not leave secrets behind in the development Vault server.
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cleanupCancel()
		if err := client.kv.DeleteMetadata(cleanupCtx, secretPath); err != nil {
			t.Errorf("clean up integration secret %q: %v", secretPath, err)
		}
	})

	require.NoError(t, client.Put(ctx, secretPath, want))

	got, err := client.Get(ctx, secretPath)
	require.NoError(t, err)
	require.Equal(t, want, got)

	require.NoError(t, client.Delete(ctx, secretPath))
	_, err = client.Get(ctx, secretPath)
	require.Error(t, err, "a soft-deleted KV v2 secret should not be readable")
}
