package auth

import (
	"testing"

	"github.com/zerjioang/etherniti/core/logger"
)

func BenchmarkAuthModel(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {

		}
	})
}
