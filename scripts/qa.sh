#!/bin/bash

#
# Copyright gaethway
# SPDX-License-Identifier: Apache License 2.0
#

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Running code QA basic scripts"

./scripts/fmt_and_simplify.sh && \
./scripts/fmt.sh && \
./scripts/goimports.sh && \
./scripts/license_header.sh

echo "qa scripts finished"