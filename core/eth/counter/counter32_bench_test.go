// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package counter_test

import (
	"testing"

	"github.com/zerjioang/etherniti/core/eth/counter"
)

func BenchmarkCounterPtr(b *testing.B) {

	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = counter.NewCounter32()
		}
	})
	b.Run("add", func(b *testing.B) {
		c := counter.NewCounter32()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Increment()
		}
	})
	b.Run("get", func(b *testing.B) {
		c := counter.NewCounter32()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = c.Get()
		}
	})
	b.Run("set-n", func(b *testing.B) {
		c := counter.NewCounter32()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Set(uint32(n))
		}
	})
	b.Run("set-fix", func(b *testing.B) {
		c := counter.NewCounter32()
		x := uint32(55)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Set(x)
		}
	})
	b.Run("unsafe-bytes", func(b *testing.B) {
		c := counter.NewCounter32()
		c.Increment()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = c.UnsafeBytes()
		}
	})
	b.Run("unsafe-bytes-fixed", func(b *testing.B) {
		c := counter.NewCounter32()
		c.Increment()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = c.UnsafeBytesFixed()
		}
	})
	b.Run("safe-bytes", func(b *testing.B) {
		c := counter.NewCounter32()
		c.Increment()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = c.SafeBytes()
		}
	})
}
