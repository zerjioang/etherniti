package common

import (
	"testing"

	"github.com/zerjioang/go-hpc/thirdparty/echo"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseController(t *testing.T) {
	t.Run("append", func(t *testing.T) {
		result := new(DatabaseController).buildCompositeId("a", "b")
		assert.Equal(t, string(result), "a.b")
	})
	t.Run("instantiate", func(t *testing.T) {
		ctl := new(DatabaseController)
		assert.NotNil(t, ctl)
	})
	t.Run("register-routes", func(t *testing.T) {
		ctl := new(DatabaseController)
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
