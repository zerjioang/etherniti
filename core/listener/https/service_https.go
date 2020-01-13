// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package https

import (
	"crypto/tls"
	"net/http"

	base "github.com/zerjioang/etherniti/core/listener/http"

	"github.com/zerjioang/etherniti/core/listener/middleware"

	"github.com/zerjioang/etherniti/core/listener/swagger"

	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/shared/def/listener"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
	"github.com/zerjioang/go-hpc/thirdparty/gommon/log"
)

var (
	//default etherniti proxy configuration
	cfg = config.GetDefaultOpts()
	//variables used when HTTPS is requested
	proxyCertificate tls.Certificate
	certErr          error
)

type HttpsListener struct {
	base.HttpListener
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
		proxyCertificate, certErr = tls.X509KeyPair(
			certBytes,
			keyBytes,
		)
	} else {
		logger.Error("failed to load SSL crypto data")
	}
}

func (l HttpsListener) GetLocalHostTLS() (tls.Certificate, error) {
	return proxyCertificate, certErr
}

//fetch specific server configuration
func (l HttpsListener) ServerConfig() *http.Server {
	return common.DefaultHttpServerConfig
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
		err := httpServerInstance.StartServer(common.DefaultHttpServerConfig)
		if err != nil {
			logger.Error("shutting down http the server: ", err)
			notifier <- err
		}
	}()

	// Start https server
	secureServer := common.NewServer(middleware.ConfigureServerRoutes)
	go func() {
		s, err := l.buildServerConfig(secureServer)
		if err != nil {
			logger.Error("failed to build https server configuration: ", err)
			notifier <- err
		} else {
			logger.Info("starting https server...")
			swagger.ConfigureFromTemplate()
			err := secureServer.StartServer(s)
			if err != nil {
				logger.Error("shutting down https the server: ", err)
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
		log.Fatal("failed to setup TLS configuration due to error: ", err)
		return nil, err
	}

	// prepare tls configuration
	// and get a perfect SSL Labs Score
	tlsConf := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
		Certificates: []tls.Certificate{cert},
	}
	if !e.DisableHTTP2 {
		tlsConf.NextProtos = append(tlsConf.NextProtos, "h2")
	}

	//configure custom secure server
	// which is default http config + tls data
	secureServerConfig := common.DefaultHttpsServerConfig
	secureServerConfig.TLSConfig = tlsConf
	return secureServerConfig, nil
}

// create new deployer instance
func NewHttpsListenerCustom() HttpsListener {
	d := HttpsListener{}
	return d
}

// create new deployer instance
func NewHttpsListener() listener.ListenerInterface {
	d := HttpsListener{}
	d.HttpListener = base.NewHttpListenerCustom()
	return d
}
