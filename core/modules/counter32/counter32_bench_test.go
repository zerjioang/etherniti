// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package counter32_test

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
			_ = counter32.NewCounter32()
		}
	})
	b.Run("add", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		c := counter32.NewCounter32()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Increment()
		}
	})
	b.Run("get", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		c := counter32.NewCounter32()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = c.Get()
		}
	})
}
