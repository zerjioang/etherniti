// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package server

import (
	"errors"

	"github.com/zerjioang/etherniti/shared/constants"

	"github.com/zerjioang/etherniti/core/util"

	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/eth/profile"
	"github.com/zerjioang/etherniti/core/eth/rpc"
)

// creating a custom context,
// allow us to add new features in a clean way
type EthernitiContext struct {
	echo.Context
	// connection profile data for interaction
	profileData profile.ConnectionProfile
}

var (
	noPeerAddressError = errors.New("no peer address to connect defined")
)

// returns connection profile from token information
func (context *EthernitiContext) ConnectionProfileSetup() (profile.ConnectionProfile, error) {
	requestProfileKeyContent := context.ReadConnectionProfileToken()
	readedProfile, err := profile.ParseConnectionProfileToken(requestProfileKeyContent)
	if err == nil {
		//save profile data
		context.profileData = readedProfile
	}
	return readedProfile, err
}

// it recovers the eth client linked to it
// if peer url is provided, this peer address is used to dial
// otherwise, token information is readed in order to custom peer dial
func (context EthernitiContext) RecoverEthClientFromTokenOrPeerUrl(peerUrl string) (ethrpc.EthRPC, string, error) {
	var client ethrpc.EthRPC
	var contextId string
	// by default, peer url is used to dial
	if peerUrl == "" {
		//no peer url found, try to read from user token
		if context.profileData.Peer == "" {
			return client, "", noPeerAddressError
		}
		contextId = context.profileData.Peer
	} else {
		// use peer url
		contextId = peerUrl
	}
	client = ethrpc.NewDefaultRPC(contextId)
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
	data, encErr := util.StdMarshal(i)
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
func NewEthernitiContext(c echo.Context) *EthernitiContext {
	ctx := new(EthernitiContext)
	ctx.Context = c
	return ctx
}
