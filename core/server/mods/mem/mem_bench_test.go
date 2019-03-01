// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package mem

import "testing"

func BenchmarkMemStatus(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = MemStatusMonitor()
		}
	})
}
