package util

import "testing"

// BenchmarkGenerateUUID/uuid-4         	 2000000	       845 ns/op	   1.18 MB/s	      64 B/op	       2 allocs/op
func BenchmarkGenerateUUID(b *testing.B) {
	b.Run("uuid", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		// run the Fib function b.N times
		for n := 0; n < b.N; n++ {
			_ = GenerateUUID()
		}
	})
}
