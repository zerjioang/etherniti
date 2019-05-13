package project

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectController(t *testing.T) {
	t.Run("create-project-controller", func(t *testing.T) {
		pc := NewProjectController()
		assert.NotNil(t, pc)
	})
	t.Run("create-project-write", func(t *testing.T) {
		pc := NewProjectController()
		assert.NotNil(t, pc)
		err := pc.Create(nil)
		assert.Nil(t, err)
	})
}
