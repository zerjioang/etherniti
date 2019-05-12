// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/handlers/clientcache"
	"github.com/zerjioang/etherniti/core/handlers/security"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
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
func (ctl SecurityController) domainBlacklist(c *echo.Context) error {
	c.OnSuccessCachePolicy = clientcache.CacheInfinite
	return api.SendSuccessBlob(c, security.DomainBlacklistBytesData())
}

// return a whitelist of non phishing sites,
// This list is maintained by the MetaMask project at
// https://github.com/MetaMask/eth-phishing-detect/blob/master/src/config.json .
func (ctl SecurityController) phisingWhitelist(c *echo.Context) error {
	c.OnSuccessCachePolicy = clientcache.CacheInfinite
	return api.SendSuccessBlob(c, security.PhishingWhitelistRawBytes())
}

// return a blacklist of phishing sites,
// as well as a whitelist and a fuzzylist.
// This list is maintained by the MetaMask project at
// https://github.com/MetaMask/eth-phishing-detect/blob/master/src/config.json .
func (ctl SecurityController) phisingBlacklist(c *echo.Context) error {
	c.OnSuccessCachePolicy = clientcache.CacheInfinite
	return api.SendSuccessBlob(c, security.PhishingBlacklistRawBytes())
}

// return a list of fuzzy domains
// This list is maintained by the MetaMask project at
// https://github.com/MetaMask/eth-phishing-detect/blob/master/src/config.json .
func (ctl SecurityController) fuzzylist(c *echo.Context) error {
	c.OnSuccessCachePolicy = clientcache.CacheInfinite
	return api.SendSuccessBlob(c, security.FuzzyDataRawBytes())
}

// return whether given domain name is dangerous or not
func (ctl SecurityController) isDangerousDomain(c *echo.Context) error {
	c.OnSuccessCachePolicy = clientcache.CacheInfinite
	domainName := c.Param("domain")
	return api.SendSuccessBlob(c, security.IsDangerousDomain(domainName))
}

func (ctl SecurityController) RegisterRouters(router *echo.Group) {
	logger.Debug("exposing index controller methods")
	router.GET("/security/check/:domain", ctl.isDangerousDomain)
	router.GET("/security/domain-blacklist", ctl.domainBlacklist)
	router.GET("/security/phishing-blacklist", ctl.phisingBlacklist)
	router.GET("/security/phishing-whitelist", ctl.phisingWhitelist)
	router.GET("/security/phishing-fuzzing", ctl.fuzzylist)
}
