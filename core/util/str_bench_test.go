// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package util

import (
	"strings"
	"testing"
)

func BenchmarkStringUtils(b *testing.B) {

	b.Run("to-lower-std", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		val := "Hello World, This is AWESOME"
		for n := 0; n < b.N; n++ {
			_ = strings.ToLower(val)
		}
	})
	b.Run("ToLowerAscii", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		val := "Hello World, This is AWESOME"
		for n := 0; n < b.N; n++ {
			_ = ToLowerAscii(val)
		}
	})
}
