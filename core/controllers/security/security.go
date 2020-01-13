// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package security

import (
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/controllers/wrap"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
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
func (ctl SecurityController) domainBlacklist(c *shared.EthernitiContext) error {
	c.OnSuccessCachePolicy = constants.CacheInfinite
	return api.SendSuccessBlob(c, DomainBlacklistBytesData())
}

// return a whitelist of non phishing sites,
// This list is maintained by the MetaMask project at
// https://github.com/MetaMask/eth-phishing-detect/blob/master/src/config.json
func (ctl SecurityController) phisingWhitelist(c *shared.EthernitiContext) error {
	c.OnSuccessCachePolicy = constants.CacheInfinite
	return api.SendSuccessBlob(c, PhishingWhitelistRawBytes())
}

// return a blacklist of phishing sites,
// as well as a whitelist and a fuzzylist.
// This list is maintained by the MetaMask project at
// https://github.com/MetaMask/eth-phishing-detect/blob/master/src/config.json
func (ctl SecurityController) phisingBlacklist(c *shared.EthernitiContext) error {
	c.OnSuccessCachePolicy = constants.CacheInfinite
	return api.SendSuccessBlob(c, PhishingBlacklistRawBytes())
}

// return a list of fuzzy domains
// This list is maintained by the MetaMask project at
// https://github.com/MetaMask/eth-phishing-detect/blob/master/src/config.json .
func (ctl SecurityController) fuzzylist(c *shared.EthernitiContext) error {
	c.OnSuccessCachePolicy = constants.CacheInfinite
	return api.SendSuccessBlob(c, FuzzyDataRawBytes())
}

// return whether given domain name is dangerous or not
func (ctl SecurityController) isDangerousDomain(c *shared.EthernitiContext) error {
	c.OnSuccessCachePolicy = constants.CacheInfinite
	domainName := c.Param("domain")
	return api.SendSuccessBlob(c, IsDangerousDomain(domainName))
}

func (ctl SecurityController) RegisterRouters(router *echo.Group) {
	logger.Debug("exposing index controller methods")
	router.GET("/security/check/:domain", wrap.Call(ctl.isDangerousDomain))
	router.GET("/security/domain-blacklist", wrap.Call(ctl.domainBlacklist))
	router.GET("/security/phishing-blacklist", wrap.Call(ctl.phisingBlacklist))
	router.GET("/security/phishing-whitelist", wrap.Call(ctl.phisingWhitelist))
	router.GET("/security/phishing-fuzzing", wrap.Call(ctl.fuzzylist))
}
