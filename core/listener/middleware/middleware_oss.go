// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// edition tags: opensource
// +build oss

package middleware

import (
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

// this is open-source edition middleware
func secure(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		// get request
		response := c.Response()
		rh := response.Header()

		// add default security headers
		// h.Set("access-control-allow-credentials", "true")
		rh.Set("x-xss-protection", "1; mode=block")
		rh.Set("strict-transport-security", "max-age=63072000; includeSubDomains; preload ") //2 years
		//public-key-pins: pin-sha256="t/OMbKSZLWdYUDmhOyUzS+ptUbrdVgb6Tv2R+EMLxJM="; pin-sha256="PvQGL6PvKOp6Nk3Y9B7npcpeL40twdPwZ4kA2IiixqA="; pin-sha256="ZyZ2XrPkTuoiLk/BR5FseiIV/diN3eWnSewbAIUMcn8="; pin-sha256="0kDINA/6eVxlkns5z2zWv2/vHhxGne/W0Sau/ypt3HY="; pin-sha256="ktYQT9vxVN4834AQmuFcGlSysT1ZJAxg+8N1NkNG/N8="; pin-sha256="rwsQi0+82AErp+MzGE7UliKxbmJ54lR/oPheQFZURy8="; max-age=600; report-uri="https://www.keycdn.com"
		rh.Set("X-Content-Type-Options", "nosniff")
		rh.Set("Expect-CT", "enforce, max-age=30")
		rh.Set("X-UA-Compatible", "IE=Edge,chrome=1")
		rh.Set("x-frame-options", "SAMEORIGIN")
		rh.Set("Referrer-Policy", "same-origin")
		rh.Set("Feature-Policy", "microphone 'none'; payment 'none'; sync-xhr 'self'")
		rh.Set("x-firefox-spdy", "h2")

		return next(c)
	}
}
