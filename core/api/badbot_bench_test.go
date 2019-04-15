// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

import "testing"

func BenchmarkBadBot(b *testing.B) {
	b.Run("first-item-access", func(b *testing.B) {

		l := BadBotsList

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = l.Contains("almaden")
		}
	})
	b.Run("last-item-access", func(b *testing.B) {

		l := BadBotsList

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = l.Contains("googlebot")
		}
	})
}
