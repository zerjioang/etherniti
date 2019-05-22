#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

# exit script on error
set -e

docker run \
    -it \
    -d \
    --network host \
    -v $(pwd)/volumes/data:/home/etherniti/.etherniti etherniti/proxy:pro
