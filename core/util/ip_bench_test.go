package util

import (
	"testing"
)

func BenchmarkIpToUint32(b *testing.B){

	b.Run("convert-bytes", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		val := "1.41.132.176"
		for n := 0; n < b.N; n++ {
			_ = Ip2int(val)
		}
	})
	b.Run("convert-string", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = Ip2int("1.41.132.176")
		}
	})
	b.Run("convert-string-unsafe-inline", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = Ip2int("1.41.132.176")
		}
	})
	b.Run("convert-string-unsafe", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = Ip2int("1.41.132.176")
		}
	})
	b.Run("decode-int-to-string", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = Int2ip(1697219760)
		}
	})
}
