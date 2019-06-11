// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func TestNetworkController(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {
		ctl := NewNetworkController()
		assert.NotNil(t, ctl)
	})
	t.Run("register-routes", func(t *testing.T) {
		ctl := NewNetworkController()
		e := echo.New()
		// create example group
		// create example group
		testGroup := e.Group("", func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c *echo.Context) error {
				return next(c)
			}
		})
		ctl.RegisterRouters(testGroup)
		assert.NotNil(t, ctl)
	})
}
