#!/bin/bash

#
# Copyright etherniti
# SPDX-License-Identifier: Apache License 2.0
#

# exit script on error
set -e

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Running ganache-cli..."

docker run \
	-d \
	--name ganache-cli \
	-p 8545:8545 \
	trufflesuite/ganache-cli:latest \
	-a 10 \
	--debug

echo "ganache-cli running"