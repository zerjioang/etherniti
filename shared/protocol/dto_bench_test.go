// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package protocol

import (
	"testing"

	"github.com/zerjioang/etherniti/core/util/str"
)

func BenchmarkNewApiError(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewApiError(200, str.UnsafeBytes("test-stack"))
		}
	})
}

func BenchmarkNewApiResponse(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewApiResponse(str.UnsafeBytes("success"), nil)
		}
	})
}
