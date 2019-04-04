// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package server

import "testing"

func BenchmarkNewEthernitiContext(b *testing.B) {
	b.Run("instantiate-nil", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewEthernitiContext(nil)
		}
	})
}
