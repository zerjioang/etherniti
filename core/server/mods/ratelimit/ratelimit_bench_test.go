package ratelimit

import (
	"net/http"
	"testing"
)

func BenchmarkRatelimit(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = NewRateLimitEngine()
		}
	})
	b.Run("eval-nil", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		limiter := NewRateLimitEngine()
		for n := 0; n < b.N; n++ {
			limiter.Eval(nil)
		}
	})
	b.Run("eval-empty", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		limiter := NewRateLimitEngine()
		h := new(http.Header)
		for n := 0; n < b.N; n++ {
			limiter.Eval(h)
		}
	})
}
