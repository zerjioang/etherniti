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

buildArm \
	--build-arg BUILD_MODE=prod \
	--build-arg ETHERNITI_GOARCH=arm \
	--build-arg BUILD_EDITION=pro \
	-t etherniti/proxy:latest