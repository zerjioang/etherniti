package jsonboost

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	// short string/[]byte sequences, as the difference between these
	// three methods is a constant overhead
	benchmarkString = "0123456789x"
	benchmarkBytes  = []byte("0123456789y")
)

func TestStringToBytes(t *testing.T) {
	t.Run("unsafe", func(t *testing.T) {
		raw := StringToBytes(benchmarkString)
		assert.Equal(t, raw, benchmarkBytes)
	})
	t.Run("safe", func(t *testing.T) {
		raw := []byte(benchmarkString)
		assert.Equal(t, raw, benchmarkBytes)
	})
}
