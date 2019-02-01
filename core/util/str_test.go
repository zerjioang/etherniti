package util

import "testing"

func TestGetJsonBytes(t *testing.T) {
	t.Run("get-bytes-nil", func(t *testing.T) {
		GetJsonBytes(nil)
	})
}

func BenchmarkGetJsonBytes(b *testing.B) {
	b.Run("get-bytes-nil", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for i := 0; i < b.N; i++ {
			GetJsonBytes(nil)
		}
	})
}
