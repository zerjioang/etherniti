// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package mtls_test

import (
	"testing"

	"github.com/zerjioang/etherniti/core/listener/mtls"
)

func BenchmarkMtlsListener(b *testing.B) {
	b.Run("instantiation", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			mtls.NewMtlsListener()
		}
	})
}
