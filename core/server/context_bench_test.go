package server

import "testing"

func BenchmarkNewEthernitiContext(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = NewEthernitiContext()
		}
	})
}
