#!/bin/bash

#
# Copyright etherniti
# SPDX-License-Identifier: Apache License 2.0
#

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Syncing go.ethereum with latest version of github.com"

ethPath="vendor/github.com/ethereum/go-ethereum"

if [[ ! -f ${ethPath} ]] ; then
	echo "etherem secp256k1 C files missing...downloading..."
	if [[ ! -d ${ethPath} ]] ; then
		echo "ethereum c files were not previously downloaded"
		mkdir -p $ethPath
		cd $ethPath
		# files are not downloaded
		echo "cloning eth files..."
		git clone https://github.com/ethereum/go-ethereum .
	else
		# files were already downloaded previously
		echo "previously downloaded ethereum c files found. skipping"
	fi
	echo "go.ethereum c files copied"
fi

echo "sync done"