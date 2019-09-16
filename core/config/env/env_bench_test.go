// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package env

import (
	"testing"

	"github.com/zerjioang/etherniti/core/logger"
)

func BenchmarkEnvironment(b *testing.B) {
	b.Run("new-env", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = New()
		}
	})
	b.Run("new-env-parallel", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		// RunParallel will create GOMAXPROCS goroutines
		// and distribute work among them.
		b.RunParallel(func(pb *testing.PB) {
			// The loop body is executed b.N times total across all goroutines.
			for pb.Next() {
				_ = New()
			}
		})
	})
	b.Run("read-key-all", func(b *testing.B) {
		logger.Enabled(false)
		cfg := New()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			cfg.Load()
		}
	})
	b.Run("read-key-all-parallel", func(b *testing.B) {
		logger.Enabled(false)
		cfg := New()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		// RunParallel will create GOMAXPROCS goroutines
		// and distribute work among them.
		b.RunParallel(func(pb *testing.PB) {
			// The loop body is executed b.N times total across all goroutines.
			for pb.Next() {
				cfg.Load()
			}
		})
	})
	b.Run("read-key-env", func(b *testing.B) {
		logger.Enabled(false)
		cfg := New()
		cfg.Load()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = cfg.Read("HOME")
		}
	})
	b.Run("read-key-env-parallel", func(b *testing.B) {
		logger.Enabled(false)
		cfg := New()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		// RunParallel will create GOMAXPROCS goroutines
		// and distribute work among them.
		b.RunParallel(func(pb *testing.PB) {
			// The loop body is executed b.N times total across all goroutines.
			for pb.Next() {
				_, _ = cfg.Read("HOME")
			}
		})
	})
}
