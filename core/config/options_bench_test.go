// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

import (
	"testing"

	"github.com/zerjioang/etherniti/core/logger"
)

func BenchmarkCommon(b *testing.B) {
	b.Run("get-defaults", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetDefaultOpts()
		}
	})
	b.Run("BlockTorConnections", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetDefaultOpts().BlockTorConnections
		}
	})
}
