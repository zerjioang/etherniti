// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package mem

import "testing"

func BenchmarkMemStatus(b *testing.B) {
	b.Run("instantiate-struct", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = MemStatusMonitor()
		}
	})
	b.Run("instantiate-ptr", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = MemStatusMonitorPtr()
		}
	})
	b.Run("instantiate-internal", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = memStatusMonitor()
		}
	})
	b.Run("start", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		m := MemStatusMonitorPtr()
		for n := 0; n < b.N; n++ {
			m.Start()
		}
	})
}
