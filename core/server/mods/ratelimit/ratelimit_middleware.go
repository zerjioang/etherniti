// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ratelimit

import (
	"net/http"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/thirdparty/echo"
)

var (
	//this variable might be concurrent accessed
	rateLimitEngine RateLimitEngine
)

func init() {
	// Create a cache with a default expiration time of 60 minutes, and which
	// purges expired items every 5 minutes
	//rateCache = cache.New(60*time.Minute, 5*time.Minute)
	rateLimitEngine = NewRateLimitEngine()
}

// REST API style rate limit middleware function.
// flood and abuse limit policy middleware
func RateLimit(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		header := c.Response().Header()
		clientIdentifier := c.RealIP()
		clientIdentifierBytes := str.UnsafeBytes(clientIdentifier)
		result := rateLimitEngine.Eval(clientIdentifierBytes, header)
		if result == Allow {
			return next(c)
		}
		return c.JSON(http.StatusTooManyRequests, rateExcedeed)
	}
}
