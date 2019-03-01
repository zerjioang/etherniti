package handlers

import (
	"strings"

	"github.com/zerjioang/etherniti/core/api/protocol"

	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/server"
	"github.com/zerjioang/etherniti/core/util"
)

// custom http error handler. returns error messages as json
func customHTTPErrorHandler(err error, c echo.Context) {
	// use code snippet below to customize http return code
	/*
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
	*/
	_ = protocol.Error(c, err)
}

// http to http redirect function
func httpsRedirect(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		scheme := c.Scheme()
		// host := req.Host
		if scheme == "http" {
			return c.Redirect(301, config.GetRedirectUrl(req.Host, req.RequestURI))
		}
		return next(c)
	}
}

// hardening middleware function.
func hardening(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add security headers
		h := c.Response().Header()
		h.Set("server", "Apache")
		// h.Set("access-control-allow-credentials", "true")
		h.Set("x-xss-protection", "1; mode=block")
		h.Set("strict-transport-security", "max-age=31536000; includeSubDomains; preload")
		//public-key-pins: pin-sha256="t/OMbKSZLWdYUDmhOyUzS+ptUbrdVgb6Tv2R+EMLxJM="; pin-sha256="PvQGL6PvKOp6Nk3Y9B7npcpeL40twdPwZ4kA2IiixqA="; pin-sha256="ZyZ2XrPkTuoiLk/BR5FseiIV/diN3eWnSewbAIUMcn8="; pin-sha256="0kDINA/6eVxlkns5z2zWv2/vHhxGne/W0Sau/ypt3HY="; pin-sha256="ktYQT9vxVN4834AQmuFcGlSysT1ZJAxg+8N1NkNG/N8="; pin-sha256="rwsQi0+82AErp+MzGE7UliKxbmJ54lR/oPheQFZURy8="; max-age=600; report-uri="https://www.keycdn.com"
		h.Set("X-Content-Type-Options", "nosniff")
		//h.Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline' font-src fonts.googleapis.com fonts.gstatic.com")
		h.Set("Expect-CT", "enforce, max-age=30")
		h.Set("X-UA-Compatible", "IE=Edge,chrome=1")
		h.Set("x-frame-options", "SAMEORIGIN")
		h.Set("Referrer-Policy", "same-origin")
		h.Set("Feature-Policy", "microphone 'none'; payment 'none'; sync-xhr 'self'")
		h.Set("x-firefox-spdy", "h2")
		return next(c)
	}
}

// fake server headers middleware function.
func fakeServer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add fake server header
		h := c.Response().Header()
		h.Set("server", "Apache/2.0.54")
		h.Set("x-powered-by", "PHP/5.1.6")
		return next(c)
	}
}

// bots blacklist function.
func antiBots(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add antibots policy
		ua := c.Request().UserAgent()
		ua = util.ToLowerAscii(ua)
		if ua == "" {
			//drop the request
			logger.Warn("drop request: no user-agent provided")
			return userAgentErr
		} else if isBotRequest(ua) {
			//drop the request
			logger.Warn("drop request: provided user-agent is considered as a bot: ", ua)
			return userAgentErr
		}
		return next(c)
	}
}

// check if user agent string contains bot strings similarities
func isBotRequest(userAgent string) bool {
	var lock = false
	var size = len(api.BadBotsList)
	for i := 0; i < size && !lock; i++ {
		lock = strings.Contains(userAgent, api.BadBotsList[i])
	}
	return lock
}

// check if http request host value is allowed or not
func hostnameCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add Host policy
		h := c.Request().Host
		chunks := strings.Split(h, ":")
		var hostname = ""
		if len(chunks) == 1 {
			//no port defined in host header
			hostname = h
		} else if len(chunks) == 2 {
			//port defined in host header
			hostname = chunks[0]
		}
		var allowed = false
		var size = len(config.AllowedHostnames)
		for i := 0; i < size && !allowed; i++ {
			allowed = strings.Compare(hostname, config.AllowedHostnames[i]) == 0
		}
		if allowed {
			// fordward request to next middleware
			return next(c)
		} else {
			// drop the request
			return nil
		}
	}
}

// keepalive middleware function.
func keepalive(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add keep alive headers in the response if requested by the client
		connectionMode := c.Request().Header.Get("Connection")
		connectionMode = util.ToLowerAscii(connectionMode)
		if connectionMode == "keep-alive" {
			// keep alive connection mode requested
			h := c.Response().Header()
			h.Set("Connection", "Keep-Alive")
			h.Set("Keep-Alive", "timeout=5, max=1000")
		}
		return next(c)
	}
}

// jwt middleware function.
func customContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// convert context in etherniti context
		cc := server.NewEthernitiContext(c)
		return next(cc)
	}
}
