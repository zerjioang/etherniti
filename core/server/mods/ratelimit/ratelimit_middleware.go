package ratelimit

import (
	"net/http"

	"github.com/labstack/echo"
)

var (
	//this variable might be concurrent accessed
	rateLimitEngine RateLimitEngine
)

func init() {
	// Create a cache with a default expiration time of 60 minutes, and which
	// purges expired items every 5 minutes
	//rateCache = cache.New(60*time.Minute, 5*time.Minute)
	rateLimitEngine = RateLimitEngine{}
}

// REST API style rate limit middleware function.
func RateLimit(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Response().Header()
		valid := rateLimitEngine.Eval(&header)
		if valid {
			return next(c)
		}
		return c.JSON(http.StatusTooManyRequests, rateExcedeed)
	}
}
