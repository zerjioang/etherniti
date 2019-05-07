// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package tokenlist

import (
	"testing"

	"github.com/zerjioang/etherniti/core/logger"
)

func BenchmarkTokenList(b *testing.B) {
	b.Run("get-address", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetTokenAddressByName("$IQN")
		}
	})
	b.Run("get-symbol", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetTokenAddressByName("0x0db8d8b76bc361bacbb72e2c491e06085a97ab31")
		}
	})
}
