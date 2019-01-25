#!/bin/bash

#
# Copyright gaethway
# SPDX-License-Identifier: Apache License 2.0
#

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "Downloading dependencies using go dep"
$GOPATH/bin/dep ensure -v

ethTemp="eth-temp"
ethPath="./vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1/include/secp256k1.h"
if [[ ! -f ${ethPath} ]] ; then
	echo "etherem secp256k1 C files missing...downloading..."
	if [[ ! -d ${ethTemp} ]] ; then
		echo "ethereum c files were not previously downloaded"
		# files are not download
		git clone https://github.com/ethereum/go-ethereum $ethTemp
		rm -rf "./vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/"
		cp -ra ${ethTemp}/crypto/secp256k1 ./vendor/github.com/ethereum/go-ethereum/crypto/
	else
		# files were already downloaded previously
		echo "previously downloaded ethereum c files found"
		cp -ra ${ethTemp}/crypto/secp256k1 ./vendor/github.com/ethereum/go-ethereum/crypto/
	fi
	echo "go.ethereum c files copied"
fi

echo "all dependencies downloaded"