package browser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasGraphicInterface(t *testing.T) {
	t.Run("detect-ui", func(t *testing.T) {
		status := detectUI()
		assert.NotNil(t, status)
	})
	t.Run("has-ui", func(t *testing.T) {
		status := HasGraphicInterface()
		assert.NotNil(t, status)
		t.Log("has ui: ", status)
	})
}
