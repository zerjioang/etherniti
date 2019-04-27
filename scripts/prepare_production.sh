#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Preparing code for production..."

# remove logging liness with just text
coreFiles=$(find . -type f -name "*.go" | grep -vendor)
regex="log\.(Debug|Info|Error|Warn|Critical){1}\(\"(\w|\s)*\"\)"
for file in ${coreFiles}
do
	echo "optimizing file: $file"
	sed -ri "s/$regex//g" ${file}
done

# after removing logs, we may need to remove some orphaned imports
./scripts/goimports.sh
./scripts/fmt.sh
./scripts/fmt_and_simplify.sh
./scripts/govet.sh -tags prod