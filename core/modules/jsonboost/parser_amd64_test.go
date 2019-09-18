package jsonboost

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssembly(t *testing.T) {
	t.Run("lookup-json-error-code-1", func(t *testing.T) {
		result := Lookup("", "")
		assert.Equal(t, result, "")
	})
	t.Run("lookup-json-error-code-2", func(t *testing.T) {
		result := Lookup("{}", "")
		assert.Equal(t, result, "")
	})
	t.Run("lookup-json-error-3", func(t *testing.T) {
		// this test must return a corrupted json error
		result := Lookup(`{"a": 23`, "a")
		assert.Equal(t, result, "")
	})
	t.Run("lookup-json-extract-value-a", func(t *testing.T) {
		result := Lookup(`{"a": 23}`, "a")
		assert.Equal(t, result, "")
	})
}
