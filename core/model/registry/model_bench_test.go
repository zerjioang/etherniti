package registry

import (
	"testing"

	"github.com/zerjioang/etherniti/core/logger"
)

func BenchmarkRegistryModel(b *testing.B) {
	b.Run("create-empty", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			// Generate a seed to determine all keys from.
			// This should be persisted, backed up, and secured
			_ = NewEmptyRegistry()
		}
	})
	b.Run("create-with-data", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			// Generate a seed to determine all keys from.
			// This should be persisted, backed up, and secured
			contract := NewEmptyRegistry()

			contract.Name = "test"
			contract.Description = "this is a demo contract"
			contract.Address = "0xf17f52151EbEF6C7334FAD080c5704D77216b732"
			contract.Version = "1.2"
		}
	})
}
