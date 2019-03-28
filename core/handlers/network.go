// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"github.com/zerjioang/etherniti/core/modules/cache"

	"github.com/labstack/echo"
)

// eth network controller
type NetworkController struct {
	//ethereum interaction cache
	cache *cache.MemoryCache
	//main connection peer address/ip
	peer string
	//connection name: mainet, ropsten, rinkeby, etc
	networkName string
}

// constructor like function
func NewNetworkController() NetworkController {
	ctl := NetworkController{}
	ctl.SetTargetName("eth")
	ctl.cache = cache.NewMemoryCache()
	return ctl
}

func (ctl *NetworkController) SetPeer(peerLocation string) {
	ctl.peer = peerLocation
}

func (ctl *NetworkController) SetTargetName(networkName string) {
	ctl.networkName = networkName
}

// implemented method from interface RouterRegistrable
func (ctl NetworkController) RegisterRouters(router *echo.Group) {
}
