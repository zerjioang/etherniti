// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package controllers

import (
	"github.com/valyala/fasthttp"

	"github.com/zerjioang/etherniti/core/controllers/network"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

const (
	httpId  = "http://"
	httpsId = "https://"

	ropsten  = "ropsten"
	rinkeby  = "rinkeby"
	kovan    = "kovan"
	mainnet  = "mainnet"
	infura   = "infura"
	quiknode = "quiknode"

	// ganache testnet node pointing always to default ganache config
	// http://127.0.0.1:7545
	ganache = "ganache"

	// private endpoints identifiers
	private = "private"

	//default endpoints
	ganacheEndpoint = httpId + "127.0.0.1:7545"

	//default infura public v3 endpoints
	ropstenInfura = httpsId + "ropsten.infura.io/v3/"
	rinkebyInfura = httpsId + "rinkeby.infura.io/v3/"
	kovanInfura   = httpsId + "kovan.infura.io/v3/"
	mainnetInfura = httpsId + "mainnet.infura.io/v3/"

	UndefinedEndpoint = ""
)

var (
	infuraToken = "" //4f61378203ca4da4a6b6601bc16a22ad

	// infura endpoints when user has a valid infura token
	ropstenInfuraEndpoint = UndefinedEndpoint
	rinkebyInfuraEndpoint = UndefinedEndpoint
	kovanInfuraEndpoint   = UndefinedEndpoint
	mainnetInfuraEndpoint = UndefinedEndpoint

	//custom user provided endpoints (if exists)
	// for connecting to different networks using it
	// prefered provider: own node, infura, etc
	ropstenCustom = UndefinedEndpoint
	rinkebyCustom = UndefinedEndpoint
	kovanCustom   = UndefinedEndpoint
	mainnetCustom = UndefinedEndpoint
)

type RestController struct {
	network network.NetworkController
	web3    network.Web3Controller
	erc20   network.Erc20Controller
	db      network.Web3DbController
	shh     network.Web3ShhController
	abi     network.AbiController
	devops  network.DevOpsController
	rpc     network.Web3RpcController
}

func init() {
	logger.Debug("loading infura token secret")
	// rad infura token
	infuraToken = cfg.InfuraToken()
	//update all infura related urls
	logger.Info("updating infura v3 endpoints with provided token")
	ropstenInfuraEndpoint = ropstenInfura + infuraToken
	rinkebyInfuraEndpoint = rinkebyInfura + infuraToken
	kovanInfuraEndpoint = kovanInfura + infuraToken
	mainnetInfuraEndpoint = mainnetInfura + infuraToken
	// load custom endpoints if exists
	logger.Info("loading user provided custom endpoints")
	ropstenCustom = cfg.RopstenCustomEndpoint
	rinkebyCustom = cfg.RinkebyCustomEndpoint
	kovanCustom = cfg.KovanCustomEndpoint
	mainnetCustom = cfg.MainnetCustomEndpoint
}

// constructor like function
func newController(client *fasthttp.Client, peer string, name string) RestController {
	logger.Debug("creating new web3 controller")
	ctl := RestController{}
	ctl.network = network.NewNetworkController()
	ctl.erc20 = network.NewErc20Controller(&ctl.network)
	ctl.web3 = network.NewWeb3Controller(&ctl.network)
	ctl.db = network.NewWeb3DbController(&ctl.network)
	ctl.shh = network.NewWeb3ShhController(&ctl.network)
	ctl.devops = network.NewDevOpsController(&ctl.network)
	ctl.rpc = network.NewWeb3RpcController(&ctl.network)
	ctl.abi = network.NewAbiController()
	//configure target network parameters
	ctl.network.SetPeer(peer)
	ctl.network.SetTargetName(name)
	ctl.network.SetClient(client)
	return ctl
}

// implemented method from interface RouterRegistrable
func (ctl RestController) RegisterRouters(router *echo.Group) {
	logger.Debug("registering rest controller api endpoints for network: ", ctl.network.Name())
	ctl.network.RegisterRouters(router)
	ctl.erc20.RegisterRouters(router)
	ctl.web3.RegisterRouters(router)
	ctl.db.RegisterRouters(router)
	ctl.shh.RegisterRouters(router)
	ctl.devops.RegisterRouters(router)
	ctl.rpc.RegisterRouters(router)
	ctl.abi.RegisterRouters(router)
}

// constructor like function
func newInfuraController(client *fasthttp.Client, networkName, infuraEndpoint, fallbackEndpoint string) RestController {
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
func NewRopstenController(client *fasthttp.Client) RestController {
	logger.Debug("creating new web3 controller for ropsten network")
	return newInfuraController(client, ropsten, ropstenInfuraEndpoint, ropstenCustom)
}

// constructor like function
func NewRinkebyController(client *fasthttp.Client) RestController {
	logger.Debug("creating new web3 controller for rinkeby network")
	return newInfuraController(client, rinkeby, rinkebyInfuraEndpoint, rinkebyCustom)
}

// constructor like function
func NewKovanController(client *fasthttp.Client) RestController {
	logger.Debug("creating new web3 controller for kovan network")
	return newInfuraController(client, kovan, kovanInfuraEndpoint, kovanCustom)
}

// constructor like function
func NewMainNetController(client *fasthttp.Client) RestController {
	logger.Debug("creating new web3 controller for mainnet network")
	return newInfuraController(client, mainnet, mainnetInfuraEndpoint, mainnetCustom)
}

// constructor like function for user provided infura based connection
func NewInfuraController(client *fasthttp.Client) RestController {
	logger.Debug("creating new web3 controller for infura network")
	return newController(client, UndefinedEndpoint, infura)
}

// constructor like function for user provided infura based connection
func NewQuikNodeController(client *fasthttp.Client) RestController {
	logger.Debug("creating new web3 controller for quiknode network")
	return newController(client, UndefinedEndpoint, quiknode)
}

// constructor like function
func NewGanacheController(client *fasthttp.Client) RestController {
	logger.Debug("creating new web3 controller for ganache testrpc")
	return newController(client, ganacheEndpoint, ganache)
}

// constructor like function
func NewPrivateNetController(client *fasthttp.Client) RestController {
	logger.Debug("creating new web3 controller for private network")
	return newController(client, "", private)
}
