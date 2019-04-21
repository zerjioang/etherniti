// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package tor

import (
	"github.com/zerjioang/etherniti/core/modules/tor"
	"net/http"

	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/thirdparty/echo"
)

var (
	torBlockedError = protocol.NewApiError(http.StatusNotAcceptable, "Tor based connection blocked")
)

// REST API style Tor IP blocker
func BlockTorConnections(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.ContextInterface) error {

		//get current request ip
		requestIp := c.Request().RemoteAddr

		found := tor.TornodeSet.Contains(requestIp)
		if !found {
			//received request IP is not blacklisted
			return next(c)
		} else {
			// received request is done using on of the blacklisted tor nodes
			//return rate limit excedeed message
			return c.JSON(http.StatusTooManyRequests, torBlockedError)
		}
	}
}
