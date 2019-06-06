// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package controllers

import (
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/controllers/network"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

const (
	ropsten  = "ropsten"
	rinkeby  = "rinkeby"
	kovan    = "kovan"
	mainnet  = "mainnet"
	infura   = "infura"
	quiknode = "quiknode"
	private  = "private"
)

var (
	ropstenInfura = "https://ropsten.infura.io/v3/"
	rinkebyInfura = "https://rinkeby.infura.io/v3/"
	kovanInfura   = "https://kovan.infura.io/v3/"
	mainnetInfura = "https://mainnet.infura.io/v3/"
	infuraToken   = "" //4f61378203ca4da4a6b6601bc16a22ad

	//custom endpoints
	ropstenCustom = ""
	rinkebyCustom = ""
	kovanCustom   = ""
	mainnetCustom = ""
)

type RestController struct {
	network network.NetworkController
	web3    network.Web3Controller
	erc20   network.Erc20Controller
	db      network.Web3DbController
	shh     network.Web3ShhController
	abi     network.AbiController
	devops  network.DevOpsController
}

func init() {
	logger.Debug("loading infura token secret")
	// rad infura token
	infuraToken = config.InfuraToken()
	//update all infura related urls
	logger.Debug("updating infura v3 endpoints with provided token")
	ropstenInfura = ropstenInfura + infuraToken
	rinkebyInfura = rinkebyInfura + infuraToken
	kovanInfura = kovanInfura + infuraToken
	infuraToken = infuraToken + infuraToken
	// load custom endpoints if exists
	//ropstenCustom := config.EndpointRopsten()

}

// constructor like function
func newController(peer string, name string) RestController {
	logger.Debug("creating new web3 controller")
	ctl := RestController{}
	ctl.network = network.NewNetworkController()
	ctl.erc20 = network.NewErc20Controller(&ctl.network)
	ctl.web3 = network.NewWeb3Controller(&ctl.network)
	ctl.db = network.NewWeb3DbController(&ctl.network)
	ctl.shh = network.NewWeb3ShhController(&ctl.network)
	ctl.devops = network.NewDevOpsController(&ctl.network)
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
	ctl.devops.RegisterRouters(router)
	ctl.abi.RegisterRouters(router)
}

// constructor like function
func newInfuraController(networkName, infuraEndpoint, fallbackEndpoint string) RestController {
	logger.Debug("creating new web3 controller for ", networkName, " network")

	if infuraToken != "" && len(infuraToken) == 32 {
		// infura token found and valid
		return newController(infuraEndpoint, networkName)
	} else if fallbackEndpoint != "" {
		// load ropsten controller with user provided URL
		return newController(fallbackEndpoint, networkName)
	} else {
		// ropsten not supported
		return newController("", "unknown")
	}
}

// constructor like function
func NewRopstenController() RestController {
	logger.Debug("creating new web3 controller for ropsten network")
	return newInfuraController(ropsten, ropstenInfura, ropstenCustom)
}

// constructor like function
func NewRinkebyController() RestController {
	logger.Debug("creating new web3 controller for rinkeby network")
	return newInfuraController(rinkeby, rinkebyInfura, rinkebyCustom)
}

// constructor like function
func NewKovanController() RestController {
	logger.Debug("creating new web3 controller for kovan network")
	return newInfuraController(kovan, kovanInfura, kovanCustom)
}

// constructor like function
func NewMainNetController() RestController {
	logger.Debug("creating new web3 controller for mainnet network")
	return newInfuraController(mainnet, mainnetInfura, mainnetCustom)
}

// constructor like function for user provided infura based connection
func NewInfuraController() RestController {
	logger.Debug("creating new web3 controller for infura network")
	return newController("", infura)
}

// constructor like function for user provided infura based connection
func NewQuikNodeController() RestController {
	logger.Debug("creating new web3 controller for quiknode network")
	return newController("", quiknode)
}

// constructor like function
func NewPrivateNetController() RestController {
	logger.Debug("creating new web3 controller for private network")
	return newController("", private)
}
