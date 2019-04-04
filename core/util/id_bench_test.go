// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package util

import "testing"

func BenchmarkGenerateUUID(b *testing.B) {
	b.Run("uuid-from-entropy", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GenerateUUIDFromEntropy()
		}
	})
}
