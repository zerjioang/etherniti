// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/eth"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/release"
	"github.com/zerjioang/etherniti/core/server/mods/ratelimit"
)

func init() {
	defer recoverName()
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

type HttpListener struct {
	manager eth.WalletManager
	limiter ratelimit.RateLimitEngine
}

func (l HttpListener) GetLocalHostTLS() (tls.Certificate, error) {
	return localhostCert, certEtr
}

func (l HttpListener) Run() {
	logger.Info("loading Etherniti Proxy, an Ethereum Multitenant WebAPI")
	if config.EnableHttpsRedirect {
		//build http server
		httpServerInstance := NewServer()
		// add redirects from http to https
		logger.Info("[LAYER] http to https redirect")
		httpServerInstance.Pre(httpsRedirect)

		// Start http server
		go func() {
			s, err := l.buildInsecureServerConfig()
			if err != nil {
				logger.Error("failed to build http server configuration", err)
			} else {
				logger.Info("starting http server...")
				err := httpServerInstance.StartServer(s)
				if err != nil {
					logger.Error("shutting down http the server", err)
				}
			}
		}()
		// Start https server
		secureServer := NewServer()
		go func() {
			s, err := l.buildSecureServerConfig(secureServer)
			if err != nil {
				logger.Error("failed to build https server configuration", err)
			} else {
				logger.Info("starting https server...")
				ConfigureServerRoutes(secureServer)
				err := secureServer.StartServer(s)
				if err != nil {
					logger.Error("shutting down https the server", err)
				}
			}
		}()
		//graceful shutdown of http and https server
		l.shutdown(httpServerInstance, secureServer)
	} else {
		//deploy http server only
		e := NewServer()
		s, err := l.buildInsecureServerConfig()
		if err != nil {
			logger.Error("failed to build server configuration", err)
		} else {
			ConfigureServerRoutes(e)
			// Start server
			go func() {
				logger.Info("starting http server...")
				err := e.StartServer(s)
				if err != nil {
					logger.Info("shutting down http server", err)
				}
			}()
			//graceful shutdown of http server
			l.shutdown(e, nil)
		}
	}
}

func (l HttpListener) shutdown(httpInstance *echo.Echo, httpsInstance *echo.Echo) {
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

func (l HttpListener) buildSecureServerConfig(e *echo.Echo) (*http.Server, error) {
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

func (l HttpListener) buildInsecureServerConfig() (*http.Server, error) {
	//configure custom secure server
	return &http.Server{
		Addr:         config.ListeningAddress,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}, nil
}

func configureSwaggerJson() {
	//read template file
	log.Debug("reading swagger json file")
	raw, err := ioutil.ReadFile(resources + "/swagger/swagger-template.json")
	if err != nil {
		logger.Error("failed reading swagger template file", err)
		return
	}
	//replace hardcoded variables
	str := string(raw)
	str = strings.Replace(str, "$title", "Etherniti Proxy REST API", -1)
	str = strings.Replace(str, "$version", release.Version, -1)
	str = strings.Replace(str, "$host", config.SwaggerAddress, -1)
	str = strings.Replace(str, "$basepath", "/v1", -1)
	str = strings.Replace(str, "$header-auth-key", config.HttpProfileHeaderkey, -1)
	//write swagger.json file
	writeErr := ioutil.WriteFile(resources+"/swagger/swagger.json", []byte(str), os.ModePerm)
	if writeErr != nil {
		logger.Error("failed writing swagger.json file", writeErr)
		return
	}
}

// create new deployer instance
func NewHttpListener() HttpListener {
	d := HttpListener{}
	d.manager = eth.NewWalletManager()
	d.limiter = ratelimit.NewRateLimitEngine()
	return d
}
