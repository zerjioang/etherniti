#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

cd "$(dirname "$0")"

# source helper files
source ./docker_helper.sh
source ./docker_build.sh

# move to project root dir from ./scripts to ./
cd ../..

build --build-arg BUILD_MODE=dev -t etherniti/proxy:latest