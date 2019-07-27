#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

# exit script on error
set -e

# move to script location
cd "$(dirname "$0")"
# move to project root dir from ./scripts to ./
cd ../..

# source helper files
source ./scripts/colors.sh
source ./scripts/docker/docker_version.sh
source ./scripts/docker/docker_build.sh
source ./scripts/docker/docker_helper.sh

buildArm \
	--build-arg BUILD_MODE=prod \
	--build-arg BUILD_EDITION=pro \
	--build-arg BUILD_VERSION=$BUILD_VERSION \
	--build-arg ETHERNITI_GOARCH=arm \
	-t etherniti/proxy:latest