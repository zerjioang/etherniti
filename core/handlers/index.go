// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"net/http"
	"runtime"
	"time"

	"github.com/zerjioang/etherniti/core/integrity"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/server/mods/disk"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type IndexController struct {
}

const (
	indexWelcome = `{
  "name": "eth-wbapi",
  "description": "Etherniti: Ethereum Multitenant API",
  "cluster_name": "eth-wbapi",
  "version": "0.0.1",
  "env": "development",
  "tagline": "dapps everywhere"
}`
	gbUnits = float64(disk.GB)
)

var (
	//read only once, the number of server cpus
	numcpus = runtime.NumCPU()
	// monitor disk usage and get basic stats
	diskMonitor = disk.DiskUsage("/")
)

func NewIndexController() IndexController {
	dc := IndexController{}
	return dc
}

func (ctl IndexController) index(c echo.Context) error {
	return c.String(http.StatusOK, indexWelcome)
}

func (ctl IndexController) status(c echo.Context) error {

	// 1 read memory stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	wrapper := map[string]interface{}{
		// memory stats
		"alloc":       m.Alloc,
		"total":       m.TotalAlloc,
		"sys":         m.Sys,
		"mallocs":     m.Mallocs,
		"frees":       m.Frees,
		"numgc":       m.NumGC,
		"numForcedGC": m.NumForcedGC,
		"heapalloc":   m.HeapAlloc,

		// cpus stats
		"numcpus": numcpus,

		// software version stats
		"runtime":           runtime.Version(),
		"http-version":      echo.Version,
		"etherniti-version": config.Version,

		// basic disk stats
		"disk-all":  float64(diskMonitor.All) / gbUnits,
		"disk-used": float64(diskMonitor.Used) / gbUnits,
		"disk-free": float64(diskMonitor.Free) / gbUnits,
	}

	return c.JSON(http.StatusOK, wrapper)
}

func (ctl IndexController) integrity(c echo.Context) error {
	// get current date time
	currentTime := time.Now()
	timeStr := currentTime.String()

	//sign message
	hash, signature := integrity.SignMsgWithIntegrity(timeStr)
	wrapper := map[string]string{
		"message":   timeStr,
		"hash":      hash,
		"signature": signature,
	}
	return c.JSON(http.StatusOK, wrapper)
}

// implemented method from interface RouterRegistrable
func (ctl IndexController) RegisterRouters(router *echo.Echo) {
	log.Info("exposing index controller methods")
	router.GET("/v1", ctl.index)
	router.GET("/", ctl.index)
	router.GET("/v1/status", ctl.status)
	router.GET("/v1/integrity", ctl.integrity)
}
