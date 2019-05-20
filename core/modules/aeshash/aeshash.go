// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build !arm
// +build amd64

package aeshash

import "unsafe"

func aeshashstr(p unsafe.Pointer, h uintptr) uintptr

// Hash hashes the given string using the algorithm used by Go's hash tables
// God knows what it really is.
func Hash(key string) uint32 {
	return uint32(aeshashstr(unsafe.Pointer(&key), 0))
}
