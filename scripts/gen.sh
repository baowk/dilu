#!/usr/bin/env bash
set -euo pipefail

if [[ $# -lt 2 ]]; then
  echo "Usage: scripts/gen.sh <db> <table> [module] [config]"
  echo "Example: scripts/gen.sh sys sys_user user resources/config.dev.yaml"
  exit 1
fi

DB_NAME="$1"
TABLE_NAME="$2"
PACKAGE_NAME="${3:-}"
CONFIG_PATH="${4:-resources/config.dev.yaml}"

CMD=(go run main.go gen -c "${CONFIG_PATH}" -d "${DB_NAME}" -t "${TABLE_NAME}" -f false)
if [[ -n "${PACKAGE_NAME}" ]]; then
  CMD+=(-p "${PACKAGE_NAME}")
fi

echo "Running: ${CMD[*]}"
"${CMD[@]}"
