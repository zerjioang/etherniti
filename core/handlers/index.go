// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"bytes"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/zerjioang/etherniti/core/api/protocol"
	"github.com/zerjioang/etherniti/core/eth/fastime"
	"github.com/zerjioang/etherniti/core/integrity"
	"github.com/zerjioang/etherniti/core/server/mods/mem"
	"github.com/zerjioang/etherniti/core/util"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/release"

	"github.com/zerjioang/etherniti/core/server/mods/disk"

	"github.com/labstack/echo"
)

type IndexController struct {
	lastIntegrityBytes []byte
	integrityLock      *sync.Mutex

	lastStatusBytes []byte
	statusLock      *sync.Mutex
}

const (
	gbUnits = float64(disk.GB)
)

var (
	//read only once, the number of server cpus
	numcpus = runtime.NumCPU()
	// monitor disk usage and get basic stats
	diskMonitor = disk.DiskUsagePtr()
	memMonitor  = mem.MemStatusMonitor()
	// integrity ticker (24h)
	integrityTicker = time.NewTicker(25 * time.Hour)
	statusTicker    = time.NewTicker(5 * time.Second)
)

// index data
const (
	/*IndexWelcomeJson = `{
	  "name": "eth-wbapi",
	  "description": "Etherniti: Ethereum REST API",
	  "cluster_name": "eth-wbapi",
	  "version": "` + release.Version + `",
	  "env": "` + config.EnvironmentName + `",
	  "tagline": "dapps everywhere"
	}`*/
	IndexWelcomeJson = `{"name":"eth-wbapi","description":"Etherniti:Ethereum REST API","cluster_name":"eth-wbapi","version":"` + release.Version + `","env":"` + config.EnvironmentName + `","tagline":"dapps everywhere"}`
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
	indexWelcomeBytes     = []byte(IndexWelcomeJson)
	indexWelcomeHtmlBytes = []byte(indexWelcomeHtml)
	// internally used struct pools to reduce GC
	statusPool = sync.Pool{
		// New creates an object when the pool has nothing available to return.
		// New must return an interface{} to make it flexible. You have to cast
		// your type after getting it.
		New: func() interface{} {
			// Pools often contain things like *bytes.Buffer, which are
			// temporary and re-usable.
			wrapper := new(protocol.ServerStatusResponse)
			wrapper.Runtime.Compiler = runtime.Compiler

			wrapper.Version.Etherniti = release.Version
			wrapper.Version.Go = runtime.Version()
			wrapper.Version.HTTP = echo.Version

			wrapper.Cpus.Cores = numcpus

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

// contructor like function
func init() {
	diskMonitor.Start("/")
	memMonitor.Start()
}

func NewIndexController() IndexController {
	dc := IndexController{}
	dc.integrityLock = new(sync.Mutex)
	dc.statusLock = new(sync.Mutex)
	// load initial value for integrity bytes
	dc.lastIntegrityBytes = util.GetJsonBytes(dc.integrityReload())
	dc.lastStatusBytes = dc.statusReload()

	go func(ctl *IndexController) {
		for range integrityTicker.C {
			// time to update integrity signature
			dc.integrityLock.Lock()
			ctl.lastIntegrityBytes = util.GetJsonBytes(ctl.integrityReload())
			dc.integrityLock.Unlock()
		}
	}(&dc)

	go func(ctl *IndexController) {
		for range statusTicker.C {
			// time to update integrity signature
			dc.statusLock.Lock()
			ctl.lastStatusBytes = ctl.statusReload()
			dc.statusLock.Unlock()
		}
	}(&dc)

	return dc
}

func Index(c echo.Context) error {
	if c.Request().Header.Get("Accept") == "application/json" {
		return CachedJsonBlob(c, true, CacheInfinite, indexWelcomeBytes)
	}
	return CachedHtml(c, true, CacheInfinite, indexWelcomeHtmlBytes)
}

func (ctl IndexController) Status(c echo.Context) error {
	data := ctl.status()
	var code int
	code, c = Cached(c, true, 5) // 5 seconds cache directive
	return c.JSONBlob(code, data)
}

func (ctl IndexController) status() []byte {
	ctl.statusLock.Lock()
	raw := ctl.lastStatusBytes
	ctl.statusLock.Unlock()
	return raw
}

func (ctl IndexController) statusReload() []byte {

	//get the wrapper from the pool, adn cast it
	wrapper := statusPool.Get().(*protocol.ServerStatusResponse)
	memMonitor.ReadPtr(wrapper)

	wrapper.Disk.All = diskMonitor.All()
	wrapper.Disk.Used = diskMonitor.Used()
	wrapper.Disk.Free = diskMonitor.Free()

	//get the buffer from the pool, adn cast it
	buffer := bufferBool.Get().(*bytes.Buffer)
	data := wrapper.Bytes(buffer)
	buffer.Reset()

	// Then put it back
	statusPool.Put(wrapper)
	bufferBool.Put(buffer)

	return data
}

// return server side integrity message signed with private ecdsa key
// concurrency safe
func (ctl IndexController) Integrity(c echo.Context) error {
	var code int
	code, c = Cached(c, true, 86400) // 24h cache directive
	return c.JSONBlob(code, ctl.integrity())
}

func (ctl IndexController) integrityReload() protocol.IntegrityResponse {
	// get current date time
	currentTime := fastime.Now()
	millis := currentTime.Unix()
	timeStr := time.Unix(millis, 0).Format(time.RFC3339)
	millisStr := strconv.FormatInt(millis, 10)

	//sign message
	hash, signature := integrity.SignMsgWithIntegrity(timeStr)

	var wrapper protocol.IntegrityResponse
	wrapper.Message = timeStr
	wrapper.Millis = millisStr
	wrapper.Hash = hash
	wrapper.Signature = signature
	return wrapper
}

func (ctl IndexController) integrity() []byte {
	ctl.integrityLock.Lock()
	raw := ctl.lastIntegrityBytes
	ctl.integrityLock.Unlock()
	return raw
}

// implemented method from interface RouterRegistrable
func (ctl IndexController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing index controller methods")
	router.GET("/", Index)
	router.GET("/status", ctl.Status)
	router.GET("/integrity", ctl.Integrity)
}
