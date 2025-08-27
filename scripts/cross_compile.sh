#!/usr/bin/env bash

set -e

DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
cd "$DIR/.."

TARGET_DIR="bin"
#PLATFORMS="darwin/amd64 darwin/arm64 linux/386 linux/amd64 linux/arm linux/arm64 windows/amd64"
PLATFORMS="linux/amd64 windows/amd64"

echo "Cleaning up directory: ${TARGET_DIR}"
rm -rf "${TARGET_DIR}"
echo "Creating directory: ${TARGET_DIR}"
mkdir -p "${TARGET_DIR}"

for cmd_path in cmd/*; do
  if [ -d "$cmd_path" ]; then
    TARGET_NAME=$(basename "$cmd_path")
    PACKAGE_NAME="./${cmd_path}"
    echo "--- Building command: ${TARGET_NAME} ---"

    for pl in ${PLATFORMS}; do
      export GOOS=$(echo "${pl}" | cut -d'/' -f1)
      export GOARCH=$(echo "${pl}" | cut -d'/' -f2)
      export CGO_ENABLED=0

      TARGET=${TARGET_DIR}/${TARGET_NAME}_${GOOS}_${GOARCH}
      if [ "${GOOS}" = "windows" ]; then
        TARGET=${TARGET_DIR}/${TARGET_NAME}_${GOOS}_${GOARCH}.exe
      fi

      echo "build => ${TARGET}"
      go build -trimpath -o "${TARGET}" "${PACKAGE_NAME}"
    done
  fi
done

echo ""
echo "All builds completed."
