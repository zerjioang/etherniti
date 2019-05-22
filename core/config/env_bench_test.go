// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

import (
	"testing"

	"github.com/zerjioang/etherniti/core/logger"
)

func BenchmarkGetRedirectUrl(b *testing.B) {
	b.Run("redirect", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetRedirectUrl("subdomain.localhost.com", "/v1/do/the/test")
		}
	})
	b.Run("cert-pem", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetCertPem()
		}
	})
	b.Run("key-pem", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetKeyPem()
		}
	})
}

func BenchmarkGetEnvironment(b *testing.B) {
	b.Run("get-env", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetEnvironment()
		}
	})
	b.Run("get-env-parallel", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		// RunParallel will create GOMAXPROCS goroutines
		// and distribute work among them.
		b.RunParallel(func(pb *testing.PB) {
			// The loop body is executed b.N times total across all goroutines.
			for pb.Next() {
				_ = GetEnvironment()
			}
		})
	})
	b.Run("read-key-env", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetEnvironment().String("X_ETHERNITI_TOKEN_SECRET")
		}
	})
	b.Run("read-key-ptr", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetEnvironment().String("X_ETHERNITI_TOKEN_SECRET")
		}
	})
}
