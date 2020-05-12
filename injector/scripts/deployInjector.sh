#!/usr/bin/env bash

set -ue

[ -z "$MESH_NAME" ] && { echo "Need to set the environment variable MESH_NAME"; exit 1; }

if ! command -v jq >/dev/null 2>&1; then
    echo "Please install jq before continue"
    exit 1
fi

if ! command -v openssl >/dev/null 2>&1; then
    echo "Please install openssl before continue"
    exit 1
fi

if ! command -v kubectl >/dev/null 2>&1; then
    echo "Please install kubectl before continue"
    exit 1
fi

kubectl apply -f deploy/inject-ns.yaml
./scripts/gen-cert.sh
./scripts/ca-bundle.sh
kubectl apply -f _output/inject.yaml
echo "waiting for sidecar-secret-inject to start"
kubectl rollout status deployment/sudecar-secret-inject -n sidecar-secret-inject

ACTUAL_APPMESH_NAME=$(kubectl get deployment sidecar-secret-inject -n sidecar-secret-inject -o=jsonpath="{.spec.template.spec.containers[0].env[?(@.name=='SECRET_NAME')].value}")
if [[ "$ACTUAL_SECRET_NAME" = "$SECRET_NAME" ]]; then
    echo "Secret name has been set up"
else
    echo "Mesh name is unexpected. Expect:${SECRET_NAME}, Actual:${ACTUAL_SECRET_NAME}"
    exit 1
fi

echo "The injector is ready"
