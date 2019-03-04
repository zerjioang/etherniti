// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package lib

import (
	"bytes"
	"strings"
)

func encode(nb uint64, buf *bytes.Buffer, base string) {
	l := uint64(len(base))
	if nb/l != 0 {
		encode(nb/l, buf, base)
	}
	buf.WriteByte(base[nb%l])
}

func decode(enc, base string) uint64 {
	var nb uint64
	lbase := len(base)
	le := len(enc)
	for i := 0; i < le; i++ {
		mult := 1
		for j := 0; j < le-i-1; j++ {
			mult *= lbase
		}
		nb += uint64(strings.IndexByte(base, enc[i]) * mult)
	}
	return nb
}
