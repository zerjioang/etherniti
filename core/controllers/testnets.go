// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package controllers

import (
	"net/http"

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
	infuraToken = cfg.InfuraToken()
	//update all infura related urls
	logger.Info("updating infura v3 endpoints with provided token")
	ropstenInfura = ropstenInfura + infuraToken
	rinkebyInfura = rinkebyInfura + infuraToken
	kovanInfura = kovanInfura + infuraToken
	// load custom endpoints if exists
	logger.Info("loading user provided custom endpoints")
	ropstenCustom = cfg.RopstenCustomEndpoint
	rinkebyCustom = cfg.RinkebyCustomEndpoint
	kovanCustom = cfg.KovanCustomEndpoint
	mainnetCustom = cfg.MainnetCustomEndpoint

}

// constructor like function
func newController(client *http.Client, peer string, name string) RestController {
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
	ctl.network.SetClient(client)
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
func newInfuraController(client *http.Client, networkName, infuraEndpoint, fallbackEndpoint string) RestController {
	logger.Debug("creating new web3 controller for ", networkName, " network")

	if infuraToken != "" && len(infuraToken) == 32 {
		// infura based controller is supported with default url
		return newController(client, infuraEndpoint, networkName)
	} else if fallbackEndpoint != "" {
		// infura based controller is supported with user provided URL
		return newController(client, fallbackEndpoint, networkName)
	} else {
		// infura based controller not supported
		return newController(client, "", "unknown")
	}
}

// constructor like function
func NewRopstenController(client *http.Client) RestController {
	logger.Debug("creating new web3 controller for ropsten network")
	return newInfuraController(client, ropsten, ropstenInfura, ropstenCustom)
}

// constructor like function
func NewRinkebyController(client *http.Client) RestController {
	logger.Debug("creating new web3 controller for rinkeby network")
	return newInfuraController(client, rinkeby, rinkebyInfura, rinkebyCustom)
}

// constructor like function
func NewKovanController(client *http.Client) RestController {
	logger.Debug("creating new web3 controller for kovan network")
	return newInfuraController(client, kovan, kovanInfura, kovanCustom)
}

// constructor like function
func NewMainNetController(client *http.Client) RestController {
	logger.Debug("creating new web3 controller for mainnet network")
	return newInfuraController(client, mainnet, mainnetInfura, mainnetCustom)
}

// constructor like function for user provided infura based connection
func NewInfuraController(client *http.Client) RestController {
	logger.Debug("creating new web3 controller for infura network")
	return newController(client, "", infura)
}

// constructor like function for user provided infura based connection
func NewQuikNodeController(client *http.Client) RestController {
	logger.Debug("creating new web3 controller for quiknode network")
	return newController(client, "", quiknode)
}

// constructor like function
func NewPrivateNetController(client *http.Client) RestController {
	logger.Debug("creating new web3 controller for private network")
	return newController(client, "", private)
}
