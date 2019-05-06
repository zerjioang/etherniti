// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

import (
	"testing"

	"github.com/zerjioang/etherniti/core/logger"
)

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
	b.Run("get-env-ptr", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetEnvironmentPtr()
		}
	})
	b.Run("get-env-ptr-parallel", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		// RunParallel will create GOMAXPROCS goroutines
		// and distribute work among them.
		b.RunParallel(func(pb *testing.PB) {
			// The loop body is executed b.N times total across all goroutines.
			for pb.Next() {
				_ = GetEnvironmentPtr()
			}
		})
	})
}
