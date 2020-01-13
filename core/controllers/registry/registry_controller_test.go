package registry

import (
	"testing"

	"github.com/zerjioang/etherniti/shared"

	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/go-hpc/thirdparty/echo"

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
		ctx := shared.AdquireContext(common.NewContext(echo.New()))
		err := pc.Create(ctx)
		assert.Nil(t, err)
		shared.ReleaseContext(ctx)
	})
}
