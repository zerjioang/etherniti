// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/thirdparty/echo"
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
			return func(c *echo.Context) error {
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
		ctx := common.NewContext(echo.New())
		err := ctl.version(ctx)
		assert.Nil(t, err)
	})
}
