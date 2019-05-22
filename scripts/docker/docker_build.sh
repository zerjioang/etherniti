#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

# exit script on error
set -e

function build(){
	if ! existsImage "preloader"; then
		echo "preloader docker image not found"
		echo "building preloader image"
		docker build -f Dockerfile-preloader -t preloader:latest .
		echo "preloader image built"
	else 
		echo "using already built preloader docker image"
	fi

	echo "Building container amd64 image from Dockerfile-dev"
	docker build -f Dockerfile-dev \
		-t etherniti/proxy:pro \
		$@ .
}

function buildArm(){
	if ! existsImage "preloader"; then
		echo "preloader docker image not found"
		echo "building preloader image"
		docker build -f Dockerfile-preloader -t preloader:latest .
		echo "preloader image built"
	else 
		echo "using already built preloader docker image"
	fi

	echo "Building container arm image from Dockerfile-dev-arm"
	docker build -f Dockerfile-dev-arm \
		-t etherniti/proxy:pro \
		$@ .
}