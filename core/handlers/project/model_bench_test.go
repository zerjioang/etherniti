package project

import (
	"encoding/json"
	"testing"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/util/ip"
)

func BenchmarkProjectModel(b *testing.B) {
	b.Run("create-model", func(b *testing.B) {
		testIp := ip.Ip2intLow("127.0.0.1")
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			// Generate a seed to determine all keys from.
			// This should be persisted, backed up, and secured
			_ = NewProject("", "", testIp)
		}
	})
	b.Run("serialization-bench", func(b *testing.B) {
		testIp := ip.Ip2intLow("127.0.0.1")
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			// Generate a seed to determine all keys from.
			// This should be persisted, backed up, and secured
			p := NewProject("", "", testIp)
			_, _ = p.Storage()
		}
	})
	b.Run("deserialization-bench", func(b *testing.B) {
		testIp := ip.Ip2intLow("127.0.0.1")
		p := NewProject("", "", testIp)
		_, v := p.Storage()
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			// Generate a seed to determine all keys from.
			// This should be persisted, backed up, and secured
			p := NewEmptyProject()
			_ = json.Unmarshal(v, &p)
		}
	})
}
