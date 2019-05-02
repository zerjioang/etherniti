// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package aeshash_test

import (
	"testing"

	"github.com/zerjioang/etherniti/core/modules/aeshash"
)

func BenchmarkAesHash(b *testing.B) {
	b.Run("test01", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = aeshash.Hash("cheese")
		}
	})
}
