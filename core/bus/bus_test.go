package bus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBus(t *testing.T) {
	t.Run("test-server", func(t *testing.T) {
		b := SharedBus()
		assert.NotNil(t, b)
	})
}
