#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

docker run -it -p 8080:8080 -v $(pwd)/volumes/data:/home/etherniti/.etherniti etherniti/proxy:develop
