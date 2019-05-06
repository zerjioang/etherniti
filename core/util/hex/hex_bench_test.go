// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package hex

import (
	"encoding/hex"
	"testing"
)

func BenchmarkEncode(b *testing.B) {
	b.Run("encode-stdlib", func(b *testing.B) {
		data := []byte("this-is-a-test")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = hex.EncodeToString(data)
		}
	})
	b.Run("encode-fast", func(b *testing.B) {
		data := []byte("this-is-a-test")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = UnsafeEncodeToString(data)
		}
	})
	b.Run("encode-fast-pooled", func(b *testing.B) {
		data := []byte("this-is-a-test")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = UnsafeEncodeToStringPooled(data)
		}
	})
	b.Run("decode-stdlib", func(b *testing.B) {
		data := "746869732d69732d612d74657374"
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _= hex.DecodeString(data)
		}
	})
	b.Run("decode-fast", func(b *testing.B) {
		data := "746869732d69732d612d74657374"
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_,_ = UnsafeDecodeString(data)
		}
	})
}
