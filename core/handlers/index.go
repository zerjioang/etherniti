// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"bytes"
	"io/ioutil"

	"github.com/zerjioang/etherniti/core/modules/interval"

	"github.com/zerjioang/etherniti/shared/constants"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/util/banner"

	"github.com/zerjioang/etherniti/core/handlers/clientcache"

	"runtime"
	"sync"
	"time"

	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/core/server/mods/mem"

	"github.com/zerjioang/etherniti/core/logger"

	"github.com/zerjioang/etherniti/core/server/mods/disk"

	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type IndexController struct {
	// use channels: https://talks.golang.org/2012/concurrency.slide#25
}

var (
	// monitor disk usage and get basic stats
	diskMonitor *disk.DiskStatus
	memMonitor  mem.MemStatus
	// integrity ticker (24h)
	integrityTicker *interval.IntervalTask
	// status ticker (5s)
	statusTicker *interval.IntervalTask
	//bytes of welcome message
	IndexWelcomeJson      string
	indexWelcomeBytes     []byte
	indexWelcomeHtmlBytes []byte
	// internally used struct pools to reduce GC
	statusPool = sync.Pool{
		// New creates an object when the pool has nothing available to return.
		// New must return an interface{} to make it flexible. You have to cast
		// your type after getting it.
		New: func() interface{} {
			// Pools often contain things like *bytes.Buffer, which are
			// temporary and re-usable.
			wrapper := protocol.ServerStatusResponse{}
			wrapper.Architecture = runtime.GOARCH
			wrapper.Os = runtime.GOOS
			wrapper.Runtime.Compiler = runtime.Compiler

			wrapper.Version.Etherniti = constants.Version
			wrapper.Version.Go = runtime.Version()
			wrapper.Version.HTTP = echo.Version

			wrapper.Cpus.Cores = runtime.NumCPU()

			return wrapper
		},
	}
	bufferBool = sync.Pool{
		// New creates an object when the pool has nothing available to return.
		// New must return an interface{} to make it flexible. You have to cast
		// your type after getting it.
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
)

func init() {
	// monitor disk usage and get basic stats
	diskMonitor = disk.DiskUsagePtr()
	// monitor memory usage and get basic stats
	memMonitor = mem.MemStatusMonitor()
	// integrity ticker (24h)
	//integrityTicker = interval.NewTask(24 * time.Hour, interval.Loop).Do()
	integrityTicker = interval.NewTask("integrity", 5*time.Second, interval.Loop, true, onNewIntegrityData).Do()
	// status ticker update each 5s
	statusTicker = interval.NewTask("status", 5*time.Second, interval.Loop, true, onNewStatusData).Do()

	/*IndexWelcomeJson = `{
	  "name": "eth-wbapi",
	  "description": "Web3 REST Proxy",
	  "cluster_name": "eth-wbapi",
	  "version": "` + constants.Version + `",
	  "env": "` + config.EnvironmentName + `",
	  "tagline": "dapps everywhere"
	}`*/
	LoadIndexConstants()
	// start monitoring root path
	diskMonitor.Start("/")
}

func LoadIndexConstants() {
	// load constants
	IndexWelcomeJson = `{"name":"eth-wbapi","description":"High Performance Web3 REST Proxy","cluster_name":"eth-wbapi","version":"` + constants.Version + `","commit":"` + banner.Commit + `","env":"` + config.GetEnvironmentName() + `","tagline":"dapps everywhere"}`
	indexWelcomeBytes = []byte(IndexWelcomeJson)
	indexWelcomeHtmlBytes, _ = ioutil.ReadFile(config.ResourcesIndexHtml)
}

func NewIndexController() *IndexController {
	return new(IndexController)
}

func Index(c *echo.Context) error {
	c.OnSuccessCachePolicy = clientcache.CacheInfinite
	if c.IsJsonRequest() {
		return c.JSONBlob(protocol.StatusOK, indexWelcomeBytes)
	}
	return c.HTMLBlob(protocol.StatusOK, indexWelcomeHtmlBytes)
}

func (ctl *IndexController) Status(c *echo.Context) error {
	data := ctl.status()
	var code int
	c.OnSuccessCachePolicy = 5
	return c.JSONBlob(code, data)
}

func (ctl *IndexController) status() []byte {
	return statusTicker.Bytes()
}

// return server side integrity message signed with private ecdsa key
// concurrency safe
func (ctl *IndexController) Integrity(c *echo.Context) error {
	data := ctl.integrity()
	var code int
	c.OnSuccessCachePolicy = clientcache.CacheOneDay
	return c.JSONBlob(code, data)
}

func (ctl *IndexController) integrity() []byte {
	return integrityTicker.Bytes()
}

// implemented method from interface RouterRegistrable
func (ctl *IndexController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing index controller methods")
	router.GET("/", Index)
	router.GET("/status", ctl.Status)
	router.GET("/integrity", ctl.Integrity)
}
