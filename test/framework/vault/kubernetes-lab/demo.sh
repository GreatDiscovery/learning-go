#!/usr/bin/env bash
set -euo pipefail

context="${KUBE_CONTEXT:-kind-vault-auth-lab}"
namespace="vault-lab"

kubectl --context "${context}" -n "${namespace}" exec vault-client -- sh -ec '
  jwt="$(cat /var/run/secrets/tokens/vault-token)"
  VAULT_TOKEN="$(vault write -field=token auth/kubernetes/login \
    role=learning-go \
    jwt="${jwt}")"
  export VAULT_TOKEN

  echo "=== Vault issued a short-lived token for this Pod ==="
  printf "display_name: "
  vault read -field=display_name auth/token/lookup-self
  echo
  printf "policies: "
  vault read -field=policies auth/token/lookup-self
  echo
  printf "ttl: "
  vault read -field=ttl auth/token/lookup-self
  echo

  echo
  echo "=== Allowed by the learning-go policy ==="
  vault kv get -mount=secret learning-go/config

  echo
  echo "=== Denied: the same token cannot read another path ==="
  if vault kv get -mount=secret other/config; then
    echo "ERROR: forbidden secret was unexpectedly readable" >&2
    exit 1
  fi
  echo "Expected result: Vault returned permission denied."
'

echo
echo "=== Denied identity: a different ServiceAccount cannot use this role ==="
if kubectl --context "${context}" -n "${namespace}" exec vault-intruder -- sh -ec '
  jwt="$(cat /var/run/secrets/tokens/vault-token)"
  vault write -field=token auth/kubernetes/login \
    role=learning-go \
    jwt="${jwt}"
'; then
  echo "ERROR: intruder ServiceAccount unexpectedly authenticated" >&2
  exit 1
fi
echo "Expected result: Vault rejected the intruder ServiceAccount during login."
