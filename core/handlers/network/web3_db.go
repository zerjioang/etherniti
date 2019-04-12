// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

// eth web3 controller
type Web3DbController struct {
	network *NetworkController
}

// constructor like function
func NewWeb3DbController(network *NetworkController) Web3DbController {
	ctl := Web3DbController{}
	ctl.network = network
	return ctl
}

// END of ERC20 functions

// implemented method from interface RouterRegistrable
func (ctl Web3DbController) RegisterRouters(router *echo.Group) {
}
