package util

import "testing"

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
