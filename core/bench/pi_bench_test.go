// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package bench

import (
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/encoding/ioproto"
	"github.com/zerjioang/etherniti/shared/protocol/io"
	"testing"
)

var (
	testSerializer, _ = ioproto.EncodingModeSelector(io.ModeJson)
	scoreVar int64
)

func BenchmarkPi(b *testing.B) {
	b.Run("calculate-score", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			calculateScore()
		}
	})

	b.Run("get-score", func(b *testing.B) {
		calculateScore()
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetScore()
		}
	})
	// we use a local variable to avoid compiler optimizations and to compare benchmark results too
	b.Run("get-score-with-local-variable", func(b *testing.B) {
		var scr int64
		calculateScore()
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			scr = GetScore()
		}
		if scr != 0 {

		}
	})
	// we use a global variable to avoid compiler optimizations and to compare benchmark results too
	b.Run("get-score-with-global-variable", func(b *testing.B) {
		calculateScore()
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			scoreVar = GetScore()
		}
	})
	b.Run("get-bench-time", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = GetBenchTime()
		}
	})
}

