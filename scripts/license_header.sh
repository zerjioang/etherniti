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

echo "Checking etherniti source files license header..."
pwd

copyrightContent=$(cat ./docs/header.txt)

echo "default header copyright content is:"
echo $copyrightContent

files=$(find . -type f -name "*.go" | grep -vendor)
for f in $files
do
	echo "checking license status of $f"
	if ! grep -q 'Copyright' $f; then
		cat ./docs/header.txt $f > $f.new && mv $f.new $f
	fi
done

