// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package common

import (
	"math/big"

	"github.com/zerjioang/etherniti/core/eth/fixtures"
)

// HexToHash sets byte representation of s to hash.
// If b is larger than len(h), b will be cropped from the left.
func HexToHash(s string) fixtures.Hash { return BytesToHash(FromHex(s)) }

// BytesToHash sets b to hash.
// If b is larger than len(h), b will be cropped from the left.
func BytesToHash(b []byte) fixtures.Hash {
	var h fixtures.Hash
	h.SetBytes(b)
	return h
}

// BigToHash sets byte representation of b to hash.
// If b is larger than len(h), b will be cropped from the left.
func BigToHash(b *big.Int) fixtures.Hash { return BytesToHash(b.Bytes()) }
