// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package dashboard

import (
	"sync/atomic"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/bus"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared/notifier"
	"github.com/zerjioang/etherniti/thirdparty/echo"
	"github.com/zerjioang/go-bus"
)

type UIStatsController struct {
}

type ProxyStats struct {
	HandledRequest int     `json:"requests"`
	Uptime         float32 `json:"uptime"`
	Transactions   int     `json:"tx"`
	Contracts      int     `json:"contracts"`
	Accounts       int     `json:"accounts"`
	Profiles       int     `json:"profiles"`
}

func (s *ProxyStats) Defaults() {
	logger.Debug("loading proxy statistics default values")
	s.Uptime = 99.9 //current service SLA
	s.Accounts = 1  //proxy management account
	s.Profiles = 1  // proxy admin profile
}

var (
	atomicStats atomic.Value
)

func init() {
	logger.Debug("loading proxy statistics wrapper")
	var stats ProxyStats
	stats.Defaults()
	//save current data
	atomicStats.Store(stats) // atomic/thread-safe
	// subscribe to listen new http requests
	bus.Subscribe(notifier.NewProxyRequest, func(message gobus.EventMessage) {
		statsData, ok := atomicStats.Load().(ProxyStats)
		if ok {
			statsData.HandledRequest++
			atomicStats.Store(statsData)
		}
	})
	// subscribe to listen new dashboard accounts creation
	bus.Subscribe(notifier.NewDashboardAccount, func(message gobus.EventMessage) {
		statsData, ok := atomicStats.Load().(ProxyStats)
		if ok {
			statsData.Accounts++
			atomicStats.Store(statsData)
		}
	})
	// subscribe to listen new profile creation
	bus.Subscribe(notifier.NewProfileRequest, func(message gobus.EventMessage) {
		statsData, ok := atomicStats.Load().(ProxyStats)
		if ok {
			statsData.Profiles++
			atomicStats.Store(statsData)
		}
	})
}

// constructor like function
func NewUIStatsController() UIStatsController {
	uiCtl := UIStatsController{}
	return uiCtl
}

// load current proxy statistics
func (ctl UIStatsController) stats(c *echo.Context) error {
	statsData, ok := atomicStats.Load().(ProxyStats)
	if ok {
		return api.SendSuccess(c, []byte("proxy_stats"), statsData)
	}
	return api.ErrorStr(c, "failed to get proxy dashboard statistics")
}

func (ctl UIStatsController) RegisterRouters(router *echo.Group) {
	logger.Debug("exposing ui stats controller methods")
	router.GET("/stats", ctl.stats)
}
