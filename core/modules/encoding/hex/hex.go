// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package hex

import (
	"encoding/hex"
	"unsafe"
)

const (
	hextable = "0123456789abcdef"
)

func ToEthHex(raw []byte) string {
	dst := make([]byte, 2+len(raw)*2)
	dst[0] = 48  // 0
	dst[1] = 120 // x
	for i, v := range raw {
		dst[2+i*2] = hextable[v>>4]
		dst[2+i*2+1] = hextable[v&0x0f]
	}
	return *(*string)(unsafe.Pointer(&dst))
}

func ToHex(raw []byte) string {
	dst := make([]byte, len(raw)*2)
	for i, v := range raw {
		dst[i*2] = hextable[v>>4]
		dst[i*2+1] = hextable[v&0x0f]
	}
	return *(*string)(unsafe.Pointer(&dst))
}

// decode an standard hex string
func FromHex(raw string) ([]byte, error) {
	return hex.DecodeString(raw)
}

// decode ethereum 0x hex string
func FromEthHex(raw string) ([]byte, error) {
	return hex.DecodeString(raw[2:])
}
