#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

# exit script on error
set -e

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

if ! [ -x "$(command -v dep)" ]; then
  echo 'error: dep is not installed.'
  echo 'installing dep...'
  export GOPATH=$HOME/go
  curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
fi

echo "Downloading dependencies using go dep"
$GOPATH/bin/dep ensure -v
echo "all dependencies downloaded"