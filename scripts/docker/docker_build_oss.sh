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

# move to project root dir from ./scripts to ./
cd ../..

#our proxy configuration
edition="oss"
version="1.0.0"

# get current os system architecture
osarch=$(arch)
# download docker image
echo "architecture : $osarch"
echo "edition      : $edition"
echo "version      : $version"
echo ""
echo "generating docker image etherniti/proxy-oss:$osarch-$version"
echo ""

build --build-arg BUILD_MODE=pre --build-arg BUILD_EDITION=oss -t etherniti/proxy-oss:$osarch-$version && \
	echo "uploading to public docker hub" && \
	docker push etherniti/proxy-oss:$osarch-$version