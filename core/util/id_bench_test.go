// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package util

import "testing"

func BenchmarkGenerateUUID(b *testing.B) {
	b.Run("uuid", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = GenerateUUID()
		}
	})
}
