#!/bin/bash

#
# Copyright etherniti
# SPDX-License-Identifier: Apache License 2.0
#

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Generating test functions base with gotests"

echo "checking if gotests is installed in $GOPATH"
if [[ ! -f ${GOPATH}/bin/gotests ]]; then
	#statements
	echo "gotests not found. Downloading via go get"
	go get -v github.com/cweill/gotests
	cd ${GOPATH}/src/github.com/cweill/gotests/gotests
	go build && go install
fi

if [[ ! -f ${GOPATH}/bin/gotests ]]; then
	echo "failed to install gotests in ${GOPATH}"
	return -1
fi

#get all files excluding vendors
filelist=$(find . -type f -name "*.go" | grep -vendor)
for file in ${filelist}
do
	echo "generating gotests for file $file"
	${GOPATH}/bin/gotests -excl Benchmark.* -w ${file}
done

echo "Code formatting done!"