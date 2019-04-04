// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

import "testing"

func BenchmarkBadBot(b *testing.B) {
	b.Run("first-item-access", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = BadBotsList[0]
		}
	})
	b.Run("last-item-access", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = BadBotsList[len(BadBotsList)-1]
		}
	})
}
