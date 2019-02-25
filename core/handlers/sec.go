// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
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
	return c.JSONBlob(http.StatusOK, indexWelcomeBytes)
}

// return a blacklist of phishing sites,
// as well as a whitelist and a fuzzylist.
// This list is maintained by the MetaMask project at
// https://github.com/MetaMask/eth-phishing-detect/blob/master/src/config.json .
func (ctl SecurityController) phisingBlacklist(c echo.Context) error {
	return c.JSONBlob(http.StatusOK, indexWelcomeBytes)
}

func (ctl SecurityController) RegisterRouters(router *echo.Echo) {
	log.Info("exposing index controller methods")
	router.GET("/v1/security/blacklist/domains", ctl.domainBlacklist)
	router.GET("/v1/security/blacklist/phishing", ctl.phisingBlacklist)
}
