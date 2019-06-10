package project

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectInteractionController(t *testing.T) {
	t.Run("create-project-interaction-controller-ptr", func(t *testing.T) {
		p := NewProjectControllerPtr()
		pc := NewProjectInteractionControllerPtr(p, nil)
		assert.NotNil(t, pc)
	})
}
