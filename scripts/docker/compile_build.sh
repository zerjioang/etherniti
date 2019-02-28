#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

cd "$(dirname "$0")"

bash ./link.sh

# move to project root dir from ./scripts to ./
cd ../..

# default configuration
hash=$(git rev-parse --short HEAD)
if [[ -z "$hash" ]]; then
    echo "no hash found. setting default"
    hash="development-build"
fi

if [[ -z "$BUILD_MODE" ]]; then
    echo "no BUILD_MODE found.           setting default to: dev"
    BUILD_MODE="pre" # dev, pre, prod
fi

if [[ -z "$ETHERNITI_GOARCH" ]]; then
    echo "no GOARCH found.               setting default to: amd64"
    ETHERNITI_GOARCH="amd64" # arm
fi

if [[ -z "$ETHERNITI_GOOS" ]]; then
    echo "no GOOS found.                 setting default to: linux"
    ETHERNITI_GOOS="linux"
fi

if [[ -z "$ETHERNITI_COMPILER" ]]; then
    echo "no CC found.                 setting default to: gcc"
    # gcc gc gccgo
    ETHERNITI_COMPILER="gcc"
fi

# Disabling CGO also removes the need for the cross-compile dependencies
# forced a rebuild with -a
# netgo to make sure we use built-in net package and not the system’s one
# -ldflags '-linkmode external'
# -ldflags '-extldflags -static'
# -ldflags '-libgcc=none'
# -w just disables debug letting the file be smaller
# -s
function compile(){
    outputname=$1
    if [[ -z "$outputname" ]]; then
        echo "no compilation filename found. setting default to: etherniti"
        outputname="etherniti"
    fi
    if [[ "$ETHERNITI_GOARCH" = "arm" ]]; then
        echo "compiling for arm..."
        ETHERNITI_COMPILER=/usr/bin/arm-linux-gnueabihf-gcc-7
        # compile for arm-v7
        sudo apt-get install gcc \
        make \
        libc6-armel-cross \
        libc6-dev-armel-cross \
        binutils-arm-linux-gnueabi \
        libncurses5-dev \
        gccgo-7-arm-linux-gnueabihf \
        gcc-arm-linux-gnueabi
        # trigger the compilation
        CC=${ETHERNITI_COMPILER} \
        CGO_ENABLED=1 \
        GOOS=linux \
        GOARCH=arm \
        GOARM=7 \
        go build -o $outputname
    else
        echo "compiling for $ETHERNITI_GOARCH..."
        if [[ "$BUILD_MODE" = "dev" ]]; then
            echo "compiling dev-stage version..."
            echo "Using commit hash '$hash' for current build"
            CGO_ENABLED=1 \
            CC=${ETHERNITI_COMPILER} \
            GOOS=${ETHERNITI_GOOS} \
            GOARCH=${ETHERNITI_GOARCH} \
            go build \
                -ldflags "-X 'main.Build=$hash'" \
                -tags dev \
                -o $outputname
        elif [[ "$BUILD_MODE" = "pre" ]]; then
            echo "compiling pre-stage version..."
            echo "Using commit hash '$hash' for current build"
            CGO_ENABLED=1 \
            CC=${ETHERNITI_COMPILER} \
            GOOS=${ETHERNITI_GOOS} \
            GOARCH=${ETHERNITI_GOARCH} \
            go build \
                -ldflags "-X 'main.Build=$hash'" \
                -tags pre \
                -o $outputname
        else
            echo "compiling production version..."
            echo "Using commit hash '$hash' for current build"
            CGO_ENABLED=1 \
            CC=${ETHERNITI_COMPILER} \
            GOOS=${ETHERNITI_GOOS} \
            GOARCH=${ETHERNITI_GOARCH} \
            go build \
            -a \
            -tags 'netgo prod' \
            -ldflags "-linkmode external" \
            -ldflags "-extldflags -static" \
            -ldflags "-libgcc=none" \
            -ldflags "-s -w -X 'main.Build=$hash'" \
            -o $outputname && \
            ls -alh
        fi
    fi
}

compile "$@"