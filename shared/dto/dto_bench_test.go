// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package dto_test

import (
	"testing"

	"github.com/zerjioang/etherniti/shared/dto"

	"github.com/zerjioang/go-hpc/util/str"
)

func BenchmarkNewApiError(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = dto.NewApiError(200, "test-stack")
		}
	})
}

func BenchmarkNewApiResponse(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = dto.NewApiResponse(str.UnsafeBytes("success"), nil)
		}
	})
}
