package ratelimit

import (
	"github.com/zerjioang/etherniti/core/api"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
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
	rateLimit    int32 = 4000
	rateLimitStr       = "4000"
)

var (
	rateRemaining = rateLimit
	rateExcedeed  = api.NewApiError(http.StatusTooManyRequests, "rate limit reached")
	//rateCache *cache.Cache
)

type RateLimitEngine struct {

}

// constructor like function
func NewRateLimitEngine() RateLimitEngine {
	rle := RateLimitEngine{}
	return rle
}

// ratelimit evaluation function
func (rte RateLimitEngine) Eval(h *http.Header) bool {
	//get current request identifier
	// read request token. if not available fallback to user ip

	//validate input data
	if h == nil {
		return false
	}

	//get current time
	timeNow := time.Now()
	//currentTime := timeNow.Unix()
	afterHourTime := timeNow.Add(60 * time.Minute)

	// read current limit
	currentRequestsLimit := atomic.LoadInt32(&rateRemaining)

	//inject rate limit headers
	// add rate limit max value: 4000
	h.Set(XRateLimit, rateLimitStr)

	if currentRequestsLimit > 0 {
		//decrease counter limit
		atomic.AddInt32(&rateRemaining, -1)

		// add current user remaining limit
		h.Set(XRateRemaining, strconv.FormatInt(int64(currentRequestsLimit), 10))

		//allow request
		return true
	} else {
		//return rate limit excedeed message
		rateResetStr := strconv.FormatInt(afterHourTime.Unix(), 10)
		h.Set(XRateReset, rateResetStr)
		return false
	}
}