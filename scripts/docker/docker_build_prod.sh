#!/bin/bash

#
# Copyright Helix Distributed Ledger. All Rights Reserved.
# SPDX-License-Identifier: GNU GPL v3
#

cd "$(dirname "$0")"

# source helper files
source ./docker_helper.sh

# move to project root dir from ./scripts to ./
cd ../..

if ! existsImage "preloader"; then
	echo "preloader docker image not found"
	echo "building preloader image"
	docker build -f Dockerfile-preloader -t preloader:latest .
	echo "preloader image built"
else 
	echo "using already compiler preloader docker image"
fi

echo "Building production container image from Dockerfile"
docker build -f ./Dockerfile-prod -t etherniti/proxy:latest .