// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package providers

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/controllers/dashboard"
	"github.com/zerjioang/etherniti/core/controllers/index"
	"github.com/zerjioang/etherniti/core/controllers/profiles"
	"github.com/zerjioang/etherniti/core/controllers/project"
	"github.com/zerjioang/etherniti/core/controllers/registry"
	"github.com/zerjioang/etherniti/core/controllers/security"
	"github.com/zerjioang/etherniti/core/controllers/solc"
	"github.com/zerjioang/etherniti/core/controllers/thirdparty"
	"github.com/zerjioang/etherniti/core/controllers/wallet"
	"github.com/zerjioang/etherniti/core/controllers/wrap"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/go-hpc/lib/codes"
	"github.com/zerjioang/go-hpc/lib/eth/rpc/client"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

var (
	cfg = config.GetDefaultOpts()
)

func infuraJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return defaultJwt(next, data.InfuraJwtErrorMessage)
}

func quiknodeJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return defaultJwt(next, data.QuiknodeJwtErrorMessage)
}

func ganacheJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func privateJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return defaultJwt(next, data.JwtErrorMessage)
}

func userJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc, ok := c.(*shared.EthernitiContext)
		if ok && cc != nil {
			tokenData := cc.ReadToken("Authorization")
			if tokenData == "" {
				logger.Error("missing authentication token")
				return api.ErrorWithMessage(cc, codes.StatusUnauthorized, data.ErrMissingAuthenticationToken, data.ErrMissingAuthentication)
			}
			decodedAuthData, err := dashboard.ParseAuthenticationToken(tokenData)
			if err != nil {
				logger.Error("failed to process authentication token: ", err)
				return api.ErrorWithMessage(cc, codes.StatusUnauthorized, data.ErrInvalidAuthenticationToken, err)
			} else {
				// todo review this
				cc.SetTokenData(decodedAuthData.Uuid)
			}
		}
		return next(c)
	}
}

// jwt middleware function.
func defaultJwt(next echo.HandlerFunc, errorMsg []byte) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc, ok := c.(*shared.EthernitiContext)
		if ok && cc != nil {
			_, parseErr := cc.ConnectionProfileSetup()
			if parseErr != nil {
				return api.ErrorBytes(cc, errorMsg)
			}
		}
		return next(c)
	}
}

// create a group for all /api/v1 functions
func next(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

// RegisterServices in echo server, allowed routes
func RegisterServices(e *echo.Echo) *echo.Group {

	// create a shared instance of http client for all controllers
	cli := client.NewEthClient()

	// /v1
	groupV1 := e.Group(constants.ApiVersion, next)
	if cfg.MetricsEnabled {
		logger.Info("registering prometheus_metrics metrics collector endpoint")
		e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	}
	// /v1/hi
	indexCtl := index.NewIndexController()
	groupV1.GET("/hi", wrap.Call(indexCtl.Index))

	// add authentication controller to root
	// /v1/auth/...
	// register ui rest
	dashboard.NewUIAuthController().RegisterRouters(groupV1)

	// add api key based authentication mechanism
	dashboard.NewApiKeysAuthController().RegisterRouters(groupV1)

	// register dashboard required apis
	// /dashboard/stats
	internalGroup := groupV1.Group("/internal", next)
	dashboard.NewProxyStatsController().RegisterRouters(internalGroup)

	// /v1/...
	indexCtl.RegisterRouters(groupV1)
	profiles.NewProfileController().RegisterRouters(groupV1)
	security.NewSecurityController().RegisterRouters(groupV1)
	wallet.NewWalletController().RegisterRouters(groupV1)
	solc.NewSolcController().RegisterRouters(groupV1)

	//register external api calls

	// ui helper calls
	uiGroup := groupV1.Group("/ui", next)
	dashboard.NewUIController().RegisterRouters(uiGroup)

	// coin market cap: get eth price data
	externalGroup := groupV1.Group("/external", next)
	thirdparty.NewExternalController(cli).RegisterRouters(externalGroup)

	//register web3 networks related services
	web3Group := groupV1.Group("/web3", next)

	// /v1/web3/ropsten
	ropstenGroup := web3Group.Group("/ropsten", next)
	NewRopstenController(cli).RegisterRouters(ropstenGroup)

	// /v1/web3/rinkeby
	rinkebyGroup := web3Group.Group("/rinkeby", next)
	NewRinkebyController(cli).RegisterRouters(rinkebyGroup)

	// /v1/web3/kovan
	kovanGroup := web3Group.Group("/kovan", next)
	NewKovanController(cli).RegisterRouters(kovanGroup)

	// /v1/web3/mainnet
	mainnetGroup := web3Group.Group("/mainnet", next)
	NewMainNetController(cli).RegisterRouters(mainnetGroup)

	// /v1/web3/infura
	infuraGroup := web3Group.Group("/infura", infuraJwt)
	NewInfuraController(cli).RegisterRouters(infuraGroup)

	// /v1/web3/quiknode
	quiknodeGroup := web3Group.Group("/quiknode", quiknodeJwt)
	NewQuikNodeController(cli).RegisterRouters(quiknodeGroup)

	// /v1/web3/ganache
	ganacheGroup := web3Group.Group("/ganache", ganacheJwt)
	NewGanacheController(cli).RegisterRouters(ganacheGroup)

	// /v1/web3/private
	privateGroup := web3Group.Group(constants.PrivateApi, privateJwt)
	NewPrivateNetController(cli).RegisterRouters(privateGroup)

	// register controllers related to user context
	// /v1/my
	userGroup := groupV1.Group("/my", userJwt)
	p := project.NewProjectControllerPtr()
	p.RegisterRouters(userGroup)
	registry.NewRegistryController().RegisterRouters(userGroup)

	// register project interaction controller
	// /v1/dapp
	dappGroup := groupV1.Group("/dapp", userJwt)
	project.NewProjectInteractionControllerPtr(p, cli).RegisterRouters(dappGroup)
	return groupV1
}
