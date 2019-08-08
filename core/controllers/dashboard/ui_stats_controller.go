// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package dashboard

import (
	"github.com/zerjioang/etherniti/core/bench"
	"sync/atomic"
	"time"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/bus"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared/notifier"
	"github.com/zerjioang/etherniti/thirdparty/echo"
	"github.com/zerjioang/go-bus"
)

type ProxyStatsController struct {
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
	scoreWrapper struct {
		Time  time.Duration `json:"time"`
		Score int64         `json:"score"`
	}
)

func init() {
	logger.Debug("loading proxy statistics wrapper")
	//load internal benchmarking data
	scoreWrapper.Score = bench.GetScore()
	scoreWrapper.Time = bench.GetBenchTime()

	// initialize proxy statistics wrapper
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
func NewProxyStatsController() ProxyStatsController {
	uiCtl := ProxyStatsController{}
	return uiCtl
}

// load current proxy statistics
func (ctl ProxyStatsController) internalAnalytics(c *echo.Context) error {
	statsData, ok := atomicStats.Load().(ProxyStats)
	if ok {
		return api.SendSuccess(c, []byte("proxy_internal_analytics"), statsData)
	}
	return api.ErrorStr(c, "failed to get proxy dashboard statistics")
}

func (ctl ProxyStatsController) score(c *echo.Context) error {
	return api.SendSuccess(c, []byte("proxy_score"), scoreWrapper)
}

func (ctl ProxyStatsController) RegisterRouters(router *echo.Group) {
	logger.Debug("exposing ui internalAnalytics controller methods")
	router.GET("/analytics", ctl.internalAnalytics)
	router.GET("/score", ctl.score)
}
