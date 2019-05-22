#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

# exit script on error
set -e

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ../..

echo "checking resources status..."

if [[ $HOME == "/home/travis/gopath" ]]; then
    echo "overwriting travis-ci HOME var"
    HOME="/home/travis"
    export HOME="/home/travis"
fi
