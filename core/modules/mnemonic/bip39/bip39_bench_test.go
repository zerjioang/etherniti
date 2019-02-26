// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package bip39

import (
	"encoding/hex"
	"testing"
)

func BenchmarkBip39(b *testing.B) {
	b.Run("bip39-generate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		entropy, _ := hex.DecodeString("066dca1a2bb7e8a1db2832148ce9933eea0f3ac9548d793112d9a95c9407efad")
		for n := 0; n < b.N; n++ {
			NewMnemonic(entropy)
		}
	})
	b.Run("is-valid", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			IsMnemonicValid("all hour make first leader extend hole alien behind guard gospel lava path output census museum junior mass reopen famous sing advance salt reform")
		}
	})
}
