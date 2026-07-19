# Vault Go client example

This example follows the official HashiCorp Vault Go client workflow:

1. Create a client with `api.DefaultConfig` and `api.NewClient`.
2. Authenticate the client with `SetToken`.
3. Access versioned secrets with `KVv2(mount).Put/Get/Delete`.

The example expects the local Vault dev server to be available at port 8200.
Never use a root token or dev mode in production.

## Run

Set the connection details used by the `vault-learn` Podman container:

```bash
export VAULT_ADDR=http://127.0.0.1:8200
export VAULT_TOKEN=root
```

Write a secret:

```bash
go run ./test/framework/vault/cmd/vault-client \
  -op put \
  -path learning/demo \
  -message "hello vault"
```

Read it:

```bash
go run ./test/framework/vault/cmd/vault-client \
  -op get \
  -path learning/demo
```

Soft-delete its latest version:

```bash
go run ./test/framework/vault/cmd/vault-client \
  -op delete \
  -path learning/demo
```

Run the isolated unit tests without a real Vault server:

```bash
go test ./test/framework/vault
```

Run the integration test against the local Vault container:

```bash
VAULT_INTEGRATION=1 \
VAULT_ADDR=http://127.0.0.1:8200 \
VAULT_TOKEN=root \
go test ./test/framework/vault -run Integration -v
```

The integration test uses a unique path under `secret/learning-go/integration/`,
verifies put, get, and soft-delete behavior, then removes all test versions and
metadata. Without `VAULT_INTEGRATION=1`, it is skipped.

## Production notes

- Do not hard-code tokens or secrets in source code.
- Use an application auth method such as AppRole, Kubernetes, or a cloud auth
  method instead of a root token.
- Use TLS and validate the Vault server certificate.
- Give the application token only the policy capabilities it needs.

Official references:

- https://developer.hashicorp.com/vault/docs/get-started/developer-qs
- https://developer.hashicorp.com/vault/api-docs/libraries
- https://developer.hashicorp.com/vault/docs/secrets/kv/kv-v2
