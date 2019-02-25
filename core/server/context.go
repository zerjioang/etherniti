// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package server

import (
	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/eth/profile"
	"github.com/zerjioang/etherniti/core/eth/rpc"
)

// creating a custom context,
// allow us to add new features in a clean way
type EthernitiContext struct {
	echo.Context
}

// returns connectio profile from token information
func (context EthernitiContext) ConnectionProfile() (profile.ConnectionProfile, error) {
	requestProfileKeyContent := context.Request().Header.Get(config.HttpProfileHeaderkey)
	return profile.ParseConnectionProfileToken(requestProfileKeyContent)
}

// from incoming http request, it recovers the eth client linked to it
func (context EthernitiContext) ClientInstance() (ethrpc.EthRPC, error) {
	var client ethrpc.EthRPC
	return client, nil
}

// constructor like function
func NewEthernitiContext() EthernitiContext {
	ctx := EthernitiContext{}
	return ctx
}
