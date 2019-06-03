// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package http

import "testing"

func BenchmarkHttpListener(b *testing.B) {
	b.Run("instantiation", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewHttpListener()
		}
	})
	b.Run("instantiation-custom", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewHttpListenerCustom()
		}
	})
}
