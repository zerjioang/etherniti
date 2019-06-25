// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package protocol_test

import (
	"testing"

	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/core/util/str"
)

func BenchmarkNewApiError(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = protocol.NewApiError(200, str.UnsafeBytes("test-stack"))
		}
	})
}

func BenchmarkNewApiResponse(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = protocol.NewApiResponse(str.UnsafeBytes("success"), nil)
		}
	})
}
