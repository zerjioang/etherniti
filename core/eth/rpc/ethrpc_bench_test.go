// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ethrpc

import (
	"testing"
)

func BenchmarkEth1(b *testing.B) {
	b.Run("eth1", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Eth1().Int64()
		}
	})
	b.Run("eth1-global", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Eth1Int64()
		}
	})
}
