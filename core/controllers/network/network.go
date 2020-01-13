// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"errors"
	"hash/fnv"
	"strconv"

	web3 "github.com/zerjioang/go-hpc/lib/eth/rpc/client"

	"github.com/zerjioang/etherniti/shared"

	"github.com/zerjioang/etherniti/core/logger"

	"github.com/zerjioang/etherniti/core/api"

	"github.com/zerjioang/etherniti/core/data"

	"github.com/zerjioang/go-hpc/lib/cache"

	"github.com/zerjioang/go-hpc/lib/eth"
	ethrpc "github.com/zerjioang/go-hpc/lib/eth/rpc"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

var (
	errGetCaller      = errors.New(data.DataBindFailedStr)
	errInvalidAddress = errors.New(data.AddressNoSetupStr)
)

// eth network controller
type NetworkController struct {
	// http client
	client *web3.EthClient
	// node connection information
	connection *NodeConnection
	//ethereum interaction cache
	cache *cache.MemoryCache
	// predefined rpc client
	rpclient *ethrpc.EthRPC
}

// constructor like function
func NewNetworkController() NetworkController {
	ctl := NetworkController{}
	ctl.connection = NewNodeConnection("http://", "127.0.0.1", "7545", "8547", "", "", "default")
	ctl.cache = cache.NewMemoryCache()
	return ctl
}

func (ctl *NetworkController) SetClient(c *web3.EthClient) {
	ctl.client = c
}

func (ctl *NetworkController) SetRpcClient(rpclient *ethrpc.EthRPC) {
	ctl.rpclient = rpclient
}

func (ctl *NetworkController) SetConnection(c *NodeConnection) {
	ctl.connection = c
}

func (ctl *NetworkController) GetRPCEndpoint() string {
	return ctl.connection.GetRPCEndpoint()
}

func (ctl *NetworkController) GetGraphQLEndpoint() string {
	return ctl.connection.GetGraphQLEndpoint()
}

func (ctl *NetworkController) SetTargetName(networkName string) {
	ctl.connection.name = networkName
}

// implemented method from interface RouterRegistrable
func (ctl *NetworkController) RegisterRouters(router *echo.Group) {
}

func (ctl *NetworkController) Name() string {
	return ctl.connection.Name()
}

func (ctl *NetworkController) getRpcClient(c *shared.EthernitiContext) (*ethrpc.EthRPC, error) {
	//check if current newtwork has a predefined rpc controller or not
	// network with predefined controllers are: rinkeby, kovan, ganache, infura
	logger.Info("checking if exists a predefined rpc client for current network")
	if ctl.rpclient != nil {
		return ctl.rpclient, nil
	} else {
		// predefined rpc client not found. resolve it and setup
		// get our client context
		peerLocation := ctl.GetRPCEndpoint()
		client, cId, cliErr := c.RecoverEthClientFromTokenOrPeerUrl(peerLocation, ctl.client)
		logger.Info("controller request using context id: ", cId)
		if cliErr != nil {
			logger.Error("failed to build an eth client from current context. missing connection url: ", cliErr)
			return nil, cliErr
		}
		ctl.rpclient = client
		return client, nil
	}
}

func (ctl *NetworkController) getCallerAddress(c *shared.EthernitiContext) (string, error) {
	from := c.CallerEthAddress()
	if !eth.IsValidAddressLow(from) {
		return "", errInvalidAddress
	}
	return from, nil
}

func (ctl *NetworkController) Noop(c *shared.EthernitiContext) error {
	return api.Error(c, errors.New("not implemented"))
}

// uniqueid is a combination of all network parameters so that it can be returned a unique
// network identifier for caching purposes, etc
func (ctl *NetworkController) UniqueId() string {
	hash := fnv.New64a()
	_, _ = hash.Write([]byte(ctl.connection.GetRPCEndpoint()))
	uid := hash.Sum64()
	// Format to a string by passing the number and it's base.
	uidstr := strconv.FormatUint(uid, 10)
	return uidstr
}
