#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

# exit script on error
set -e

cd "$(dirname "$0")"

# move to project root dir from ./ci scripts to ./
cd ..

golint -set_exit_status ./...
go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...