// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package solc

import (
	"testing"

	"github.com/zerjioang/etherniti/shared"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

func TestSolcController(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {
		ctl := NewSolcController()
		assert.NotNil(t, ctl)
	})
	t.Run("register-routes", func(t *testing.T) {
		ctl := NewSolcController()
		e := echo.New()
		// create example group
		testGroup := e.Group("", func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				return next(c)
			}
		})
		ctl.RegisterRouters(testGroup)
		assert.NotNil(t, ctl)
	})
	t.Run("version", func(t *testing.T) {
		ctl := NewSolcController()
		assert.NotNil(t, ctl)

		echo.New()
		ctx := shared.AdquireContext(common.NewContext(echo.New()))
		err := ctl.version(ctx)
		assert.Nil(t, err)
	})
}
