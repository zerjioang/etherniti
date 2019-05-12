package eth

import "testing"

func BenchmarkAddress(b *testing.B) {
	b.Run("convert-address", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ConvertAddress(address0)
		}
	})
	b.Run("is-zero-address", func(b *testing.B) {
		b.Run("invalid-length-address", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = IsZeroAddress("0x-invalid-address")
			}
		})
		b.Run("valid-length-address", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = IsZeroAddress(address0)
			}
		})
	})
	b.Run("is-valid-address", func(b *testing.B) {
		b.Run("invalid-length-address", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = IsValidAddress("0x-invalid-address")
			}
		})
		b.Run("valid-length-address", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = IsValidAddress(address0)
			}
		})
	})
}
