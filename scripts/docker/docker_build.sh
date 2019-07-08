#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

# exit script on error
set -e

function build(){
	if ! existsImage "preloader"; then
		log "preloader docker image not found"
		log "building preloader image"
		docker build -f Dockerfile-preloader -t preloader:latest .
		log "preloader image built"
	else 
		log "using already built preloader docker image"
	fi

	title "Building container amd64 image from Dockerfile-dev"
	docker build -f Dockerfile-dev "$@" .
}

function buildArm(){
	if ! existsImage "preloader"; then
		log "preloader docker image not found"
		log "building preloader image"
		docker build -f Dockerfile-preloader -t preloader:latest .
		log "preloader image built"
	else 
		log "using already built preloader docker image"
	fi

	title "Building container arm image from Dockerfile-dev-arm"
	docker build -f Dockerfile-dev-arm "$@" .
}