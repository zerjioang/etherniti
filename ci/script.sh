#!/bin/bash

#
# Copyright Etherniti. All Rights Reserved.
# SPDX-License-Identifier: Apache 2
#

# exit script on error
set -e

cd "$(dirname "$0")"

# move to project root dir from ./ci scripts to ./
cd ..

echo "script script"

# check if gocyclo is installed
# STEP 1: calculate cyclomatic complexity

# check if golint is installed or not
if ! [ -x "$(command -v gocyclo)" ]; then
  echo 'error: gocyclo is not installed.'
  echo 'installing gocyclo...'
  go get github.com/fzipp/gocyclo
fi
# run go cyclo
gocyclo main.go
echo "cyclomatic complexity of shared package"
gocyclo shared
echo "cyclomatic complexity of thirdparty package"
gocyclo thirdparty
echo "cyclomatic complexity of core package"
gocyclo core

# STEP 2: run copyfighter
if ! [ -x "$(command -v copyfighter)" ]; then
  echo 'error: copyfighter is not installed.'
  echo 'installing copyfighter...'
  go get -u github.com/jmhodges/copyfighter
fi
# run copyfighter
echo "running copyfighter"
#copyfighter $(pwd)

go get -u github.com/alexkohler/prealloc

# STEP 3: run gosec
if ! [ -x "$(command -v gosec)" ]; then
  echo 'error: gosec is not installed.'
  echo 'installing gosec...'
	# binary will be $GOPATH/bin/gosec
	curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $GOPATH/bin vX.Y.Z
fi
# run gosec and audit full source code
# to scan test files and vendored data use
# gosec -tests -vendor ./...
gosec ./...

# check if golint is installed or not
# STEP 3: GOLINT
if ! [ -x "$(command -v golint)" ]; then
  echo 'error: golint is not installed.'
  echo 'installing golint...'
  go get -u golang.org/x/lint/golint
fi
# run golint
golint -set_exit_status ./...

go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...