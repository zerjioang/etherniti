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
	ropstenInfura = "ropsten.infura.io/v3/"
	rinkebyInfura = "rinkeby.infura.io/v3/"
	kovanInfura   = "kovan.infura.io/v3/"
	mainnetInfura = "mainnet.infura.io/v3/"
)

var (
	infuraToken = "" //4f61378203ca4da4a6b6601bc16a22ad

	// configure endpoints
	UndefinedEndpoint = network.NewUndefinedConnection()

	// infura endpoints when user has a valid infura token
	ropstenInfuraEndpoint = network.NewUndefinedConnection()
	rinkebyInfuraEndpoint = network.NewUndefinedConnection()
	kovanInfuraEndpoint   = network.NewUndefinedConnection()
	mainnetInfuraEndpoint = network.NewUndefinedConnection()

	//custom user provided endpoints (if exists)
	// for connecting to different networks using it
	// prefered provider: own node, infura, etc
	ropstenCustom = ""
	rinkebyCustom = ""
	kovanCustom   = ""
	mainnetCustom = ""

	privateInfura   = network.NewNodeConnection("", "", "", "", "", "", infura)
	privateQuiknode = network.NewNodeConnection("", "", "", "", "", "", quiknode)
	privateCustom   = network.NewNodeConnection("", "", "", "", "", "", private)
	localGanache    = network.NewNodeConnection("", "", "", "", "", "", ganache)
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
	graphql network.GraphqlController
}

func init() {
	logger.Debug("loading infura token secret")
	// rad infura token
	infuraToken = cfg.InfuraToken()
	//update all infura related urls
	logger.Info("updating infura v3 endpoints with provided token")
	ropstenInfuraEndpoint = network.NewNodeConnection(httpsId, ropstenInfura+infuraToken, "", "", "", "", ropsten)
	rinkebyInfuraEndpoint = network.NewNodeConnection(httpsId, rinkebyInfura+infuraToken, "", "", "", "", rinkeby)
	kovanInfuraEndpoint = network.NewNodeConnection(httpsId, kovanInfura+infuraToken, "", "", "", "", kovan)
	mainnetInfuraEndpoint = network.NewNodeConnection(httpsId, mainnetInfura+infuraToken, "", "", "", "", mainnet)
	// load custom endpoints if exists
	logger.Info("loading user provided custom endpoints")
	ropstenCustom = cfg.RopstenCustomEndpoint
	rinkebyCustom = cfg.RinkebyCustomEndpoint
	kovanCustom = cfg.KovanCustomEndpoint
	mainnetCustom = cfg.MainnetCustomEndpoint
}

// constructor like function
func newController(client *fasthttp.Client, connection *network.NodeConnection) RestController {
	logger.Debug("creating new web3 controller")
	ctl := RestController{}
	ctl.network = network.NewNetworkController()
	ctl.erc20 = network.NewErc20Controller(&ctl.network)
	ctl.web3 = network.NewWeb3Controller(&ctl.network)
	ctl.db = network.NewWeb3DbController(&ctl.network)
	ctl.shh = network.NewWeb3ShhController(&ctl.network)
	ctl.devops = network.NewDevOpsController(&ctl.network)
	ctl.rpc = network.NewWeb3RpcController(&ctl.network)
	ctl.graphql = network.NewGraphqlController(&ctl.network)
	ctl.abi = network.NewAbiController()
	//configure target network parameters
	ctl.network.SetConnection(connection)
	ctl.network.SetTargetName(connection.Name())
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
	ctl.graphql.RegisterRouters(router)
	ctl.abi.RegisterRouters(router)
}

// constructor like function
func newInfuraController(client *fasthttp.Client, connection *network.NodeConnection, fallbackConnStr string) RestController {
	logger.Debug("creating new web3 controller for ", connection.Name(), " network")

	if infuraToken != "" && len(infuraToken) == 32 {
		// infura based controller is supported with default url
		return newController(client, connection)
	} else if fallbackConnStr != "" {
		// infura based controller is supported with user provided URL
		fllbackConn := network.NodeConnectionFromString(fallbackConnStr)
		return newController(client, fllbackConn)
	} else {
		// infura based controller not supported
		return newController(client, connection)
	}
}

// constructor like function
func NewRopstenController(client *fasthttp.Client) RestController {
	logger.Debug("creating new web3 controller for ropsten network")
	return newInfuraController(client, ropstenInfuraEndpoint, ropstenCustom)
}

// constructor like function
func NewRinkebyController(client *fasthttp.Client) RestController {
	logger.Debug("creating new web3 controller for rinkeby network")
	return newInfuraController(client, rinkebyInfuraEndpoint, rinkebyCustom)
}

// constructor like function
func NewKovanController(client *fasthttp.Client) RestController {
	logger.Debug("creating new web3 controller for kovan network")
	return newInfuraController(client, kovanInfuraEndpoint, kovanCustom)
}

// constructor like function
func NewMainNetController(client *fasthttp.Client) RestController {
	logger.Debug("creating new web3 controller for mainnet network")
	return newInfuraController(client, mainnetInfuraEndpoint, mainnetCustom)
}

// constructor like function for user provided infura based connection
func NewInfuraController(client *fasthttp.Client) RestController {
	logger.Debug("creating new web3 controller for infura network")
	return newController(client, privateInfura)
}

// constructor like function for user provided infura based connection
func NewQuikNodeController(client *fasthttp.Client) RestController {
	logger.Debug("creating new web3 controller for quiknode network")
	return newController(client, privateQuiknode)
}

// constructor like function
func NewGanacheController(client *fasthttp.Client) RestController {
	logger.Debug("creating new web3 controller for ganache testrpc")
	return newController(client, localGanache)
}

// constructor like function
func NewPrivateNetController(client *fasthttp.Client) RestController {
	logger.Debug("creating new web3 controller for private network")
	return newController(client, privateCustom)
}
