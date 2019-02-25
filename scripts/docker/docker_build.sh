#!/bin/bash

#
# Copyright Helix Distributed Ledger. All Rights Reserved.
# SPDX-License-Identifier: GNU GPL v3
#

function build(){
	if ! existsImage "preloader"; then
		echo "preloader docker image not found"
		echo "building preloader image"
		docker build -f Dockerfile-preloader -t preloader:latest .
		echo "preloader image built"
	else 
		echo "using already built preloader docker image"
	fi

	echo "Building container image from Dockerfile"
	docker build -f Dockerfile-dev \
		-t etherniti/proxy:develop \
		$@ .
}