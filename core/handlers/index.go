// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/release"
	"net/http"
	"runtime"
	"time"

	"github.com/zerjioang/etherniti/core/integrity"

	"github.com/zerjioang/etherniti/core/server/mods/disk"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type IndexController struct {
}

const (
	gbUnits = float64(disk.GB)
)

var (
	//read only once, the number of server cpus
	numcpus = runtime.NumCPU()
	// monitor disk usage and get basic stats
	diskMonitor = disk.DiskUsagePtr()
)

// index data
const (
	indexWelcomeJson = `{
  "name": "eth-wbapi",
  "description": "Etherniti: Ethereum REST API",
  "cluster_name": "eth-wbapi",
  "version": "` + release.Version + `",
  "env": "` + config.EnvironmentName + `",
  "tagline": "dapps everywhere"
}`
	indexWelcomeHtml = `<!doctype html>
<title>Etherniti: Ethereum Multitenant API</title>
<style>
  body { text-align: center; padding: 150px; }
  h1 { font-size: 50px; }
  body { font: 20px Helvetica, sans-serif; color: #333; }
  article { display: block; text-align: left; width: 800px; margin: 0 auto; }
  a { color: #dc8100; text-decoration: none; }
  a:hover { color: #333; text-decoration: none; }
</style>

<article>
    <h1 style="text-align:center">Etherniti: an Ethereum REST API</h1>
	<h2 style="text-align:center">a High Performance, Multitenant RESTful proxy service for Ethereum</h2>
    <div>
        <p>Please refer to official API documentation for further details or visit <a href="http://dev-proxy.etherniti.org/swagger">http://dev-proxy.etherniti.org/swagger</a></p>
        <p>&mdash; Etherniti core team</p>
    </div>
</article>
	`
)

var (
	//bytes of welcome message
	indexWelcomeBytes = []byte(indexWelcomeJson)
)

func init() {
	var monErr error
	monErr = diskMonitor.Eval("/")
	if monErr != nil {
		log.Error("failed to start disk status monitor on path /. Caused by: ", monErr)
	}
}

func NewIndexController() IndexController {
	dc := IndexController{}
	return dc
}

func Index(c echo.Context) error {
	if c.Request().Header.Get("Accept") == "application/json" {
		return c.JSONBlob(http.StatusOK, indexWelcomeBytes)
	}
	return c.HTML(http.StatusOK, indexWelcomeHtml)
}

func (ctl IndexController) status(c echo.Context) error {

	// read memory stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	wrapper := map[string]interface{}{
		// memory stats
		"memory": map[string]uint64{
			"alloc":     m.Alloc,
			"total":     m.TotalAlloc,
			"sys":       m.Sys,
			"mallocs":   m.Mallocs,
			"frees":     m.Frees,
			"heapalloc": m.HeapAlloc,
		},
		// gc stats
		"gc": map[string]uint32{
			"numgc":       m.NumGC,
			"numForcedGC": m.NumForcedGC,
		},
		// cpus stats
		"cpus": map[string]int{
			"cores": numcpus,
		},
		// runtime stats
		"runtime": map[string]string{
			"version":  runtime.Version(),
			"compiler": runtime.Compiler,
		},
		// software version stats
		"version": map[string]string{
			"http":      echo.Version,
			"etherniti": release.Version,
		},
		// basic disk stats
		"disk": map[string]float64{
			"all":  float64(diskMonitor.All()) / gbUnits,
			"used": float64(diskMonitor.Used()) / gbUnits,
			"free": float64(diskMonitor.Free()) / gbUnits,
		},
	}

	return c.JSON(http.StatusOK, wrapper)
}

// return server side integrity message signed with private ecdsa key
func (ctl IndexController) integrity(c echo.Context) error {
	// get current date time
	currentTime := time.Now()
	timeStr := currentTime.String()
	millisStr := currentTime.Unix()

	//sign message
	hash, signature := integrity.SignMsgWithIntegrity(timeStr)
	wrapper := map[string]interface{}{
		"message":   timeStr,
		"millis":    millisStr,
		"hash":      hash,
		"signature": signature,
	}
	return c.JSON(http.StatusOK, wrapper)
}

// implemented method from interface RouterRegistrable
func (ctl IndexController) RegisterRouters(router *echo.Group) {
	log.Info("exposing index controller methods")
	router.GET("/", Index)
	router.GET("/status", ctl.status)
	router.GET("/integrity", ctl.integrity)
}
