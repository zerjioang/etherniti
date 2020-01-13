package project

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

func TestProjectReleaseController(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		t.Run("instantiate", func(t *testing.T) {
			pc := NewProjectReleaseController()
			assert.NotNil(t, pc)
		})
		t.Run("register", func(t *testing.T) {
			pc := NewProjectReleaseController()
			pc.RegisterRouters(echo.New().Group("", nil))
			assert.NotNil(t, pc)
		})
	})
	t.Run("pointer", func(t *testing.T) {
		t.Run("instantiate", func(t *testing.T) {
			pc := NewProjectReleaseControllerPtr()
			assert.NotNil(t, pc)
		})
		t.Run("register", func(t *testing.T) {
			pc := NewProjectReleaseControllerPtr()
			pc.RegisterRouters(echo.New().Group("", nil))
			assert.NotNil(t, pc)
		})
	})
	t.Run("register-routes", func(t *testing.T) {
		ctl := NewProjectReleaseController()
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
