package security

import (
	"testing"

	"github.com/zerjioang/etherniti/core/logger"
)

func BenchmarkDomainBlacklist(b *testing.B) {
	b.Run("get-domain-list", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = DomainBlacklist()
		}
	})
	b.Run("get-domain-list-bytes", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = DomainBlacklistBytesData()
		}
	})
}
