#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

# exit script on error
set -e

cd "$(dirname "$0")"

bash ./link.sh

# move to project root dir from ./scripts to ./
cd ../..

# default configuration
# short version of the hash
hash=$(git rev-parse --short HEAD)

# large version of the hash
# hash=$(git rev-list -1 HEAD)

if [[ -z "$hash" ]]; then
    echo "no hash found. setting default"
    hash="development-build"
fi

if [[ -z "$BUILD_MODE" ]]; then
    echo "no BUILD_MODE found.           setting default to: dev"
    BUILD_MODE="pre" # dev, pre, prod
fi

if [[ -z "$BUILD_EDITION" ]]; then
    echo "no BUILD_EDITION found.           setting default to: oss"
    BUILD_MODE="oss" # oss, pro
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

# defined compilation custom compilation go tags
TAGS="${BUILD_MODE} ${BUILD_EDITION}"

echo "
HASH:                       $hash
BUILD_MODE:                 $BUILD_MODE
BUILD_EDITION:              $BUILD_EDITION
BUILD_VERSION:              $BUILD_VERSION
BUILD_TAGS:                 $TAGS
ETHERNITI_GOARCH:           $ETHERNITI_GOARCH
ETHERNITI_GOOS:             $ETHERNITI_GOOS
ETHERNITI_COMPILER:         $ETHERNITI_COMPILER
ETHERNITI_GOGCCFLAGS:       $ETHERNITI_GOGCCFLAGS
ETHERNITI_CGO_ENABLED:      $ETHERNITI_CGO_ENABLED
DNS RESOLVER:               Pure GO
"

# core=$(uname -m)
# armv7l is 32 bit processor.

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
#
# static compilation issues:
#    https://github.com/golang/go/issues/23265
#
function compile(){

    CGO_ENABLED=${ETHERNITI_CGO_ENABLED}
    CC=${ETHERNITI_COMPILER}
    GOGCCFLAGS=${X_ETHERNITI_GOGCCFLAGS}
    GOOS=${ETHERNITI_GOOS}
    GOARCH=${ETHERNITI_GOARCH}
    GODEBUG=netdns=go

    outputname=$1
    if [[ -z "$outputname" ]]; then
        echo "no compilation filename found. setting default to: etherniti"
        outputname="etherniti"
    fi
    if [[ "$ETHERNITI_GOARCH" = "arm" || "$ETHERNITI_GOARCH" = "arm64" || "$ETHERNITI_GOARCH" = "armv7l" ]]; then
        echo "compiling for arm..."
        echo "Using commit hash '$hash' for current build"

        #ETHERNITI_COMPILER=/usr/bin/arm-linux-gnueabihf-gcc-7
        ETHERNITI_COMPILER=/usr/bin/arm-linux-gnueabi-gcc-6
        # compile for arm-v7
        # trigger the compilation
        env CGO_ENABLED=${ETHERNITI_CGO_ENABLED} \
        CC=${ETHERNITI_COMPILER} \
        GOGCCFLAGS=${X_ETHERNITI_GOGCCFLAGS} \
        GOOS=${ETHERNITI_GOOS} \
        GOARCH=${ETHERNITI_GOARCH} \
        GODEBUG=${GODEBUG} \
        go build \
            -tags "${TAGS}"\
            -ldflags "-s -w -libgcc=none  -X 'main.Commit=${hash}' -X 'main.Edition=${BUILD_EDITION}' -X 'main.Version=${BUILD_VERSION}' -linkmode=external -extldflags -static" \
            -o $outputname
        ls -alh && file $outputname
        # docker run -it --entrypoint=/bin/sh etherniti/proxy-arm:develop
    else
        echo "compiling for ${ETHERNITI_GOARCH}..."
        if [[ "$BUILD_MODE" = "dev" ]]; then
            echo "compiling dev-stage version..."
            echo "Using commit hash '$hash' for current build"
            env CGO_ENABLED=${ETHERNITI_CGO_ENABLED} \
            CC=${ETHERNITI_COMPILER} \
            GOGCCFLAGS=${X_ETHERNITI_GOGCCFLAGS} \
            GOOS=${ETHERNITI_GOOS} \
            GOARCH=${ETHERNITI_GOARCH} \
            GODEBUG=${GODEBUG} \
            go build \
                -tags "${TAGS}"\
                -ldflags "-s -w -X 'main.Commit=${hash}' -X 'main.Edition=${BUILD_EDITION}' -X 'main.Version=${BUILD_VERSION}'" \
                -o $outputname
        elif [[ "$BUILD_MODE" = "pre" ]]; then
            echo "compiling pre-stage version..."
            echo "Using commit hash '$hash' for current build"
            env CGO_ENABLED=${ETHERNITI_CGO_ENABLED} \
            CC=${ETHERNITI_COMPILER} \
            GOGCCFLAGS=${X_ETHERNITI_GOGCCFLAGS} \
            GOOS=${ETHERNITI_GOOS} \
            GOARCH=${ETHERNITI_GOARCH} \
            GODEBUG=${GODEBUG} \
            go build \
                -tags "${TAGS}"\
                -ldflags "-s -w -X 'main.Commit=${hash}' -X 'main.Edition=${BUILD_EDITION}' -X 'main.Version=${BUILD_VERSION}' -linkmode=external -extldflags -static" \
                -o $outputname
        else
            echo "compiling production version..."
            echo "Using commit hash '$hash' for current build"
            env CGO_ENABLED=${ETHERNITI_CGO_ENABLED} \
            CC=${ETHERNITI_COMPILER} \
            GOGCCFLAGS=${X_ETHERNITI_GOGCCFLAGS} \
            GOOS=${ETHERNITI_GOOS} \
            GOARCH=${ETHERNITI_GOARCH} \
            GODEBUG=${GODEBUG} \
            go build \
                -a \
                -tags "netgo ${TAGS}" \
                -ldflags "-s -w -libgcc=none  -X 'main.Commit=${hash}' -X 'main.Edition=${BUILD_EDITION}' -X 'main.Version=${BUILD_VERSION}' -linkmode=external -extldflags -static" \
                -o $outputname && \
            ls -alh
        fi
    fi
}

compile "$@"