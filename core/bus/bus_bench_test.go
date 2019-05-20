package bus

import (
	"testing"

	"github.com/zerjioang/etherniti/core/logger"
)

func BenchmarkBus(b *testing.B) {
	b.Run("get-bus", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = SharedBus()
		}
	})
}
