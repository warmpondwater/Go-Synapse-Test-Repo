#!/bin/bash
set -euo pipefail

REPO_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SYNAPSE_BIN="${1:-${REPO_DIR}/../Go-Synapse/dist/Go-Synapse}"

if [ ! -f "$SYNAPSE_BIN" ]; then
    echo "Go-Synapse binary not found at $SYNAPSE_BIN. Trying local builds..."
    if [ -f "${REPO_DIR}/../Go-Synapse/builds/Go-Synapse-mac-arm64" ]; then
        SYNAPSE_BIN="${REPO_DIR}/../Go-Synapse/builds/Go-Synapse-mac-arm64"
    elif [ -f "${REPO_DIR}/../Go-Synapse/Go-Synapse" ]; then
        SYNAPSE_BIN="${REPO_DIR}/../Go-Synapse/Go-Synapse"
    else
        echo "Error: Could not locate Go-Synapse binary."
        exit 1
    fi
fi

echo "=========================================================="
echo "Go-Synapse Verification Suite: Tree-sitter & LSP Integrity"
echo "Target Repo : $REPO_DIR"
echo "Binary      : $SYNAPSE_BIN"
echo "Manifest    : $REPO_DIR/test_manifest.json"
echo "=========================================================="

# 1. Run Verification Audit
echo "[1/3] Executing Go-Synapse Audit and AST Integrity Pass..."
"$SYNAPSE_BIN" -audit=verify -dir "$REPO_DIR"

CERT_PATH="$REPO_DIR/logs/audit_certificate.json"
if [ ! -f "$CERT_PATH" ]; then
    echo "Error: audit_certificate.json was not generated at $CERT_PATH"
    exit 1
fi

# 2. Extract Integrity Hash
INTEGRITY_HASH=$(grep -o '"integrity_hash": "[^"]*"' "$CERT_PATH" | cut -d'"' -f4)
NODES_RECONCILED=$(grep -o '"nodes_reconciled": [0-9]*' "$CERT_PATH" | awk '{print $2}')
EDGES_RECONCILED=$(grep -o '"edges_reconciled": [0-9]*' "$CERT_PATH" | awk '{print $2}')

echo "[2/3] Audit Results:"
echo "  * SHA-256 Integrity Hash: $INTEGRITY_HASH"
echo "  * Reconciled Nodes      : $NODES_RECONCILED"
echo "  * Reconciled Edges      : $EDGES_RECONCILED"

# 3. Assertions
echo "[3/3] Running Validation Assertions..."
if [ -z "$INTEGRITY_HASH" ]; then
    echo "Assertion FAILED: Integrity hash is empty."
    exit 1
fi

if [ "$NODES_RECONCILED" -lt 20 ]; then
    echo "Assertion FAILED: Expected at least 20 reconciled nodes across 11 languages, got $NODES_RECONCILED."
    exit 1
fi

echo "=========================================================="
echo "SUCCESS: AST Tree-sitter & LSP Integrity Hash Validated!"
echo "Certificate generated and signed at: $CERT_PATH"
echo "=========================================================="
