// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func init() {
	// register all prometheus metrics collectors
	registerMetrics()
}

func registerMetrics() {
	// Prometheus: Histogram to collect required metrics
	histogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "greeting_seconds",
		Help:    "Time take to greet someone",
		Buckets: []float64{1, 2, 5, 6, 10}, //defining small buckets as this app should not take more than 1 sec to respond
	}, []string{"code"}) // this will be partitioned by the HTTP code.

	//Registering the defined metric with Prometheus
	regErr := prometheus.Register(histogram)
	if regErr != nil {
		logger.Error("failed to register histogram metrics: ", regErr.Error())
	}
}

func MetricsCollector(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		// todo update all required manual prometheus metrics here
		return next(c)
	}
}
