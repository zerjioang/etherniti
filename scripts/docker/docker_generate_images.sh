#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

# exit script on error
set -e

cd "$(dirname "$0")"

# source helper files
source ./docker_helper.sh
source ./docker_build.sh
source ./docker_version.sh

# move to project root dir from ./scripts to ./
cd ../..

echo "building all docker images"

echo "building oss image for arm"
buildArm \
	--build-arg BUILD_MODE=dev \
	--build-arg BUILD_EDITION=oss \
	--build-arg ETHERNITI_GOARCH=arm \
	-t etherniti/proxy-oss:armv7l-$BUILD_VERSION

echo "building pro image for arm"
buildArm \
	--build-arg BUILD_MODE=prod \
	--build-arg BUILD_EDITION=pro \
	--build-arg ETHERNITI_GOARCH=arm \
	-t etherniti/proxy:armv7l-$BUILD_VERSION

echo "building for x86_64"
echo "building oss image for x86_64"
build \
	--build-arg BUILD_MODE=dev \
	--build-arg BUILD_EDITION=oss \
	--build-arg ETHERNITI_GOARCH=amd64 \
	-t etherniti/proxy-oss:x86_64-$BUILD_VERSION

echo "building pro image for x86_64"
build \
	--build-arg BUILD_MODE=prod \
	--build-arg BUILD_EDITION=pro \
	--build-arg ETHERNITI_GOARCH=amd64 \
	-t etherniti/proxy:x86_64-$BUILD_VERSION

echo "uploading docker etherniti proxy images"

echo "uploading docker etherniti proxy oss images"
docker push etherniti/proxy-oss:armv7l-$BUILD_VERSION
docker push etherniti/proxy-oss:x86_64-$BUILD_VERSION

echo "uploading docker etherniti proxy pro images"
docker push etherniti/proxy:armv7l-$BUILD_VERSION
docker push etherniti/proxy:x86_64-$BUILD_VERSION