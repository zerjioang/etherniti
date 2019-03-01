// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"context"
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/release"
	"github.com/zerjioang/etherniti/core/server/mods/ratelimit"

	"github.com/zerjioang/etherniti/core/eth"

	"github.com/zerjioang/etherniti/core/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
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

type Deployer struct {
	manager eth.WalletManager
	limiter ratelimit.RateLimitEngine
}

func (deployer Deployer) GetLocalHostTLS() (tls.Certificate, error) {
	return localhostCert, certEtr
}

func (deployer Deployer) Run() {
	logger.Info("loading Etherniti Proxy, an Ethereum Multitenant WebAPI")
	if config.EnableHttpsRedirect {
		//build http server
		httpServerInstance := deployer.newServerInstance()
		// add redirects from http to https
		logger.Info("[LAYER] http to https redirect")
		httpServerInstance.Pre(httpsRedirect)

		// Start http server
		go func() {
			s, err := deployer.buildInsecureServerConfig(httpServerInstance)
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
		secureServer := deployer.newServerInstance()
		go func() {
			s, err := deployer.buildSecureServerConfig(secureServer)
			if err != nil {
				logger.Error("failed to build https server configuration", err)
			} else {
				logger.Info("starting https server...")
				ConfigureRoutes(secureServer)
				err := secureServer.StartServer(s)
				if err != nil {
					logger.Error("shutting down https the server", err)
				}
			}
		}()
		//graceful shutdown of http and https server
		deployer.shutdown(httpServerInstance, secureServer)
	} else {
		//deploy http server only
		e := deployer.newServerInstance()
		s, err := deployer.buildInsecureServerConfig(e)
		if err != nil {
			logger.Error("failed to build server configuration", err)
		} else {
			ConfigureRoutes(e)
			// Start server
			go func() {
				logger.Info("starting http server...")
				err := e.StartServer(s)
				if err != nil {
					logger.Info("shutting down http server", err)
				}
			}()
			//graceful shutdown of http server
			deployer.shutdown(e, nil)
		}
	}
}

func (deployer Deployer) shutdown(httpInstance *echo.Echo, httpsInstance *echo.Echo) {
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

func (deployer Deployer) buildSecureServerConfig(e *echo.Echo) (*http.Server, error) {
	cert, err := deployer.GetLocalHostTLS()
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

func (deployer Deployer) buildInsecureServerConfig(e *echo.Echo) (*http.Server, error) {
	//configure custom secure server
	return &http.Server{
		Addr:         config.ListeningAddress,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}, nil
}

// build new http server instance
func (deployer Deployer) newServerInstance() *echo.Echo {
	// build a the server
	e := echo.New()
	// enable debug mode
	e.Debug = config.DebugServer
	e.HidePort = config.HideServerDataInConsole
	//hide the banner
	e.HideBanner = config.HideServerDataInConsole
	return e
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
func NewDeployer() Deployer {
	d := Deployer{}
	d.manager = eth.NewWalletManager()
	d.limiter = ratelimit.NewRateLimitEngine()
	return d
}
