package auth

import (
	"bytes"
	"testing"

	"github.com/zerjioang/etherniti/core/logger"
)

func BenchmarkResponse(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewLoginResponse("foo-bar")
		}
	})
	b.Run("json", func(b *testing.B) {
		logger.Enabled(false)
		r := NewLoginResponse("foo-bar")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = r.Json()
		}
	})
	b.Run("writer", func(b *testing.B) {
		b.Run("nil", func(b *testing.B) {
			logger.Enabled(false)
			r := NewLoginResponse("foo-bar")
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = r.Writer(nil)
			}
		})
		b.Run("bytes-buffer", func(b *testing.B) {
			logger.Enabled(false)
			r := NewLoginResponse("foo-bar")
			var buf bytes.Buffer
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				buf.Reset()
				_ = r.Writer(&buf)
			}
		})
	})
}
