#!/usr/bin/env bash

set -e

TARGET_DIR="release"
TARGET_NAME="your-go-project-name"
#PLATFORMS="darwin/amd64 darwin/arm64 linux/386 linux/amd64 linux/arm linux/arm64 windows/amd64"
PLATFORMS="linux/amd64"

rm -rf ${TARGET_DIR}
mkdir ${TARGET_DIR}

for pl in ${PLATFORMS}; do
  export GOOS=$(echo ${pl} | cut -d'/' -f1)
  export GOARCH=$(echo ${pl} | cut -d'/' -f2)
  export CGO_ENABLED=0

  export TARGET=${TARGET_DIR}/${TARGET_NAME}
  if [ "${GOOS}" == "windows" ]; then
    export TARGET=${TARGET_DIR}/${TARGET_NAME}_${GOOS}_${GOARCH}.exe
  else
		export TARGET=${TARGET_DIR}/${TARGET_NAME}_${GOOS}_${GOARCH}
  fi

  echo "build => ${TARGET}"
  go build -trimpath -o ${TARGET}
done
