// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package fastime

import (
	"testing"
	"time"
)

func BenchmarkFastTime(b *testing.B) {

	b.Run("fastime-now", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Now()
		}
	})
	b.Run("fastime-now-unix", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Now().Unix()
		}
	})
}

func BenchmarkStandardTime(b *testing.B) {

	b.Run("standard-now", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = time.Now()
		}
	})
	b.Run("standard-now-unix", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = time.Now().Unix()
		}
	})
}
