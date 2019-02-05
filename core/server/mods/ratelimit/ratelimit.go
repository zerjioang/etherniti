package ratelimit

import (
	"errors"
	"github.com/labstack/echo"
	"strconv"
	"sync/atomic"
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

const (
	rateLimit int32 = 4000
	rateLimitStr = "4000"
)

var (
	rateRemaining = rateLimit
	rateExcedeed = errors.New("rate limit reached")
	//rateCache *cache.Cache
)

func init(){
	// Create a cache with a default expiration time of 60 minutes, and which
	// purges expired items every 5 minutes
	//rateCache = cache.New(60*time.Minute, 5*time.Minute)
}

// REST API style rate limit middleware function.
func RateLimit(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		//get current request identifier
		// read request token. if not available fallback to user ip

		// read current limit
		currentRequestsLimit := atomic.LoadInt32(&rateRemaining)

		//inject rate limit headers
		h := c.Response().Header()
		// add rate limit max value: 4000
		h.Set(XRateLimit, rateLimitStr)
		// add current user remaining limit
		h.Set(XRateRemaining, strconv.FormatInt(int64(currentRequestsLimit), 10))

		if currentRequestsLimit > 0 {
			//allow request
			//decrease counter limit
			atomic.AddInt32(&rateRemaining, -1)
			return next(c)
		} else {
			//return rate limit excedeed message
			h.Set(XRateReset, "1549361984")
			return rateExcedeed
		}
	}
}
