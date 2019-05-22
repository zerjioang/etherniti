package registry

import (
	"testing"

	"github.com/zerjioang/etherniti/core/logger"
)

func BenchmarkRegistryController(b *testing.B) {
	b.Run("create-controller", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewRegistryController()
		}
	})
}
