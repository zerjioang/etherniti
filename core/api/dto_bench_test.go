package api

import (
	"testing"
)

func BenchmarkNewApiError(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		// run the Fib function b.N times
		for n := 0; n < b.N; n++ {
			_ = NewApiError(200, "test-error")
		}
	})
}

func BenchmarkNewApiResponse(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		// run the Fib function b.N times
		for n := 0; n < b.N; n++ {
			_ = NewApiResponse("success", 12345)
		}
	})
}