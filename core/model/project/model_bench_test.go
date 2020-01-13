package project

import (
	"encoding/json"
	"testing"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/go-hpc/thirdparty/echo/protocol"
	"github.com/zerjioang/go-hpc/thirdparty/echo/protocol/encoding"
)

var (
	testSerializer, _ = encoding.EncodingModeSelector(protocol.ModeJson)
)

func BenchmarkProjectModel(b *testing.B) {
	b.Run("create-model", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			// Generate a seed to determine all keys from.
			// This should be persisted, backed up, and secured
			_ = NewProject("", nil)
		}
	})
	b.Run("serialization-pibench", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			// Generate a seed to determine all keys from.
			// This should be persisted, backed up, and secured
			p := NewProject("", nil)
			_ = p.Value(testSerializer)
		}
	})
	b.Run("deserialization-pibench", func(b *testing.B) {
		p := NewProject("", nil)
		v := p.Value(testSerializer)
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
