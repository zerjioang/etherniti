// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package https

import (
	"crypto/tls"
	"net/http"

	http2 "github.com/zerjioang/etherniti/core/listener/http"

	"github.com/zerjioang/etherniti/core/listener/middleware"

	"github.com/zerjioang/etherniti/core/listener/swagger"

	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/shared/def/listener"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
	"github.com/zerjioang/etherniti/thirdparty/gommon/log"
)

var (
	//default etherniti proxy configuration
	cfg = config.GetDefaultOpts()
	//variables used when HTTPS is requested
	localhostCert tls.Certificate
	certEtr       error
)

type HttpsListener struct {
	http2.HttpListener
}

func recoverFromPem() {
	if r := recover(); r != nil {
		logger.Error("recovered from pem", r)
	}
}

func init() {
	defer recoverFromPem()
	certBytes := cfg.GetCertPem()
	keyBytes := cfg.GetKeyPem()
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

//fetch specific server configuration
func (l HttpsListener) ServerConfig() *http.Server {
	return &common.DefaultHttpServerConfig
}

func (l HttpsListener) Listen(notifier chan error) {
	logger.Info("loading Etherniti Proxy, a High Performance Web3 Multitenant REST Proxy")
	logger.Info("loading https listener")
	//build http server
	httpServerInstance := common.NewServer(nil)
	// add redirects from http to https
	logger.Info("[LAYER] http to https redirect")
	httpServerInstance.Pre(middleware.HttpsRedirect)

	// Start http server
	go func() {
		logger.Info("starting http server...")
		err := httpServerInstance.StartServer(&common.DefaultHttpServerConfig)
		if err != nil {
			logger.Error("shutting down http the server", err)
			notifier <- err
		}
	}()

	// Start https server
	secureServer := common.NewServer(middleware.ConfigureServerRoutes)
	go func() {
		s, err := l.buildServerConfig(secureServer)
		if err != nil {
			logger.Error("failed to build https server configuration", err)
			notifier <- err
		} else {
			logger.Info("starting https server...")
			swagger.ConfigureFromTemplate()
			err := secureServer.StartServer(s)
			if err != nil {
				logger.Error("shutting down https the server", err)
				notifier <- err
			}
		}
	}()
	//graceful shutdown of http
	l.ShutdownListener("http", httpServerInstance, notifier)
	//graceful shutdown of https
	l.ShutdownListener("https", secureServer, notifier)
}

func (l HttpsListener) buildServerConfig(e *echo.Echo) (*http.Server, error) {
	cert, err := l.GetLocalHostTLS()
	if err != nil {
		log.Fatal("failed to setup TLS configuration due to stack", err)
		return nil, err
	}

	//prepare tls configuration
	var tlsConf tls.Config
	tlsConf.Certificates = []tls.Certificate{cert}
	if !e.DisableHTTP2 {
		tlsConf.NextProtos = append(tlsConf.NextProtos, "h2")
	}

	//configure custom secure server
	// which is default http config + tls data
	secureServerConfig := common.DefaultHttpServerConfig
	secureServerConfig.TLSConfig = &tlsConf
	return &secureServerConfig, nil
}

// create new deployer instance
func NewHttpsListenerCustom() HttpsListener {
	d := HttpsListener{}
	return d
}

// create new deployer instance
func NewHttpsListener() listener.ListenerInterface {
	d := HttpsListener{}
	d.HttpListener = http2.NewHttpListenerCustom()
	return d
}
