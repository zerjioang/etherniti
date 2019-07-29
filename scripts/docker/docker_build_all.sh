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

arm_arch="armv7l" #ARMv7
arm64_arch="armv7l" #ARMv8
x64_arch="x86_64"
