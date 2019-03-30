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

func ToHex(raw []byte) string {
	dst := make([]byte, len(raw)*2)
	for i, v := range raw {
		dst[i*2] = hextable[v>>4]
		dst[i*2+1] = hextable[v&0x0f]
	}
	return *(*string)(unsafe.Pointer(&dst))
}

func FromHex(raw string) ([]byte, error) {
	return hex.DecodeString(raw)
}
