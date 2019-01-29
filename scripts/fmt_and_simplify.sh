#!/bin/bash

#
# Copyright gaethway
# SPDX-License-Identifier: Apache License 2.0
#

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Formatting source code with gofmt (and simplification)"

packageName="github.com/zerjioang/gaethway"

#get all files excluding vendors
filelist=$(find ./ -name "*.go" | grep -vendor)
toreplace="./"
toreplaceBy="/"
for file in ${filelist}
do
	echo "Formatting file $file"
	gofmt -s -w ${file}
done

echo "Code formatting done!"
