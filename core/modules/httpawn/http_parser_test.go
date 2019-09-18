package httpawn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	example = `GET / HTTP/1.1
Host: localhost:3333
User-Agent: curl/7.58.0
Accept: */*
`
)

func TestResolveHttpMethod(t *testing.T) {
	t.Run("http-method-detector", func(t *testing.T) {
		result, _ := resolveHttpMethod([]byte(example))
		assert.Equal(t, result, GET)
	})
}

func BenchmarkResolveHttpMethod(b *testing.B) {
	b.Run("method-resolver", func(b *testing.B) {
		req := []byte(example)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = resolveHttpMethod(req)
		}
	})
}
