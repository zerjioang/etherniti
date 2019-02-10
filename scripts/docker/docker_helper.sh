#!/bin/bash

#
# Copyright Helix Distributed Ledger. All Rights Reserved.
# SPDX-License-Identifier: GNU GPL v3


# Check whether a given container (filtered by name) exists or not
function existsImage(){
    imageName=${1}
    ids=$(docker images -aq $imageName)
    if [[ ! -z ${ids} ]]; then
        return 0 #true
    else
        return 1 #false
    fi
}

# Check whether a given container (filtered by name) exists or not
function existsContainer(){
	containerName=${1}
	ids=$(docker ps -qa -f name=${containerName})
	if [[ -n ${ids} ]]; then
	    return 0 #true
	else
		return 1 #false
	fi
}

# Check whether a docker network exists or not
function existsNetwork(){
	netName=${1}
	ids=$(docker network ls -q -f name=${netName})
	if [[ -n ${ids} ]]; then
	    return 0 #true
	else
		return 1 #false
	fi
}

function run_api_instance() {
    containerName=${1}
    echo "checking the presence of old peer containers named: ${containerName}"
    if existsContainer ${containerName}; then
        echo "stopping old peer container: ${containerName}"
        docker stop ${containerName}
        echo "removing old peer container: ${containerName}"
        docker rm ${containerName}
    fi

    echo "starting etherniti-peer container with docker"
    echo "running container as: ${X_ETHERNITI_ENV}"

    docker run \
        -d \
        -ti \
        --name ${containerName} \
        --cpus="1" \
        --memory="128m" \
        -e "X_ETHERNITI_ENV=${X_ETHERNITI_ENV}" \
        --log-driver json-file \
        --log-opt mode=non-blocking \
        --log-opt max-buffer-size=4m \
        --log-opt max-file=3 \
        --log-opt max-size=20m \
        etherniti:latest   
}