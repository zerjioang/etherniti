package project

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectInteractionController(t *testing.T) {
	t.Run("create-project-interaction-controller-struct", func(t *testing.T) {
		p := NewProjectController()
		pc := NewProjectInteractionController(p)
		assert.NotNil(t, pc)
	})
	t.Run("create-project-interaction-controller-ptr", func(t *testing.T) {
		p := NewProjectControllerPtr()
		pc := NewProjectInteractionControllerPtr(p)
		assert.NotNil(t, pc)
	})
}
