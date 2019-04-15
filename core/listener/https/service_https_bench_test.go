// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package https_test

import (
	"github.com/zerjioang/etherniti/core/listener/https"
	"testing"
)

func BenchmarkHttpListener(b *testing.B) {
	b.Run("instantiation", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			https.NewHttpsListener()
		}
	})
}
