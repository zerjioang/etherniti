// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package https

import (
	"context"
	"crypto/tls"
	"github.com/zerjioang/etherniti/core/listener/middleware"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/zerjioang/etherniti/core/listener/swagger"

	"github.com/zerjioang/etherniti/core/util/banner"

	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/shared/def/listener"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/server/mods/ratelimit"
	"github.com/zerjioang/etherniti/thirdparty/echo"
	"github.com/zerjioang/etherniti/thirdparty/gommon/log"
)

var (
	//variables used when HTTPS is requested
	localhostCert tls.Certificate
	certEtr       error
	// define http server config for http to https redirection
	defaultHttpServerConfig = http.Server{
		Addr:         config.ListeningAddress,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
)

type HttpsListener struct {
	limiter ratelimit.RateLimitEngine
}

func recoverFromPem() {
	if r := recover(); r != nil {
		logger.Info("recovered from pem", r)
	}
}

func init() {
	defer recoverFromPem()
	certBytes := config.GetCertPem()
	keyBytes := config.GetKeyPem()
	if certBytes != nil && len(certBytes) > 0 &&
		keyBytes != nil && len(keyBytes) > 0 {
		localhostCert, certEtr = tls.X509KeyPair(
			certBytes,
			keyBytes,
		)
	} else {
		logger.Error("failed to load SSL crypto data")
	}
}

func (l HttpsListener) GetLocalHostTLS() (tls.Certificate, error) {
	return localhostCert, certEtr
}

func (l HttpsListener) RunMode(address string, background bool) {
}

func (l HttpsListener) Listen() error {
	logger.Info("loading Etherniti Proxy, an Ethereum Multitenant WebAPI")
	//build http server
	httpServerInstance := common.NewServer(nil)
	// add redirects from http to https
	logger.Info("[LAYER] http to https redirect")
	httpServerInstance.Pre(middleware.HttpsRedirect)

	// Start http server
	go func() {
		logger.Info("starting http server...")
		println(banner.WelcomeBanner())
		err := httpServerInstance.StartServer(&defaultHttpServerConfig)
		if err != nil {
			logger.Error("shutting down http the server", err)
		}
	}()
	// Start https server
	secureServer := common.NewServer(middleware.ConfigureServerRoutes)
	go func() {
		s, err := l.buildServerConfig(secureServer)
		if err != nil {
			logger.Error("failed to build https server configuration", err)
		} else {
			logger.Info("starting https server...")
			swagger.ConfigureFromTemplate()
			err := secureServer.StartServer(s)
			if err != nil {
				logger.Error("shutting down https the server", err)
			}
		}
	}()
	//graceful shutdown of http and https server
	l.shutdown(httpServerInstance, secureServer)
	return nil
}

func (l HttpsListener) shutdown(httpInstance *echo.Echo, httpsInstance *echo.Echo) {
	// The make built-in returns a value of type T (not *T), and it's memory is
	// initialized.
	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	logger.Info("graceful shutdown of the service requested")
	if httpInstance != nil {
		logger.Info("shutting down http server...")
		if err := httpInstance.Shutdown(ctx); err != nil {
			logger.Error(err)
		}
	}
	if httpsInstance != nil {
		logger.Info("shutting down https secure server...")
		if err := httpsInstance.Shutdown(ctx); err != nil {
			logger.Error(err)
		}
	}
	logger.Info("graceful shutdown executed")
	logger.Info("exiting...")
}

func (l HttpsListener) buildServerConfig(e *echo.Echo) (*http.Server, error) {
	cert, err := l.GetLocalHostTLS()
	if err != nil {
		log.Fatal("failed to setup TLS configuration due to trycatch", err)
		return nil, err
	}

	//prepare tls configuration
	var tlsConf tls.Config
	tlsConf.Certificates = []tls.Certificate{cert}
	if !e.DisableHTTP2 {
		tlsConf.NextProtos = append(tlsConf.NextProtos, "h2")
	}

	//configure custom secure server
	return &http.Server{
		Addr:         config.ListeningAddress,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		TLSConfig:    &tlsConf,
	}, nil
}

// create new deployer instance
func NewHttpsListener() listener.ListenerInterface {
	d := HttpsListener{}
	d.limiter = ratelimit.NewRateLimitEngine()
	return d
}
