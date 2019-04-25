// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

import "testing"

func BenchmarkGetRedirectUrl(b *testing.B) {
	b.Run("redirect", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetRedirectUrl("subdomain.localhost.com", "/v1/do/the/test")
		}
	})
	b.Run("cert-pem", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetCertPem()
		}
	})
	b.Run("key-pem", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetKeyPem()
		}
	})
}
