#!/bin/bash

#
# Copyright gaethway
# SPDX-License-Identifier: Apache License 2.0
#

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Downloading dependencies using go dep"
$GOPATH/bin/dep ensure -v
echo "all dependencies downloaded"