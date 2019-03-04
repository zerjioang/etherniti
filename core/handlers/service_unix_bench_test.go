// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import "testing"

func BenchmarkUnixSocketListener(b *testing.B) {
	b.Run("instantiation", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			NewSocketListener()
		}
	})
}
