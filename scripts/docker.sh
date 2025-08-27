#!/usr/bin/env bash

set -ex

PROFILE=${PROFILE:-"test"}
SRC_PATH=${SRC_PATH:-"."}
BUILD_HOST=${BUILD_HOST:-"unix:///var/run/docker.sock"}
DOCKERFILE=${DOCKERFILE:-"./Dockerfile"}

while getopts "e:o:s:h:f:" OPT; do
    case $OPT in
        e)
            PROFILE=$OPTARG;;
        s)
            SRC_PATH=$OPTARG;;
        h)
            BUILD_HOST=$OPTARG;;
        f)
            DOCKERFILE=$OPTARG;;
        *)
    esac
done

# build
docker -H ${BUILD_HOST} build -t ${IMAGE_NAME} -f ${DOCKERFILE} \
    --build-arg PROFILE="${PROFILE}" \
    --build-arg SRC_PATH="${SRC_PATH}" .

# login
docker -H ${BUILD_HOST} login ${REG_HOST} -u ${REG_USER} -p ${REG_PWD}

# push
docker -H ${BUILD_HOST} tag ${IMAGE_NAME} ${LATEST_IMAGE_NAME}
docker -H ${BUILD_HOST} push ${IMAGE_NAME}
docker -H ${BUILD_HOST} push ${LATEST_IMAGE_NAME}
docker -H ${BUILD_HOST} rmi ${IMAGE_NAME}
docker -H ${BUILD_HOST} rmi ${LATEST_IMAGE_NAME}

echo "Build image success --> ${IMAGE_NAME}"

