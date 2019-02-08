#!/usr/bin/env bash

#
# Copyright etherniti
# SPDX-License-Identifier: Apache License 2.0
#

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Checking code quality with linters..."

# go vet
./scripts/govet.sh -tags dev
# go fmt
./scripts/fmt.sh
# go fmt simplify
./scripts/fmt_and_simplify.sh
# go imports
./scripts/goimports.sh