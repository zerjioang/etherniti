#!/bin/bash

#
# Copyright Etherniti Project. All Rights Reserved.
# SPDX-License-Identifier: GNU GPL v3
#

#source ./colors.sh

# exit script on error
set -e

echo "downloading etherniti/proxy:pro docker image"
docker pull etherniti/proxy:pro

echo "stopping previous proxy containers..."
docker stop api && docker rm api

echo "deploying proxy container"
docker run \
        -d \
        -t \
        -p 80:8080 \
	-p 443:4430 \
        --cpus=4 \
        --hostname apollo-arm \
        --memory=2gb \
        --restart unless-stopped \
        --name api \
        --log-opt max-size=20m \
        --log-opt max-file=5 \
        --log-opt labels=production_status \
        -v $(pwd)/volumes/data:/home/root/.etherniti \
        etherniti/proxy:pro

echo "etherniti proxy container deployed"