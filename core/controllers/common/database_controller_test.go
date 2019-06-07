package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDatabaseController(t *testing.T) {
	t.Run("append", func(t *testing.T) {
		result := new(DatabaseController).buildCompositeId("a", "b")
		assert.Equal(t, string(result), "a.b")
	})
}
