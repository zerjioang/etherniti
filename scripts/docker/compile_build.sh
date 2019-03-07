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
    echo "no CC found.                   setting default to: gcc"
    # gcc gc gccgo
    ETHERNITI_COMPILER="gcc"
fi

if [[ -z "$ETHERNITI_GOGCCFLAGS" ]]; then
    defaultFlags="-C -m64 -pthread -lpthread all=-trimpath=$GOPATH -fmessage-length=0"
    echo "no GOGCCFLAGS.                 setting default to: $defaultFlags"
    ETHERNITI_GOGCCFLAGS="$defaultFlags"
fi

if [[ -z "$ETHERNITI_CGO_ENABLED" ]]; then
    echo "no ETHERNITI_CGO_ENABLED.      setting default to: 1"
    ETHERNITI_CGO_ENABLED=1
fi

echo "
HASH:                       $hash
BUILD_MODE:                 $BUILD_MODE
ETHERNITI_GOARCH:           $ETHERNITI_GOARCH
ETHERNITI_GOOS:             $ETHERNITI_GOOS
ETHERNITI_COMPILER:         $ETHERNITI_COMPILER
ETHERNITI_GOGCCFLAGS:       $ETHERNITI_GOGCCFLAGS
ETHERNITI_CGO_ENABLED:      $ETHERNITI_CGO_ENABLED
"

# Disabling CGO also removes the need for the cross-compile dependencies
# forced a rebuild with -a
# netgo to make sure we use built-in net package and not the system’s one
# -trimpath is used to tell go compile to trim $GOPATH from source file path in stack trace.
# -fmessage-length=0 You can use the options described below to control
# the formatting algorithm for diagnostic messages, e.g. how many characters
# per line, how often source location information should be reported. 
# -C is used to disable printing of columns in error messages
# -ldflags '-linkmode external'
# -ldflags '-extldflags -static'
# -ldflags '-libgcc=none'
# libgcc: compiler support lib for internal linking; use "none" to disable
# -w just disables debug letting the file be smaller
# -s
function compile(){

    CGO_ENABLED=${ETHERNITI_CGO_ENABLED}
    CC=${ETHERNITI_COMPILER}
    GOGCCFLAGS=${X_ETHERNITI_GOGCCFLAGS}
    GOOS=${ETHERNITI_GOOS}
    GOARCH=${ETHERNITI_GOARCH}

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
        GOARCH=arm \
        GOARM=7 \
        go build \
            -o $outputname
    else
        echo "compiling for $ETHERNITI_GOARCH..."
        if [[ "$BUILD_MODE" = "dev" ]]; then
            echo "compiling dev-stage version..."
            echo "Using commit hash '$hash' for current build"
            go build \
                -tags dev \
                -ldflags "-X 'main.Build=$hash'" \
                -o $outputname
        elif [[ "$BUILD_MODE" = "pre" ]]; then
            echo "compiling pre-stage version..."
            echo "Using commit hash '$hash' for current build"
            go build \
                -tags pre \
                -ldflags "-X 'main.Build=$hash'" \
                -o $outputname
        else
            echo "compiling production version..."
            echo "Using commit hash '$hash' for current build"
            go build \
            -a \
            -tags 'netgo prod' \
            -ldflags "-s -w -libgcc=none -pthread -lpthread -X 'main.Build=$hash' -linkmode external -extldflags '-static'" \
            -o $outputname && \
            ls -alh
        fi
    fi
}

compile "$@"