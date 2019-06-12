// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"errors"
	"net/http"

	"github.com/zerjioang/etherniti/core/data"

	"github.com/zerjioang/etherniti/core/modules/cache"

	"github.com/zerjioang/etherniti/core/eth"
	ethrpc "github.com/zerjioang/etherniti/core/eth/rpc"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

var (
	errGetCaller      = errors.New(data.DataBindFailedStr)
	errInvalidAddress = errors.New(data.AddressNoSetupStr)
)

// eth network controller
type NetworkController struct {
	// http client
	client *http.Client
	//main connection peer address/ip
	peer string
	//connection name: mainet, ropsten, rinkeby, etc
	networkName string
	//ethereum interaction cache
	cache *cache.MemoryCache
}

// constructor like function
func NewNetworkController() NetworkController {
	ctl := NetworkController{}
	ctl.cache = cache.NewMemoryCache()
	return ctl
}

func (ctl *NetworkController) SetClient(c *http.Client) {
	ctl.client = c
}

func (ctl *NetworkController) SetPeer(peerLocation string) {
	ctl.peer = peerLocation
}

func (ctl *NetworkController) GetPeer() string {
	return ctl.peer
}

func (ctl *NetworkController) SetTargetName(networkName string) {
	ctl.networkName = networkName
}

// implemented method from interface RouterRegistrable
func (ctl *NetworkController) RegisterRouters(router *echo.Group) {
}

func (ctl *NetworkController) Name() string {
	return ctl.networkName
}

func (ctl *NetworkController) getRpcClient(c *echo.Context) (*ethrpc.EthRPC, error) {
	// get our client context
	client, cId, cliErr := c.RecoverEthClientFromTokenOrPeerUrl(ctl.peer, ctl.client)
	logger.Info("controller request using context id: ", cId)
	if cliErr != nil {
		logger.Error("failed to build an eth client from current context. missing connection url: ", cliErr)
		return nil, cliErr
	}
	return client, nil
}

func (ctl *NetworkController) getCallerAddress(c *echo.Context) (string, error) {
	from := c.CallerEthAddress()
	if !eth.IsValidAddressLow(from) {
		return "", errInvalidAddress
	}
	return from, nil
}
