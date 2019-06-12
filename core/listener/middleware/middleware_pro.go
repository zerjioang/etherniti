// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// edition tags: organization and enterprise
// +build pro

package middleware

import (
	"strings"

	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/modules/bots"
	"github.com/zerjioang/etherniti/core/modules/tor"
	ip2 "github.com/zerjioang/etherniti/core/util/ip"

	"github.com/zerjioang/etherniti/core/modules/badips"
	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

// this is enterprise edition middleware
func secure(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		// get request
		request := c.Request()
		h := request.Header
		response := c.Response()
		rh := response.Header()

		// add abuseIP policy
		ip := c.RealIP()
		if ip == "" {
			//drop the request
			logger.Warn("drop request: no ip provided")
			return securityErr
		} else if badips.IsBackListedIp(ip) {
			//drop the request
			logger.Warn("drop request: blacklisted ip detected: ", ip)
			return securityErr
		}
		// add antibots policy
		ua := request.UserAgent()
		ua = str.ToLowerAscii(ua)
		if ua == "" {
			//drop the request
			logger.Warn("drop request: no user-agent provided")
			return securityErr
		} else if len(ua) < 4 || bots.GetBadBotsList().MatchAny(ua) {
			//drop the request
			logger.Warn("drop request: provided user-agent is considered as a bot: ", ua)
			return securityErr
		}

		// add hostname policy
		host := request.Host
		chunks := strings.Split(host, ":")
		var hostname = ""
		if len(chunks) == 1 {
			//no port defined in host header
			hostname = host
		} else if len(chunks) == 2 {
			//port defined in host header
			hostname = chunks[0]
		}
		allowed := cfg.AllowedHostnames.Contains(hostname)
		if !allowed {
			// drop the request
			logger.Warn("drop request: provided request does not specifies a valid host name in http headers")
			return securityErr
		}

		if cfg.BlockTorConnections {
			// add rate limit control
			logger.Info("[LAYER] tor connections blocker middleware added")
			//get current request ip
			requestIp := request.RemoteAddr
			ipUint32 := ip2.Ip2intLow(requestIp)
			found := tor.TornodeSet.Contains(ipUint32)
			if !found {
				//received request IP is not blacklisted
				return next(c)
			} else {
				// received request is done using on of the blacklisted tor nodes
				//return rate limit excedeed message
				logger.Warn("drop request: provided request is done using on of the blacklisted tor nodes")
				return c.FastBlob(200, echo.MIMEApplicationJSON, data.ErrBlockTorConnection)
			}
		}
		// add keep alive headers in the response if requested by the client
		connectionMode := h.Get("Connection")
		if connectionMode != "" {
			connectionMode = str.ToLowerAscii(connectionMode)
			/*
				Lista de parámetros separados por coma,
				cada uno consiste en un identificador y un valor separado por el signo igual ('=').
				Es posible establecer los siguientes identificadores:
				* timeout: indica la cantidad de  tiempo mínima  en la cual una conexión ociosa
				se debe mantener abierta (en segundos).
				Nótese que los timeouts mas largos que el timeout de TCP
				pueden ser ignorados si no se establece un mensaje de TCP
				keep-alive  en la capa de transporte.
				* max: indica el número máximo de peticiones que pueden ser
				enviadas en esta conexión antes de que sea cerrada. Si es  0,
				este valor es ignorado para las conexiones no segmentadas,
				ya que se enviara otra solicitud en la próxima respuesta.
				Una canalización de HTTP puede ser usada para limitar la división.
			*/
			if strings.Contains(connectionMode, "keep-alive") {
				// keep alive connection mode requested
				rh.Set("Connection", "Keep-Alive")
				rh.Set("Keep-Alive", "timeout=5, max=1000")
			}
		}

		// add fake server header
		rh.Set("server", "Apache/2.0.54")
		rh.Set("x-powered-by", "PHP/5.1.6")

		// add default security headers
		// h.Set("access-control-allow-credentials", "true")
		rh.Set("x-xss-protection", "1; mode=block")
		rh.Set("strict-transport-security", "max-age=63072000; includeSubDomains; preload ") //2 years
		//public-key-pins: pin-sha256="t/OMbKSZLWdYUDmhOyUzS+ptUbrdVgb6Tv2R+EMLxJM="; pin-sha256="PvQGL6PvKOp6Nk3Y9B7npcpeL40twdPwZ4kA2IiixqA="; pin-sha256="ZyZ2XrPkTuoiLk/BR5FseiIV/diN3eWnSewbAIUMcn8="; pin-sha256="0kDINA/6eVxlkns5z2zWv2/vHhxGne/W0Sau/ypt3HY="; pin-sha256="ktYQT9vxVN4834AQmuFcGlSysT1ZJAxg+8N1NkNG/N8="; pin-sha256="rwsQi0+82AErp+MzGE7UliKxbmJ54lR/oPheQFZURy8="; max-age=600; report-uri="https://www.keycdn.com"
		rh.Set("X-Content-Type-Options", "nosniff")
		// report-uri http://reportcollector.example.com/collector.cgi
		if !config.IsDevelopment() {
			rh.Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline' 'unsafe-eval' *.etherniti.org cdnjs.cloudflare.com fonts.googleapis.com fonts.gstatic.com")
		}
		rh.Set("Expect-CT", "enforce, max-age=30")
		rh.Set("X-UA-Compatible", "IE=Edge,chrome=1")
		rh.Set("x-frame-options", "SAMEORIGIN")
		rh.Set("Referrer-Policy", "same-origin")
		rh.Set("Feature-Policy", "microphone 'none'; payment 'none'; sync-xhr 'self'")
		rh.Set("x-firefox-spdy", "h2")

		return next(c)
	}
}
