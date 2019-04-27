// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"errors"
	"github.com/zerjioang/etherniti/core/modules/cache"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/eth"
	"github.com/zerjioang/etherniti/core/eth/rpc"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/server"

	"github.com/zerjioang/etherniti/thirdparty/echo"
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
	ctl.cache = cache.NewMemoryCache()
	return ctl
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

func (ctl *NetworkController) getRpcClient(c echo.ContextInterface) (*ethrpc.EthRPC, error) {
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return nil, api.ErrorStr(c, "failed to execute requested operation")
	}
	// get our client context
	client, cId, cliErr := cc.RecoverEthClientFromTokenOrPeerUrl(ctl.peer)
	logger.Info("controller request using context id: ", cId)
	if cliErr != nil {
		return nil, api.Error(c, cliErr)
	}
	return client, nil
}

func (ctl *NetworkController) getCallerAddress(c echo.ContextInterface) (string, error) {
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return "", errors.New("failed to execute requested operation")
	}
	from := cc.CallerEthAddress()
	if !eth.IsValidAddress(from) {
		return "", errors.New("invalid ethereum address setup when creating connection profile. Please provide a alid address as 'from'")
	}
	return from, nil
}