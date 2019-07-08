#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

# exit script on error
set -e

# load colored logs
source ../colors.sh

# read the content of VERSION file and use it as version string
# content example: 1.0.1 
log "reading build version information"
BUILD_VERSION=$(cat ../../VERSION)
ok "build version information: $BUILD_VERSION"