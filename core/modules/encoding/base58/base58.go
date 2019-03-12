// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package base58

import (
	"github.com/zerjioang/etherniti/core/modules/encoding"
)

const (
	base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

var (
	// Base58 represents bytes as a base-58 number [1-9A-GHJ-LM-Za-z].
	Base58 = encoding.NewEncoding(base58Alphabet)
)
