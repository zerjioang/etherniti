// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ratelimit

import (
	"net/http"
	"strconv"
	"time"

	"github.com/zerjioang/etherniti/core/modules/gocache"

	"github.com/zerjioang/etherniti/core/config"
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
	value    uint32
	reset    int64
	resetStr string
}
type RateLimitEngine struct {
	rateCache *gocache.Cache
}

// constructor like function
func NewRateLimitEngine() RateLimitEngine {
	rle := RateLimitEngine{}
	// Create a cache with a default expiration time of 1 minute, and which
	// purges expired items every 10 minutes
	rle.rateCache = gocache.New(1*time.Minute, 10*time.Minute)
	return rle
}

// ratelimit evaluation function
func (rte RateLimitEngine) Eval(clientIdentifier string, h http.Header) RateLimitResult {
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
		r := resetTime.Unix()
		currentRequestsLimit = limit{value: maxRatelimitValue, reset: r, resetStr: strconv.FormatInt(r, 10)}
	} else {
		currentRequestsLimit = rateRemaining.(limit)
	}

	//set rate limit reset time
	h.Set(XRateReset, currentRequestsLimit.resetStr)

	if currentRequestsLimit.value > 0 {
		//decrease counter limit and save it again
		currentRequestsLimit.value--
		// Want performance? Store pointers!
		rte.rateCache.Set(clientIdentifier, currentRequestsLimit, gocache.DefaultExpiration)

		// add current user remaining limit
		h.Set(XRateRemaining, strconv.FormatInt(int64(currentRequestsLimit.value), 10))

		//allow request
		return Allow
	}
	return Deny
}
