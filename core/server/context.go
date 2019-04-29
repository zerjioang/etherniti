// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package server

import (
	"errors"
	"sync"

	"github.com/zerjioang/etherniti/core/config"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/shared/constants"

	"github.com/zerjioang/etherniti/core/eth/profile"
	"github.com/zerjioang/etherniti/core/eth/rpc"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

// creating a custom context,
// allow us to add new features in a clean way
type EthernitiContext struct {
	echo.ContextInterface
	// connection profile data for interaction
	profileLock *sync.Mutex
	profileData profile.ConnectionProfile

	// http response data
	response *echo.Response
}

var (
	noPeerAddressError = errors.New("no peer address to connect defined")
	isDebug            = config.IsDevelopment()
)

// returns connection profile from token information
func (context *EthernitiContext) ConnectionProfileSetup() (profile.ConnectionProfile, error) {
	requestProfileKeyContent := context.ReadConnectionProfileToken()
	readedProfile, err := profile.ParseConnectionProfileToken(requestProfileKeyContent)
	if err == nil {
		//save profile data
		context.profileLock.Lock()
		context.profileData = readedProfile
		context.profileLock.Unlock()
	}
	return readedProfile, err
}

// get caller eth address
func (context *EthernitiContext) CallerEthAddress() string {
	context.profileLock.Lock()
	from := context.profileData.Address
	context.profileLock.Unlock()
	return from
}

// it recovers the eth client linked to it
// if peer url is provided, this peer address is used to dial
// otherwise, token information is readed in order to custom peer dial
func (context EthernitiContext) RecoverEthClientFromTokenOrPeerUrl(peerUrl string) (*ethrpc.EthRPC, string, error) {
	client := new(ethrpc.EthRPC)
	var contextId string
	// by default, peer url is used to dial
	if peerUrl == "" {
		//no peer url found, try to read from user token
		if context.profileData.RpcEndpoint == "" {
			return client, "", noPeerAddressError
		}
		contextId = context.profileData.RpcEndpoint
	} else {
		// use peer url
		contextId = peerUrl
	}
	client = ethrpc.NewDefaultRPCPtr(contextId, isDebug)
	return client, contextId, nil
}

// reads connection profile token from allowed sources
func (context EthernitiContext) ReadConnectionProfileToken() string {
	req := context.Request()

	var tokenDataStr string
	// read if token provided via header key
	tokenDataStr = req.Header.Get(constants.HttpProfileHeaderkey)
	if tokenDataStr == "" {
		//read if token provided via query param
		tokenDataStr = context.QueryParam("token")
	}
	return tokenDataStr
}

//custom json encoder
func (context EthernitiContext) JSON(code int, i interface{}) (err error) {
	data, encErr := str.StdMarshal(i)
	if encErr != nil {
		return encErr
	}
	return context.JSONBlob(code, data)
}

func (context *EthernitiContext) writeContentType(value string) {
	header := context.Response().Header()
	if header.Get(echo.HeaderContentType) == "" {
		header.Set(echo.HeaderContentType, value)
	}
}

func (context *EthernitiContext) FastBlob(code int, contentType string, b []byte) (err error) {
	response := context.Response()
	header := response.HeaderPtr()
	if header.Get(echo.HeaderContentType) == "" {
		header.Set(echo.HeaderContentType, contentType)
	}
	response.WriteHeader(code)
	_, err = response.Write(b)
	return
}

func (context *EthernitiContext) Blob(code int, contentType string, b []byte) (err error) {
	context.writeContentType(contentType)
	response := context.Response()
	response.WriteHeader(code)
	_, err = response.Write(b)
	return
}

func (context *EthernitiContext) HTMLBlob(code int, b []byte) (err error) {
	return context.Blob(code, echo.MIMETextHTMLCharsetUTF8, b)
}

// constructor like function
func NewEthernitiContext(c echo.ContextInterface) *EthernitiContext {
	ctx := new(EthernitiContext)
	ctx.ContextInterface = c
	ctx.profileLock = new(sync.Mutex)
	return ctx
}
