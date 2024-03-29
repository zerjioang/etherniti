// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ratelimit

import (
	"net/http"
	"testing"
)

func BenchmarkRatelimit(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewRateLimitEngine()
		}
	})
	b.Run("eval-nil", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		limiter := NewRateLimitEngine()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = limiter.Eval("id", nil)
		}
	})
	b.Run("eval-empty", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		limiter := NewRateLimitEngine()
		h := http.Header{}
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = limiter.Eval("127.0.0.1", h)
		}
	})
}
