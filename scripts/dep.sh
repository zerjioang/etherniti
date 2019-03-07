#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Downloading dependencies using go dep"
$GOPATH/bin/dep ensure -v
echo "all dependencies downloaded"