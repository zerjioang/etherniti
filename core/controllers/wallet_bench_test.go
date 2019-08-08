package controllers

import (
	"github.com/zerjioang/etherniti/core/eth"
	"github.com/zerjioang/etherniti/core/logger"
	"testing"
)

func BenchmarkWalletController(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {

	})
	b.Run("generate-key", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_,_ = eth.GenerateNewKey()
		}
	})
}
