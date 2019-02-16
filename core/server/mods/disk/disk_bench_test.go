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
	b.Run("read-all", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		disk := DiskUsage()
		for n := 0; n < b.N; n++ {
			disk, _ = disk.Eval("/")
			_ = disk.All
		}
	})
	b.Run("read-used", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		disk := DiskUsage()
		for n := 0; n < b.N; n++ {
			disk, _ = disk.Eval("/")
			_ = disk.Used
		}
	})
	b.Run("read-free", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		disk := DiskUsage()
		for n := 0; n < b.N; n++ {
			disk, _ = disk.Eval("/")
			_ = disk.Free
		}
	})
}
