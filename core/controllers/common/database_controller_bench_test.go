package common

import (
	"github.com/zerjioang/etherniti/core/logger"
	"testing"
)
func BenchmarkDatabaseControllerAppend(b *testing.B) {
	b.Run("append", func(b *testing.B) {
		logger.Enabled(false)
		ctl := new(DatabaseController)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ctl.buildCompositeId("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
		}
	})
	b.Run("append-2", func(b *testing.B) {
		logger.Enabled(false)
		ctl := new(DatabaseController)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ctl.buildCompositeId2("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
		}
	})
	b.Run("append-3", func(b *testing.B) {
		logger.Enabled(false)
		ctl := new(DatabaseController)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ctl.buildCompositeId3("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
		}
	})
	b.Run("append-4", func(b *testing.B) {
		logger.Enabled(false)
		ctl := new(DatabaseController)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ctl.buildCompositeId4("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
		}
	})
	b.Run("append-5", func(b *testing.B) {
		logger.Enabled(false)
		ctl := new(DatabaseController)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ctl.buildCompositeId5("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
		}
	})
	b.Run("append-6", func(b *testing.B) {
		logger.Enabled(false)
		ctl := new(DatabaseController)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ctl.buildCompositeId6("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
		}
	})
	b.Run("append-7", func(b *testing.B) {
		logger.Enabled(false)
		l := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		r := []byte("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
		ctl := new(DatabaseController)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ctl.buildCompositeId7(l, r)
		}
	})
	b.Run("append-8", func(b *testing.B) {
		logger.Enabled(false)
		l := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
		r := []byte("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
		ctl := new(DatabaseController)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ctl.buildCompositeId8(l, r)
		}
	})
}
