package api

import "testing"

func BenchmarkBadBot(b *testing.B) {
	b.Run("first", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = BadBotsList[0]
		}
	})
	b.Run("last", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = BadBotsList[len(BadBotsList)-1]
		}
	})
}
