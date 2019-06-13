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

echo "Optimizing source code with tools"


# keyify 	Transforms an unkeyed struct literal into a keyed one.
# rdeps 	Find all reverse dependencies of a set of packages
# staticcheck 	Go static analysis, detecting bugs, performance issues, and much more.
# structlayout 	Displays the layout (field sizes and padding) of structs.
# structlayout-optimize 	Reorders struct fields to minimize the amount of padding.
# structlayout-pretty 	Formats the output of structlayout with ASCII art.

echo "checking if go-tools are installed in GOPATH ($GOPATH)"
if [[ ! -f ${GOPATH}/bin/staticcheck ]]; then
	#statements
	echo "go-tools not found. Downloading via go get"
	go get -u honnef.co/go/tools/cmd/...
else
	echo "go tools installed"
fi

# keyify github.com/zerjioang/etherniti
# staticcheck github.com/zerjioang/etherniti
# structlayout-optimize github.com/zerjioang/etherniti