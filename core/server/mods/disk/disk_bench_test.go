// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package disk

import "testing"

func BenchmarkDiskUsage(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = DiskUsage()
		}
	})
	b.Run("instantiate-concurrent", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			go DiskUsage()
		}
	})
	b.Run("is-monitoring", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		disk := DiskUsage()
		for n := 0; n < b.N; n++ {
			_ = disk.IsMonitoring()
		}
	})
	b.Run("read-all", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		disk := DiskUsage()
		disk.Start("/")
		for n := 0; n < b.N; n++ {
			_ = disk.All()
		}
	})

	b.Run("read-all-concurrent", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		disk := DiskUsage()
		disk.Start("/")
		for n := 0; n < b.N; n++ {
			go disk.All()
		}
	})

	b.Run("read-used", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		disk := DiskUsage()
		disk.Start("/")
		for n := 0; n < b.N; n++ {
			_ = disk.Used()
		}
	})
	b.Run("read-free", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		disk := DiskUsage()
		disk.Start("/")
		for n := 0; n < b.N; n++ {
			_ = disk.Free()
		}
	})
}
