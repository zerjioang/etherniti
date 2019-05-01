// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"bytes"
	"io/ioutil"
	"strings"

	"github.com/zerjioang/etherniti/shared/constants"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/util/banner"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/handlers/clientcache"

	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/zerjioang/etherniti/core/modules/concurrentbuffer"

	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/core/eth/fastime"
	"github.com/zerjioang/etherniti/core/modules/integrity"
	"github.com/zerjioang/etherniti/core/server/mods/mem"

	"github.com/zerjioang/etherniti/core/logger"

	"github.com/zerjioang/etherniti/core/server/mods/disk"

	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type IndexController struct {
	// use channels: https://talks.golang.org/2012/concurrency.slide#25
	statusData    concurrentbuffer.ConcurrentBuffer
	integrityData concurrentbuffer.ConcurrentBuffer
}

var (
	//read only once, the number of server cpus
	numcpus int
	// monitor disk usage and get basic stats
	diskMonitor *disk.DiskStatus
	memMonitor  mem.MemStatus
	// integrity ticker (24h)
	integrityTicker *time.Ticker
	// status ticker (5s)
	statusTicker *time.Ticker
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

func init() {
	//read only once, the number of server cpus
	numcpus = runtime.NumCPU()
	// monitor disk usage and get basic stats
	diskMonitor = disk.DiskUsagePtr()
	memMonitor = mem.MemStatusMonitor()
	// integrity ticker (24h)
	integrityTicker = time.NewTicker(25 * time.Hour)
	statusTicker = time.NewTicker(5 * time.Second)

	/*IndexWelcomeJson = `{
	  "name": "eth-wbapi",
	  "description": "Web3 REST Proxy",
	  "cluster_name": "eth-wbapi",
	  "version": "` + constants.Version + `",
	  "env": "` + config.EnvironmentName + `",
	  "tagline": "dapps everywhere"
	}`*/
	LoadIndexConstants()

	diskMonitor.Start("/")
}

func LoadIndexConstants() {
	// load constants
	IndexWelcomeJson = `{"name":"eth-wbapi","description":"High Performance Web3 REST Proxy","cluster_name":"eth-wbapi","version":"` + constants.Version + `","commit":"` + banner.Commit + `","env":"` + config.GetEnvironmentName() + `","tagline":"dapps everywhere"}`
	indexWelcomeBytes = []byte(IndexWelcomeJson)
	indexWelcomeHtmlBytes, _ = ioutil.ReadFile(config.ResourcesIndexHtml)
}

func NewIndexController() *IndexController {
	ctl := new(IndexController)
	ctl.statusData = concurrentbuffer.NewConcurrentBuffer()
	ctl.integrityData = concurrentbuffer.NewConcurrentBuffer()
	// load initial value for integrity bytes
	ctl.refreshIntegrityData()
	// load initial value for status bytes
	ctl.refreshStatusData()

	go func() {
		for range integrityTicker.C {
			// time to update integrity signature
			ctl.refreshIntegrityData()
		}
	}()

	go func() {
		for range statusTicker.C {
			// time to update status health data
			ctl.refreshStatusData()
		}
	}()

	return ctl
}

func Index(c *echo.Context) error {
	if strings.Contains(c.Request().Header.Get("Accept"), "application/json") {
		return clientcache.CachedJsonBlob(c, true, clientcache.CacheInfinite, indexWelcomeBytes)
	}
	return clientcache.CachedHtml(c, true, clientcache.CacheInfinite, indexWelcomeHtmlBytes)
}

func (ctl *IndexController) Status(c *echo.Context) error {
	data := ctl.status()
	var code int
	code, c = clientcache.Cached(c, true, 5) // 5 seconds cache directive
	return c.JSONBlob(code, data)
}

func (ctl *IndexController) status() []byte {
	return ctl.statusData.Bytes()
}

func (ctl *IndexController) refreshStatusData() {

	//get the wrapper from the pool, and cast it
	wrapper := statusPool.Get().(protocol.ServerStatusResponse)
	// force a new read memory
	memMonitor.ReadMemory()
	// read values
	memMonitor.ReadPtr(&wrapper)

	wrapper.Disk.All = diskMonitor.All()
	wrapper.Disk.Used = diskMonitor.Used()
	wrapper.Disk.Free = diskMonitor.Free()

	//get the buffer from the pool, and cast it
	buffer := bufferBool.Get().(*bytes.Buffer)
	data := wrapper.Bytes(buffer)
	buffer.Reset()

	// Then put it back
	statusPool.Put(wrapper)
	bufferBool.Put(buffer)

	ctl.statusData.Reset()
	_, _ = ctl.statusData.Write(data)
}

// return server side integrity message signed with private ecdsa key
// concurrency safe
func (ctl *IndexController) Integrity(c *echo.Context) error {
	data := ctl.integrity()
	var code int
	code, c = clientcache.Cached(c, true, 86400) // 24h cache directive
	return c.JSONBlob(code, data)
}

func (ctl *IndexController) refreshIntegrityData() {
	// get current date time
	millis := fastime.Now().Unix()
	timeStr := time.Unix(millis, 0).Format(time.RFC3339)
	millisStr := strconv.FormatInt(millis, 10)

	//sign message
	signMessage := "Hello from Etherniti Proxy. Today message generated at " + timeStr
	hash, signature := integrity.SignMsgWithIntegrity(signMessage)

	var wrapper protocol.IntegrityResponse
	wrapper.Message = signMessage
	wrapper.Millis = millisStr
	wrapper.Hash = hash
	wrapper.Signature = signature

	data := str.GetJsonBytes(wrapper)
	ctl.integrityData.Reset()
	_, _ = ctl.integrityData.Write(data)
}

func (ctl *IndexController) integrity() []byte {
	return ctl.integrityData.Bytes()
}

// implemented method from interface RouterRegistrable
func (ctl *IndexController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing index controller methods")
	router.GET("/", Index)
	router.GET("/status", ctl.Status)
	router.GET("/integrity", ctl.Integrity)
}
