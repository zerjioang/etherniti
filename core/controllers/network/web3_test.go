// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network_test

import (
	"github.com/zerjioang/etherniti/core/eth/rpc"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/core/controllers/network"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func TestWeb3Controller(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {
		ctl := network.NewWeb3Controller(nil)
		assert.NotNil(t, ctl)
	})
	t.Run("register-routes", func(t *testing.T) {
		ctl := network.NewWeb3Controller(nil)
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
	t.Run("resolve-network-id", func(t *testing.T) {
		t.Run("id-1", func(t *testing.T) {
			name := ethrpc.ResolveNetworkId("1")
			assert.Equal(t, name, "Ethereum Mainnet")
		})
	})
}
