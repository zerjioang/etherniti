// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/handlers/clientcache"
	"github.com/zerjioang/etherniti/core/handlers/security"
	"github.com/zerjioang/etherniti/core/logger"
)

type SecurityController struct {
}

// contructor like function
func NewSecurityController() SecurityController {
	dc := SecurityController{}
	return dc
}

// return a blacklist of phishing sites.
// This list is maintained by GitHub user 409H at
// https://github.com/409H/EtherAddressLookup/blob/master/blacklists/domains.json
func (ctl SecurityController) domainBlacklist(c echo.Context) error {
	var code int
	code, c = clientcache.Cached(c, true, clientcache.CacheOneDay) // 24h cache directive
	return c.JSONBlob(code, security.DomainBlacklistBytesData())
}

// return a whitelist of non phishing sites,
// This list is maintained by the MetaMask project at
// https://github.com/MetaMask/eth-phishing-detect/blob/master/src/config.json .
func (ctl SecurityController) phisingWhitelist(c echo.Context) error {
	var code int
	code, c = clientcache.Cached(c, true, clientcache.CacheOneDay) // 24h cache directive
	return c.JSONBlob(code, security.PhishingWhitelistRawBytes())
}

// return a blacklist of phishing sites,
// as well as a whitelist and a fuzzylist.
// This list is maintained by the MetaMask project at
// https://github.com/MetaMask/eth-phishing-detect/blob/master/src/config.json .
func (ctl SecurityController) phisingBlacklist(c echo.Context) error {
	var code int
	code, c = clientcache.Cached(c, true, clientcache.CacheOneDay) // 24h cache directive
	return c.JSONBlob(code, security.PhishingBlacklistRawBytes())
}

func (ctl SecurityController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing index controller methods")
	router.GET("/security/domains/blacklist", ctl.domainBlacklist)
	router.GET("/security/phishing/blacklist", ctl.phisingBlacklist)
	router.GET("/security/phishing/whitelist", ctl.phisingWhitelist)
}
