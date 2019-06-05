// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package edition_test

import (
	"testing"

	"github.com/zerjioang/etherniti/core/config/edition"

	"github.com/zerjioang/etherniti/core/logger"
)

func BenchmarkConfigExtra(b *testing.B) {
	b.Run("edition", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			_ = edition.Edition()
		}
	})
	b.Run("edition-parallel", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()

		// RunParallel will create GOMAXPROCS goroutines
		// and distribute work among them.
		b.RunParallel(func(pb *testing.PB) {
			// The loop body is executed b.N times total across all goroutines.
			for pb.Next() {
				_ = edition.Edition()
			}
		})
	})
	b.Run("is-oss", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			_ = edition.IsOpenSource()
		}
	})
	b.Run("is-pro", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			_ = edition.IsEnterprise()
		}
	})
	b.Run("is-valid-edition", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			_ = edition.IsValidEdition()
		}
	})
}
