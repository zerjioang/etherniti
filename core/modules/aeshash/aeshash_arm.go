// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package aeshash

import (
	"github.com/zerjioang/etherniti/core/util/str"
	"hash/fnv"
)

// Hash hashes the given string using the algorithm used by Go's hash tables
// God knows what it really is.
func Hash(key string) uint32 {
	h := fnv.New32a()
	h.Write(str.UnsafeBytes(key))
	return h.Sum32()
}