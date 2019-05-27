package project

import (
	"testing"

	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/thirdparty/echo"

	"github.com/stretchr/testify/assert"
)

func TestProjectController(t *testing.T) {
	t.Run("create-project-controller-struct", func(t *testing.T) {
		pc := NewProjectController()
		assert.NotNil(t, pc)
	})
	t.Run("create-project-controller-ptr", func(t *testing.T) {
		pc := NewProjectControllerPtr()
		assert.NotNil(t, pc)
	})
	t.Run("create-project-write", func(t *testing.T) {
		pc := NewProjectController()
		assert.NotNil(t, pc)
		err := pc.Create(common.NewContext(echo.New()))
		assert.Nil(t, err)
	})
}
