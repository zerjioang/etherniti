#!/bin/bash

#
# Copyright Etherniti Project. All Rights Reserved.
# SPDX-License-Identifier: GNU GPL v3
#

# exit script on error
set -e

#source ./colors.sh

echo "downloading etherniti/proxy:pro docker image"
docker pull etherniti/proxy:pro

echo "stopping previous proxy containers..."
docker stop api && docker rm api

echo "deploying proxy container"
docker run \
        -d \
        -t \
        --network host \
        --cpus=2 \
        --hostname apollo \
        --memory=2gb \
        --restart unless-stopped \
        --name api \
        --log-opt max-size=20m \
        --log-opt max-file=5 \
        --log-opt labels=production_status \
        -v $(pwd)/volumes/data:/home/etherniti/.etherniti \
        etherniti/proxy:pro

echo "etherniti proxy container deployed"
