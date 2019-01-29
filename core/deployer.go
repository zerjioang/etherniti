// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package core

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/zerjioang/gaethway/core/config"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/zerjioang/gaethway/core/api"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/zerjioang/gaethway/core/handlers"
)

var (
	userAgentErr = errors.New("not authorized. security policy not satisfied")
	gopath       = os.Getenv("GOPATH")
	resources    = gopath + "/src/github.com/zerjioang/gaethway/resources"
	corsConfig   = middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}
	gzipConfig = middleware.GzipConfig{
		Level: 5,
	}
	defaultRedirectConfig = middleware.RedirectConfig{
		Skipper: middleware.DefaultSkipper,
		Code:    http.StatusMovedPermanently,
	}
)

const (
	localhostCertPem = `-----BEGIN CERTIFICATE-----
MIIC+jCCAeKgAwIBAgIRAI4ga6WaCWzhnIgevZi02qgwDQYJKoZIhvcNAQELBQAw
EjEQMA4GA1UEChMHQWNtZSBDbzAeFw0xOTAxMjgxNjA0NDNaFw0yMDAxMjgxNjA0
NDNaMBIxEDAOBgNVBAoTB0FjbWUgQ28wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw
ggEKAoIBAQCTSjnQuVl/vC++FnPFcREbwIsrPLY8G4CQHgKqxYkR0J/NCyCSwNfC
/dGDRSE0Lun9lvUjjXBi94Ftd+r+f3okgYPrmgn7K/R5N/K+3kGvgfxUgZFXEYtK
z5wojb+pUFwTdgfT3BHp2naBFLMKI838A3Jt5MHEXJWENHs8ovchMWivlVoBjEJD
B+SaJUGD7+PC1vvGKda/P52X+sYKPrwnlze0sNdtYD1OUX4W+YntJZdr9CgznPMg
QYSZsqRr4oGiS7ONJCfxFnGvHL/WwyBfin+QXLUTbqkSa4aXPYVD5om8tk7eGL3F
eKmwaGkC5xybq7oEUCa/DlkpJgxyDTCjAgMBAAGjSzBJMA4GA1UdDwEB/wQEAwIF
oDATBgNVHSUEDDAKBggrBgEFBQcDATAMBgNVHRMBAf8EAjAAMBQGA1UdEQQNMAuC
CWxvY2FsaG9zdDANBgkqhkiG9w0BAQsFAAOCAQEAeCNiTCXCwKNkXvXZaP+xfYcs
fSB3S/UAnUxfmOBCdfyK4JCM47BA3Hz9SzLMMwrR4IP53a9hXfQxYiMffZi8R3XF
YWJTnS3giuOhe8aFH91PhPDF+sC5dlm1cd2B3i1ylv0ogbrO9ZGtO47zA41bTPiy
E8IccKiKru2bL0llj4aqg0sdHmdLMBtsjWbT/yQaveBG/bNNDk0u5IqgJWSVePwk
jFPtgDvxFkDoDAhzrJcenMSt6LtTAoBLKkWPSRC3u+iwVLacIv0pmxj+1nGW+H18
mklI/9mByeejncVBGPp5vHastJpTFyRJ4V8CRZOQ4j9fRx7sEmQ7N+9pqDNtsw==
-----END CERTIFICATE-----`
	localhostKeyPem = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAk0o50LlZf7wvvhZzxXERG8CLKzy2PBuAkB4CqsWJEdCfzQsg
ksDXwv3Rg0UhNC7p/Zb1I41wYveBbXfq/n96JIGD65oJ+yv0eTfyvt5Br4H8VIGR
VxGLSs+cKI2/qVBcE3YH09wR6dp2gRSzCiPN/ANybeTBxFyVhDR7PKL3ITFor5Va
AYxCQwfkmiVBg+/jwtb7xinWvz+dl/rGCj68J5c3tLDXbWA9TlF+FvmJ7SWXa/Qo
M5zzIEGEmbKka+KBokuzjSQn8RZxrxy/1sMgX4p/kFy1E26pEmuGlz2FQ+aJvLZO
3hi9xXipsGhpAuccm6u6BFAmvw5ZKSYMcg0wowIDAQABAoIBAHHvkRmsx1bQM/5P
T+8Dr8BQCVfA9xc4DxNso5OGiqmFQJhUazYahs0HmvJ4n17Gi6rnA2olFzL3Ut9j
TBzib5GdvnaaCe6J6et7JAQR2a3yV0bnk45Ou/l678lPHVvUFeXX/+Ya7qB/pfvk
Dztgxw6TfAkWU+2Z0O8bydj2F0VMzskrxwwF384mZgB5ysjiVcMmKI2kCJa/ovUd
B6kFMd6Y77ohlyI0jf9YWIOKeMnRcsKorFfQtzqBpTpa3purCieUMY4jh/5iYa0n
UTxmQOAoAYerNVu/d5Qayy9Y1VAW+Zl8pBOknvK+Zo49O0Dx1/vRoIEUSolczexY
CdqUVkECgYEAw1JHLGq3Cf0YjEJeSMi1IMPpx2PdGGmtznKnnxXCsYHkue/mXl8K
l58f0QmDtlwkfNRRfSLj4XsZLQk1UvzJ1aCbehwfGqEAdtCOb2A1GWlKipvIR4pr
b5NpMToJ+3jcP1cJJTh18bOUG75y5axKz2MkiG8I62dFcG0b/puLRh0CgYEAwQwN
gBUM/VpinnqZ3xC5vQeXn6k3BLu773CgSaBFcyMJedWrpKA7d+87kvn5N9onN1Ww
aJF8MckREwefKp6D0UfJYusJD5DAInYfeKPX2PT+OKncT3G4G8kq0MSjsvQks8B/
zvPsTszJKIYJqfe7KShSDIZY6GDBpCIsw8aZlb8CgYBNWrGLWrwg/ZaSPdqfUrXB
QzW73MX8XCYUg/30mCaiLEJMjUEcEOHeCIwOOolqWHWu5ltbhszfSORAnMv8kNbS
fyf0JV0AK9FGPPScEWsWJEf8OxQHmT9RUf0wHL9FU6lOgIbDseesEKXQkw1n/mMm
XSpjyi2rJRwwGVYj8LAo1QKBgQCbbbbY7xn8Sm+opZGJ9g910M0VccqodvbDu+xy
GyaPoyAYBh8idxgqYmWW2sj7XRvCA637I1fZRcgHiFVwnRwIvkG48P/klmj71htU
qKY7OlYNDUYieK8BQCDG4evjQ4rhZxYAbIhQkbVMeU8CmEEKzDnzd5/RyUVff1yH
bDlwRQKBgGxz/v8dIP5xRXXXQYre+KxXtohY7QsDJxuC3R2NMCh9lovrsqWof5OQ
xC1++6t6BnPJnMe4vdpMeuW8QTAKhHvm+XvPiPqnNeVSj7SLbOZDlivUiNZrr87t
DagBWzI58Ymmo2EJHbe48ChjOf5aeZpH7l8ZtSDbdHRFOKcUPDUJ
-----END RSA PRIVATE KEY-----`
)

var (
	localhostCert, certEtr = tls.X509KeyPair(
		[]byte(localhostCertPem),
		[]byte(localhostKeyPem),
	)
)

type Deployer struct {
}

func (deployer Deployer) GetLocalHostTLS() (tls.Certificate, error) {
	return localhostCert, certEtr
}

func (deployer Deployer) Run() {
	log.Info("loading Ethereum Multitenant Webapi (gaethway)")

	httpServerInstance := echo.New()
	httpServerInstance.HideBanner = true
	// add redirects from http to https
	httpServerInstance.Pre(deployer.httpsRedirect)

	// Start http server
	go func() {
		log.Info("starting http server...")
		err := httpServerInstance.Start(config.HttpAddress)
		if err != nil {
			log.Error("shutting down http the server")
		}
	}()

	// build a secure http server
	e := echo.New()

	// enable debug mode
	e.Debug = true

	cert, err := deployer.GetLocalHostTLS()
	if err != nil {
		log.Fatal("failed to setup TLS configuration due to error", err)
		return
	}

	//prepare tls configuration
	var tlsConf tls.Config
	tlsConf.Certificates = []tls.Certificate{cert}
	if !e.DisableHTTP2 {
		tlsConf.NextProtos = append(tlsConf.NextProtos, "h2")
	}

	//configure custom secure server
	s := &http.Server{
		Addr:         config.HttpsAddress,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		TLSConfig:    &tlsConf,
	}
	//hide the banner
	e.HideBanner = true

	// add a custom error handler
	log.Info("[LAYER] custom error handler")
	e.HTTPErrorHandler = deployer.customHTTPErrorHandler

	// antibots, crawler middleware
	// avoid bots and crawlers
	e.Pre(deployer.antiBots)

	// remove trailing slash for better usage
	log.Info("[LAYER] trailing slash remover")
	e.Pre(middleware.RemoveTrailingSlash())

	// log all single request
	// configure logging level
	log.Info("[LAYER] logger at warn level")
	e.Logger.SetLevel(log.WARN)
	e.Use(middleware.Logger())

	// add CORS support
	log.Info("[LAYER] cors support")
	e.Use(middleware.CORSWithConfig(corsConfig))

	// add server api request hardening using http headers
	e.Use(deployer.hardening)

	// add fake server header
	e.Use(deployer.fakeServer)

	// Request ID middleware generates a unique id for a request.
	e.Use(middleware.RequestID())

	// add gzip support if client requests it
	log.Info("[LAYER] gzip compression")
	e.Use(middleware.GzipWithConfig(gzipConfig))

	// avoid panics
	e.Use(middleware.Recover())

	//load root static folder
	e.Static("/", resources+"/root")

	// load swagger ui files
	e.Static("/swagger", resources+"/swagger")

	deployer.register(e)

	// Start secure server
	go func() {
		log.Info("starting https secure server...")
		err := e.StartServer(s)
		if err != nil {
			e.Logger.Info("shutting down https secure the server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Info("graceful shutdown of the service requested")
	log.Info("shutting down http server...")
	if err := httpServerInstance.Shutdown(ctx); err != nil {
		log.Error(err)
	}
	log.Info("shutting down https secure server...")
	if err := e.Shutdown(ctx); err != nil {
		log.Error(err)
	}
	log.Info("graceful shutdown executed")
	log.Info("exiting...")
}

// http to http redirect function
func (deployer Deployer) httpsRedirect(next echo.HandlerFunc) echo.HandlerFunc {
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
func (deployer Deployer) hardening(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add security headers
		h := c.Response().Header()
		h.Set("server", "Apache")
		h.Set("access-control-allow-credentials", "true")
		h.Set("x-xss-protection", "1; mode=block")
		h.Set("strict-transport-security", "max-age=31536000; includeSubDomains; preload")
		//public-key-pins: pin-sha256="t/OMbKSZLWdYUDmhOyUzS+ptUbrdVgb6Tv2R+EMLxJM="; pin-sha256="PvQGL6PvKOp6Nk3Y9B7npcpeL40twdPwZ4kA2IiixqA="; pin-sha256="ZyZ2XrPkTuoiLk/BR5FseiIV/diN3eWnSewbAIUMcn8="; pin-sha256="0kDINA/6eVxlkns5z2zWv2/vHhxGne/W0Sau/ypt3HY="; pin-sha256="ktYQT9vxVN4834AQmuFcGlSysT1ZJAxg+8N1NkNG/N8="; pin-sha256="rwsQi0+82AErp+MzGE7UliKxbmJ54lR/oPheQFZURy8="; max-age=600; report-uri="https://www.keycdn.com"
		h.Set("X-Content-Type-Options", "nosniff")
		h.Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline'")
		h.Set("Expect-CT", "enforce, max-age=30")
		h.Set("X-UA-Compatible", "IE=Edge,chrome=1")
		h.Set("x-frame-options", "SAMEORIGIN")
		h.Set("Referrer-Policy", "same-origin")
		h.Set("Feature-Policy", "microphone 'none'; payment 'none'; sync-xhr 'self'")
		h.Set("x-firefox-spdy", "h2")
		h.Set("x-powered-by", "PHP/5.6.38")
		return next(c)
	}
}

// fakeServer middleware function.
func (deployer Deployer) fakeServer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add fake server header
		h := c.Response().Header()
		h.Set("server", "Apache")
		h.Set("x-powered-by", "PHP/5.6.38")
		return next(c)
	}
}

// fakeServer antiBots function.
func (deployer Deployer) antiBots(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add antibots policy
		ua := c.Request().UserAgent()
		if ua == "" || deployer.isBotRequest(ua) {
			//drop the request
			return userAgentErr
		}
		return next(c)
	}
}

func (deployer Deployer) isBotRequest(userAgent string) bool {
	var lock = false
	for i:=0; i <len(api.BadBotsList) && !lock ;i++ {
		lock = strings.Contains(userAgent, api.BadBotsList[i])
	}
	return lock
}

// keepalive middleware function.
func (deployer Deployer) keepalive(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add keep alive headers
		c.Response().Header().Set("Connection", "Keep-Alive")
		c.Response().Header().Set("Keep-Alive", "timeout=5, max=1000")
		return next(c)
	}
}

func (deployer Deployer) customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	_ = c.JSON(code, api.NewApiError(
		code, err.Error()),
	)
}

// register in echo server, allowed routes
func (deployer Deployer) register(server *echo.Echo) *echo.Echo {
	log.Info("registering routes")
	handlers.NewIndexController().RegisterRouters(server)
	handlers.NewProfileController().RegisterRouters(server)
	handlers.NewEthController().RegisterRouters(server)
	return server
}

func NewDeployer() Deployer {
	d := Deployer{}
	return d
}
