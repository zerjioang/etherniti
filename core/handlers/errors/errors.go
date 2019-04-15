// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package errors

import (
	"errors"

	"github.com/zerjioang/etherniti/core/util/str"
)

const (
	invalidAddress   = `{"message": "please, provide a valid ethereum or quorum address"}`
	accountKeyGenErr = `{"message": "failed to generate ecdsa private key"}`
	noConnErrMsg     = "invalid connection profile key provided in the request header. Please, make sure you have created a connection profile indicating your peer node IP address or domain name."
)

var (
	ErrNoConnectionProfile = errors.New(noConnErrMsg)
	AccountKeyGenErrBytes  = str.UnsafeBytes(accountKeyGenErr)
	InvalidAddressBytes    = str.UnsafeBytes(invalidAddress)
)
