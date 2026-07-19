# Vault Kubernetes authentication lab

This lab runs a single-node Kind cluster on Podman and deploys a separate Vault
dev server inside Kubernetes. It demonstrates the complete machine-auth flow:

1. Kubernetes projects a short-lived ServiceAccount JWT into `vault-client`.
2. The Pod submits that JWT to `auth/kubernetes/login` with role `learning-go`.
3. Vault calls Kubernetes' TokenReview API to verify the JWT.
4. The Vault role maps the `vault-lab/learning-go` ServiceAccount to the
   `learning-go` policy.
5. Vault returns a five-minute token that can read only
   `secret/data/learning-go/*`.

The projected JWT uses the explicit audience `vault`; it is not the root token
and it is not stored in an environment variable.

## Run

Create the cluster with Kind's Podman provider:

```bash
KIND_EXPERIMENTAL_PROVIDER=podman kind create cluster \
  --name vault-auth-lab \
  --wait 5m
```

If the Vault image is already available to Podman, load it into Kind to avoid a
second registry pull. Using an image archive is reliable with the experimental
Podman provider:

```bash
podman save --format docker-archive \
  -o /tmp/vault-latest.tar \
  docker.io/hashicorp/vault:latest
KIND_EXPERIMENTAL_PROVIDER=podman kind load image-archive \
  /tmp/vault-latest.tar \
  --name vault-auth-lab
```

Then configure and run the demo:

```bash
chmod +x setup.sh demo.sh
./setup.sh
./demo.sh
```

Useful inspection commands:

```bash
kubectl --context kind-vault-auth-lab -n vault-lab get pods,serviceaccounts
kubectl --context kind-vault-auth-lab -n vault-lab logs deployment/vault
```

Both the ServiceAccount JWT and the Vault token are bearer credentials. The demo
does not print either complete token. Do not log production tokens.

## Cleanup

```bash
KIND_EXPERIMENTAL_PROVIDER=podman kind delete cluster --name vault-auth-lab
```

Vault runs in dev mode with in-memory storage and root token `root`. This is
deliberately insecure and must not be reused as a production deployment.
