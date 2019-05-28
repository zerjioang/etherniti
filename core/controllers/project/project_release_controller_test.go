package project

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/thirdparty/echo"
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
}
