#!/bin/bash

set -e

# cd to project root directory
DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
cd "$DIR/.."

# create bin directory if not exists
mkdir -p bin

OS_NAME=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH_NAME=$(uname -m)
# 适配 anme -m 的输出为 go arch 的格式
if [ "$ARCH_NAME" = "x86_64" ]; then
    ARCH_NAME="amd64"
elif [ "$ARCH_NAME" = "aarch64" ]; then
    ARCH_NAME="arm64"
fi

go mod tidy

echo "Building commands for ${OS_NAME}/${ARCH_NAME}"

for cmd_path in cmd/*; do
  if [ -d "$cmd_path" ]; then
    TARGET_NAME=$(basename "$cmd_path")
    PACKAGE_NAME="./${cmd_path}"
    TARGET="bin/${TARGET_NAME}_${OS_NAME}_${ARCH_NAME}"

    echo "Building ${PACKAGE_NAME} => ${TARGET}"
    go build -o "${TARGET}" "${PACKAGE_NAME}"
  fi
done

echo "Build completed."
