// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"crypto/tls"
	"errors"
	"os"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	userAgentErr = errors.New("not authorized. security policy not satisfied")
	gopath       = os.Getenv("GOPATH")
	resources    = gopath + "/src/github.com/zerjioang/etherniti/resources"
	corsConfig   = middleware.CORSConfig{
		AllowOrigins: config.AllowedCorsOriginList,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			"X-Language",
			config.HttpProfileHeaderkey,
		},
	}
	accessLogFormat = `{"time":"${time_unix}","id":"${id}","ip":"${remote_ip}",` +
		`"host":"${host}","method":"${method}","referer":"${referer}","uri":"${uri}","ua":"${user_agent}",` +
		`"status":${status},"err":"${trycatch}","latency":${latency},"latency_human":"${latency_human}"` +
		`,"in":${bytes_in},"out":${bytes_out}}` + "\n"
	gzipConfig = middleware.GzipConfig{
		Level: 5,
	}
	localhostCert tls.Certificate
	certEtr       error
)

func recoverName() {
	if r := recover(); r != nil {
		logger.Info("recovered from ", r)
	}
}
