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

if [[ -z "$FILE" ]]; then
    echo "no compilation filename found. setting default to: etherniti"
    FILE="etherniti"
fi

if [[ -z "$BUILD_MODE" ]]; then
    echo "no BUILD_MODE found.           setting default to: dev"
    BUILD_MODE="dev" # prod
fi

if [[ -z "$HELIX_GOARCH" ]]; then
    echo "no GOARCH found.               setting default to: amd64"
    HELIX_GOARCH="amd64" # arm
fi

if [[ -z "$HELIX_GOOS" ]]; then
    echo "no GOOS found.                 setting default to: linux"
    HELIX_GOOS="linux"
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
    if [[ "$HELIX_GOARCH" = "arm" ]]; then
        echo "compiling for arm..."
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
        CC=/usr/bin/arm-linux-gnueabihf-gcc-7 \
        CGO_ENABLED=1 \
        GOOS=linux \
        GOARCH=arm \
        GOARM=7 \
        go build 
    else
        echo "compiling for $HELIX_GOARCH..."
        if [[ "$BUILD_MODE" = "dev" ]]; then
            echo "compiling development version..."
            echo "Using commit hash '$hash' for current build"
            GOOS=${HELIX_GOOS} \
            GOARCH=${HELIX_GOARCH} \
            go build -ldflags "-X 'main.Build=$hash'" -tags dev
        else
            echo "compiling production version..."
            echo "Using commit hash '$hash' for current build"
            CGO_ENABLED=1 \
            GOOS=${HELIX_GOOS} \
            GOARCH=${HELIX_GOARCH} \
            go build \
            -a \
            -tags 'netgo prod' \
            -ldflags "-linkmode external" \
            -ldflags "-extldflags -static" \
            -ldflags "-libgcc=none" \
            -ldflags "-s -w -X 'main.Build=$hash'" \
            -o $FILE && \
            ls -alh
        fi
    fi
}

compile "$@"