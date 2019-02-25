// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/zerjioang/etherniti/core/handlers/security"
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
	return c.JSONBlob(http.StatusOK, security.DomainBlacklistRawBytes)
}

// return a whitelist of non phishing sites,
// This list is maintained by the MetaMask project at
// https://github.com/MetaMask/eth-phishing-detect/blob/master/src/config.json .
func (ctl SecurityController) phisingWhitelist(c echo.Context) error {
	return c.JSONBlob(http.StatusOK, security.PhishingWhitelistRawBytes)
}

// return a blacklist of phishing sites,
// as well as a whitelist and a fuzzylist.
// This list is maintained by the MetaMask project at
// https://github.com/MetaMask/eth-phishing-detect/blob/master/src/config.json .
func (ctl SecurityController) phisingBlacklist(c echo.Context) error {
	return c.JSONBlob(http.StatusOK, security.PhishingBlacklistRawBytes)
}

func (ctl SecurityController) RegisterRouters(router *echo.Echo) {
	log.Info("exposing index controller methods")
	router.GET("/v1/security/domains/blacklist", ctl.domainBlacklist)
	router.GET("/v1/security/phishing/blacklist", ctl.phisingBlacklist)
	router.GET("/v1/security/phishing/whitelist", ctl.phisingWhitelist)
}
