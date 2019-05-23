#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

# exit script on error
set -e

cd "$(dirname "$0")"

# move to project root dir from ./ci scripts to ./
cd ..

echo "deployment script"
docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD && \
	docker push etherniti/proxy:$VERSION-oss