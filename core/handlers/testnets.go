// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"github.com/zerjioang/etherniti/core/handlers/network"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

const (
	ropsten       = "ropsten"
	ropstenInfura = "https://ropsten.infura.io/v3/4f61378203ca4da4a6b6601bc16a22ad"

	rinkeby       = "rinkeby"
	rinkebyInfura = "https://rinkeby.infura.io/v3/4f61378203ca4da4a6b6601bc16a22ad"

	kovan       = "kovan"
	kovanInfura = "https://kovan.infura.io/v3/4f61378203ca4da4a6b6601bc16a22ad"

	mainnet       = "mainnet"
	mainnetInfura = "https://mainnet.infura.io/v3/4f61378203ca4da4a6b6601bc16a22ad"

	infura   = "infura"
	quiknode = "quiknode"
	private  = "private"
)

type RestController struct {
	network network.NetworkController
	web3    network.Web3Controller
	erc20   network.Erc20Controller
	db      network.Web3DbController
	shh     network.Web3ShhController
	abi     network.AbiController
}

// constructor like function
func newController(peer string, name string) RestController {
	ctl := RestController{}
	ctl.network = network.NewNetworkController()
	ctl.erc20 = network.NewErc20Controller(&ctl.network)
	ctl.web3 = network.NewWeb3Controller(&ctl.network)
	ctl.db = network.NewWeb3DbController(&ctl.network)
	ctl.shh = network.NewWeb3ShhController(&ctl.network)
	ctl.abi = network.NewAbiController()
	ctl.network.SetPeer(peer)
	ctl.network.SetTargetName(name)
	return ctl
}

// implemented method from interface RouterRegistrable
func (ctl RestController) RegisterRouters(router *echo.Group) {
	logger.Debug("registering rest controller api endpoints for network: ", ctl.network.Name())
	ctl.network.RegisterRouters(router)
	ctl.web3.RegisterRouters(router)
	ctl.erc20.RegisterRouters(router)
	ctl.db.RegisterRouters(router)
	ctl.shh.RegisterRouters(router)
	ctl.abi.RegisterRouters(router)
}

// constructor like function
func NewRopstenController() RestController {
	return newController(ropstenInfura, ropsten)
}

// constructor like function
func NewRinkebyController() RestController {
	return newController(rinkebyInfura, rinkeby)
}

// constructor like function
func NewKovanController() RestController {
	return newController(kovanInfura, kovan)
}

// constructor like function
func NewMainNetController() RestController {
	return newController(mainnetInfura, mainnet)
}

// constructor like function for user provided infura based connection
func NewInfuraController() RestController {
	return newController("", infura)
}

// constructor like function for user provided infura based connection
func NewQuikNodeController() RestController {
	return newController("", quiknode)
}

// constructor like function
func NewPrivateNetController() RestController {
	return newController("", private)
}
