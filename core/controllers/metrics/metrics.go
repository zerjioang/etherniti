// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package metrics

import (
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/controllers/wrap"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared"
	"github.com/zerjioang/go-hpc/lib/counter32"
	"github.com/zerjioang/go-hpc/lib/fastime"

	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

type MetricsController struct {
	// use channels: https://talks.golang.org/2012/concurrency.slide#25
}

var (
	startTime      fastime.FastTime
	requestCounter counter32.Count32
)

func init() {
	startTime = fastime.Now()
}

func NewMetricsController() *MetricsController {
	return new(MetricsController)
}

func (ctl *MetricsController) Uptime(c *shared.EthernitiContext) error {
	c.OnSuccessCachePolicy = 5
	return api.SendSuccessBlob(c, ctl.uptime())
}

func (ctl *MetricsController) uptime() []byte {
	return startTime.SafeBytes()
}

func (ctl *MetricsController) Requests(c *shared.EthernitiContext) error {
	c.OnSuccessCachePolicy = 5
	return api.SendSuccessBlob(c, ctl.requests())
}

func (ctl *MetricsController) requests() []byte {
	return requestCounter.UnsafeBytes()
}

// implemented method from interface RouterRegistrable
func (ctl *MetricsController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing metrics controller methods")
	router.GET("/metrics/uptime", wrap.Call(ctl.Uptime))
	router.GET("/metrics/requests", wrap.Call(ctl.Requests))
}
