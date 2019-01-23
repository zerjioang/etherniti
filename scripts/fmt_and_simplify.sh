#!/bin/bash

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Formatting source code with gofmt (and simplification)"

packageName="github.com/zerjioang/methw"

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
