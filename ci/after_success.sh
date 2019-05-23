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

echo "after success script"
bash <(curl -s https://codecov.io/bash)