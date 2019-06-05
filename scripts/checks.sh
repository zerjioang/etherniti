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

# load colored logs
source ./scripts/colors.sh

function install(){
	name=$1
	package=$2
	log "checking if $name is installed in $GOPATH"
	if [[ ! -f ${GOPATH}/bin/$name ]]; then
		#statements
		log "$name not found. Downloading via go get"
		go get -u $package
	fi
}

log "Checking aligment in source code"

# A set of utilities for checking Go sources.
install "aligncheck" "gitlab.com/opennota/check/cmd/aligncheck"
install "structcheck" "gitlab.com/opennota/check/cmd/structcheck"
install "varcheck" "gitlab.com/opennota/check/cmd/varcheck"
# tool to detect Go structs that would take less memory if their fields were sorted.
install "maligned" "github.com/mdempsky/maligned"
# prealloc is a Go static analysis tool to find slice declarations that could potentially be preallocated.
install "prealloc" "github.com/alexkohler/prealloc"

project="github.com/zerjioang/etherniti"
log "aligncheck of ${project}"
${GOPATH}/bin/aligncheck ${project}

#get all files excluding vendors
filelist=$(find . -type f -name "*.go" | grep -vendor)
for file in ${filelist}
do
	log "checking file $file"
	#${GOPATH}/bin/goimports -v -w ${file}
done

ok "Code checks done!"

