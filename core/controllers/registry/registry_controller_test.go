package registry

import (
	"testing"

	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/thirdparty/echo"

	"github.com/stretchr/testify/assert"
)

func TestRegistryController(t *testing.T) {
	t.Run("create-project-controller", func(t *testing.T) {
		pc := NewRegistryController()
		assert.NotNil(t, pc)
	})
	t.Run("create-project-write", func(t *testing.T) {
		pc := NewRegistryController()
		assert.NotNil(t, pc)
		err := pc.Create(common.NewContext(echo.New()))
		assert.Nil(t, err)
	})
}
