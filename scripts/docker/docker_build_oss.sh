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

log "building Etherniti OSS"
#our proxy configuration
edition="oss"
# get current os system architecture
osarch=$(arch)
# download docker image
log "creating docker image with following configuration:"
echo "architecture : $osarch"
echo "edition      : $edition"
echo "version      : $BUILD_VERSION"
echo ""
log "generating docker image etherniti/proxy-oss:$osarch-$BUILD_VERSION"
echo ""

build --build-arg BUILD_MODE=pre --build-arg BUILD_EDITION=oss -t etherniti/proxy-oss:$osarch-$BUILD_VERSION && \

# show message: Press any key to continue
log "uploading to public docker hub"
read -n 1 -s -r -p "Press any key to continue"

docker push etherniti/proxy-oss:$osarch-$BUILD_VERSION
# retag
docker tag etherniti/proxy-oss:$osarch-$version etherniti/proxy-oss:$osarch-latest
docker push etherniti/proxy-oss:$osarch-$version
docker push etherniti/proxy-oss:$osarch-latest
