#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

docker run -it --network host -v $(pwd)/volumes/data:/home/etherniti/.etherniti etherniti/proxy:develop
