package config

import "testing"

func BenchmarkGetRedirectUrl(b *testing.B) {
	b.Run("redirect", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		// run the Fib function b.N times
		for n := 0; n < b.N; n++ {
			_ = GetRedirectUrl("subdomain.localhost.com", "/v1/do/the/test")
		}
	})
}