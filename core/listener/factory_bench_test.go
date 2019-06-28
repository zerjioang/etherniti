// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package listener

import (
	"testing"

	"github.com/zerjioang/etherniti/shared/def/listener"
)

func BenchmarkFactoryListener(b *testing.B) {
	b.Run("factory-http", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		//b.SetBytes(int64(unsafe.Sizeof(http.HttpListener{})))
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = FactoryListener(listener.HttpMode)
		}
	})
	b.Run("factory-https", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		//b.SetBytes(int64(unsafe.Sizeof(https.HttpsListener{})))
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = FactoryListener(listener.HttpsMode)
		}
	})
	b.Run("factory-mtls", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		//b.SetBytes(int64(unsafe.Sizeof(mtls.MtlsListener{})))
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = FactoryListener(listener.MTLSMode)
		}
	})
	b.Run("factory-unix", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		//b.SetBytes(int64(unsafe.Sizeof(socket.UnixSocketListener{})))
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = FactoryListener(listener.UnixMode)
		}
	})
	b.Run("factory-other", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		//b.SetBytes(int64(unsafe.Sizeof(socket.UnixSocketListener{})))
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = FactoryListener(listener.UnknownMode)
		}
	})
}
