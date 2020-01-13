package project

import (
	"testing"

	"github.com/zerjioang/go-hpc/thirdparty/echo"

	"github.com/stretchr/testify/assert"
)

func TestProjectInteractionController(t *testing.T) {
	t.Run("create-project-interaction-controller-ptr", func(t *testing.T) {
		p := NewProjectControllerPtr()
		pc := NewProjectInteractionControllerPtr(p, nil)
		assert.NotNil(t, pc)
	})
	t.Run("register-routes", func(t *testing.T) {
		p := NewProjectControllerPtr()
		ctl := NewProjectInteractionControllerPtr(p, nil)
		e := echo.New()
		// create example group
		// create example group
		testGroup := e.Group("", func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				return next(c)
			}
		})
		ctl.RegisterRouters(testGroup)
		assert.NotNil(t, ctl)
	})
}
