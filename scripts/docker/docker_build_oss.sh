#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

# exit script on error
set -e

cd "$(dirname "$0")"

# move to project root dir from ./scripts/docker to ./
cd ../..

# docker helper files
source ./scripts/docker/docker_helper.sh
source ./scripts/docker/docker_build.sh
source ./scripts/docker/docker_version.sh

# load colored logs
source ./scripts/colors.sh

log "Etherniti Proxy OSS ($BUILD_VERSION, $edition, $mode) docker image"
#our proxy configuration
edition="oss"
mode="pre"
# get current os system architecture
osarch=$(arch)
# download docker image
log "creating docker image with following configuration:"
echo "architecture : $osarch"
echo "edition      : $edition"
echo "version      : $BUILD_VERSION"
echo "mode         : $mode"
echo ""
log "generating docker image etherniti/proxy-oss:$osarch-$BUILD_VERSION"
echo ""

log "starting docker image compilation"
build --build-arg BUILD_MODE=$mode \
    --build-arg BUILD_EDITION=$edition \
    --build-arg BUILD_VERSION=$BUILD_VERSION \
    -t etherniti/proxy-oss:$osarch-$BUILD_VERSION && \
log "docker image compilation finished"

# show message: Press any key to continue
log "uploading to public docker hub"
read -n 1 -s -r -p "Press any key to continue"
echo ""

# retag
log "creating latest docker image from current version ($BUILD_VERSION)"
docker tag etherniti/proxy-oss:$osarch-$BUILD_VERSION etherniti/proxy-oss:$osarch-latest
# push
log "uploading version $BUILD_VERSION to dockerhub"
docker push etherniti/proxy-oss:$osarch-$BUILD_VERSION
log "uploading version 'latest' to dockerhub"
docker push etherniti/proxy-oss:$osarch-latest
ok "Etherniti Proxy OSS ($BUILD_VERSION, $edition, $mode) docker image successfully uploaded"
