package base58

import "testing"

func BenchmarkTrivialBase58Encoding(b *testing.B) {
	initTestPairs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = TrivialBase58Encoding([]byte(testPairs[i].dec))
	}
}

func BenchmarkFastBase58Encoding(b *testing.B) {
	initTestPairs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = FastBase58Encoding(testPairs[i].dec)
	}
}

func BenchmarkTrivialBase58Decoding(b *testing.B) {
	initTestPairs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = TrivialBase58Decoding(testPairs[i].enc)
	}
}

func BenchmarkFastBase58Decoding(b *testing.B) {
	initTestPairs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = FastBase58Decoding(testPairs[i].enc)
	}
}
