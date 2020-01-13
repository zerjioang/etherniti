// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ratelimit

import (
	"github.com/pkg/errors"
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/shared"

	"github.com/zerjioang/go-hpc/lib/codes"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

var (
	//this variable might be concurrent accessed
	rateLimitEngine     RateLimitEngine
	rateLimitExcededErr = errors.New("rate limit reached")
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
	return func(c echo.Context) error {
		cc, ok := c.(*shared.EthernitiContext)
		if ok && cc != nil {
			header := c.Response().Header()
			clientIdentifier := cc.RateLimitIdentifier()
			result := rateLimitEngine.Eval(clientIdentifier, header)
			if result == Allow {
				return next(cc)
			}
			return api.ErrorCode(cc, codes.StatusTooManyRequests, rateLimitExcededErr)
		}
		return errors.New("invalid request")
	}
}
