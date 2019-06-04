// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ratelimit

import (
	"net/http"
	"strconv"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/modules/cache"
	"github.com/zerjioang/etherniti/core/modules/fastime"
)

// 4,000 requests per hour.
// example:
// X-RateLimit-Limit: 4000
// X-RateLimit-Remaining: 56
// X-RateLimit-Reset: 1434037662
const (
	// The maximum number of requests that the user is allowed to make per hour.
	XRateLimit = "X-Rate-Limit-Limit"
	// The number of requests remaining in the current rate limit window.
	XRateRemaining = "X-Rate-Limit-Remaining"
	// The time at which the current rate limit window resets in UTC epoch seconds.
	XRateReset = "X-Rate-Limit-Reset"
)

type RateLimitResult bool

const (
	Deny  RateLimitResult = false
	Allow RateLimitResult = true
)

var (
	defaultCacheMeasurementUnitFt = config.RateLimitUnitsFt()
	maxRatelimitStr               = config.RateLimitStr()
	maxRatelimitValue             = config.RateLimit()
)

type limit struct {
	value uint32
	reset int64
}
type RateLimitEngine struct {
	rateCache *cache.MemoryCache
}

// constructor like function
func NewRateLimitEngine() RateLimitEngine {
	rle := RateLimitEngine{}
	rle.rateCache = cache.Instance()
	return rle
}

// ratelimit evaluation function
func (rte RateLimitEngine) Eval(clientIdentifier []byte, h http.Header) RateLimitResult {
	if h == nil {
		return Deny
	}

	//get current time
	timeNow := fastime.Now()
	resetTime := timeNow.Add(defaultCacheMeasurementUnitFt)

	//inject rate limit header: X-Rate-Limit-Limit
	h.Set(XRateLimit, maxRatelimitStr)

	// read current limit
	var currentRequestsLimit limit
	rateRemaining, found := rte.rateCache.Get(clientIdentifier)
	if !found {
		// initialize client max allowed rate limit
		currentRequestsLimit = limit{value: maxRatelimitValue, reset: resetTime.Unix()}
	} else {
		currentRequestsLimit = rateRemaining.(limit)
	}

	if currentRequestsLimit.value > 0 {
		//decrease counter limit and save it again
		currentRequestsLimit.value--
		rte.rateCache.Set(clientIdentifier, currentRequestsLimit)

		// add current user remaining limit
		h.Set(XRateRemaining, strconv.FormatInt(int64(currentRequestsLimit.value), 10))

		//set rate limit reset time
		rateResetStr := strconv.FormatInt(currentRequestsLimit.reset, 10)
		h.Set(XRateReset, rateResetStr)

		//allow request
		return Allow
	}
	return Deny
}
