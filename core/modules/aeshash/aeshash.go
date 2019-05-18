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

// noescape hides a pointer from escape analysis.  noescape is
// the identity function but escape analysis doesn't think the
// output depends on the input.  noescape is inlined and currently
// compiles down to zero instructions.
// USE CAREFULLY!
//go:nosplit
func noescape(p unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) ^ 0)
}
