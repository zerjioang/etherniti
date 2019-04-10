// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package http

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	constants2 "github.com/zerjioang/etherniti/shared/constants"

	"github.com/zerjioang/etherniti/core/util/banner"

	"github.com/zerjioang/etherniti/core/listener/base"
	"github.com/zerjioang/etherniti/shared/constants"
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
)

type HttpListener struct {
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

func (l HttpListener) GetLocalHostTLS() (tls.Certificate, error) {
	return localhostCert, certEtr
}

func (l HttpListener) RunMode(address string, background bool) {
}

func (l HttpListener) Listen() error {
	logger.Info("loading Etherniti Proxy, an Ethereum Multitenant WebAPI")
	if config.EnableHttpsRedirect {
		//build http server
		httpServerInstance := base.NewServer(nil)
		// add redirects from http to https
		logger.Info("[LAYER] http to https redirect")
		httpServerInstance.Pre(base.HttpsRedirect)

		// Start http server
		go func() {
			s, err := l.buildInsecureServerConfig()
			if err != nil {
				logger.Error("failed to build http server configuration", err)
			} else {
				logger.Info("starting http server...")
				println(banner.WelcomeBanner())
				err := httpServerInstance.StartServer(s)
				if err != nil {
					logger.Error("shutting down http the server", err)
				}
			}
		}()
		// Start https server
		secureServer := base.NewServer(base.ConfigureServerRoutes)
		go func() {
			s, err := l.buildSecureServerConfig(secureServer)
			if err != nil {
				logger.Error("failed to build https server configuration", err)
			} else {
				logger.Info("starting https server...")
				configureSwaggerJson()
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
		e := base.NewServer(base.ConfigureServerRoutes)
		s, err := l.buildInsecureServerConfig()
		if err != nil {
			logger.Error("failed to build server configuration", err)
		} else {
			// Start server
			go func() {
				logger.Info("starting http server...")
				configureSwaggerJson()
				println(banner.WelcomeBanner())
				err := e.StartServer(s)
				if err != nil {
					logger.Info("shutting down http server", err)
				}
			}()
			//graceful shutdown of http server
			l.shutdown(e, nil)
		}
	}
	return nil
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
	configureSwaggerJsonWithDir(config.ResourcesDir)
}

func configureSwaggerJsonWithDir(resources string) {
	//read template file
	log.Debug("reading swagger json file")
	raw, err := ioutil.ReadFile(resources + "/swagger/swagger-template.json")
	if err != nil {
		logger.Error("failed reading swagger template file", err)
		return
	}
	//replace hardcoded variables
	str := string(raw)
	str = strings.Replace(str, "$title", "Etherniti REST API Proxy", -1)
	str = strings.Replace(str, "$version", constants2.Version, -1)
	str = strings.Replace(str, "$host", config.SwaggerAddress, -1)
	str = strings.Replace(str, "$basepath", "/v1", -1)
	str = strings.Replace(str, "$header-auth-key", constants.HttpProfileHeaderkey, -1)
	//write swagger.json file
	writeErr := ioutil.WriteFile(resources+"/swagger/swagger.json", []byte(str), os.ModePerm)
	if writeErr != nil {
		logger.Error("failed writing swagger.json file", writeErr)
		return
	}
}

// create new deployer instance
func NewHttpListener() listener.ListenerInterface {
	d := HttpListener{}
	d.limiter = ratelimit.NewRateLimitEngine()
	return d
}
