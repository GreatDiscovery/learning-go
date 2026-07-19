#!/usr/bin/env bash
set -euo pipefail

context="${KUBE_CONTEXT:-kind-vault-auth-lab}"
namespace="vault-lab"
script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

kubectl --context "${context}" apply -f "${script_dir}/vault.yaml"
kubectl --context "${context}" -n "${namespace}" rollout status deployment/vault --timeout=120s

kubectl --context "${context}" -n "${namespace}" exec deployment/vault -- sh -ec '
  export VAULT_ADDR=http://127.0.0.1:8200
  export VAULT_TOKEN=root

  if ! vault auth list -format=json | grep -q '"kubernetes/"'; then
    vault auth enable kubernetes
  fi

  vault write auth/kubernetes/config \
    kubernetes_host="https://${KUBERNETES_SERVICE_HOST}:${KUBERNETES_SERVICE_PORT}"

  vault policy write learning-go /vault/lab/learning-go.hcl

  vault kv put -mount=secret learning-go/config \
    username=demo \
    password=vault-authenticated

  vault kv put -mount=secret other/config \
    username=should-not-be-readable

  vault write auth/kubernetes/role/learning-go \
    bound_service_account_names=learning-go \
    bound_service_account_namespaces=vault-lab \
    audience=vault \
    token_policies=learning-go \
    token_ttl=5m \
    token_max_ttl=15m
'

kubectl --context "${context}" apply -f "${script_dir}/workload.yaml"
kubectl --context "${context}" -n "${namespace}" wait \
  --for=condition=Ready pod/vault-client pod/vault-intruder \
  --timeout=120s

echo "Vault Kubernetes auth lab is ready. Run ./demo.sh to authenticate from the Pod."
