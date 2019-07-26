// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package controllers

import (
	"fmt"
	"io/ioutil"
	"runtime"

	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/util/ip"
	"github.com/zerjioang/etherniti/core/util/net/ping"

	"github.com/zerjioang/etherniti/core/bench"

	"github.com/zerjioang/etherniti/core/api"

	"github.com/zerjioang/etherniti/core/modules/cpuid"

	"github.com/zerjioang/etherniti/core/modules/interval"

	"github.com/zerjioang/etherniti/shared/constants"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/util/banner"

	"sync"
	"time"

	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/core/server/mem"

	"github.com/zerjioang/etherniti/core/logger"

	"github.com/zerjioang/etherniti/core/server/disk"

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
	serverInfo            protocol.ServerInfoResponse
	// internally used struct pools to reduce GC
	statusPool = sync.Pool{
		// New creates an object when the pool has nothing available to return.
		// New must return an interface{} to make it flexible. You have to cast
		// your type after getting it.
		New: func() interface{} {
			// Pools often contain things like *bytes.Buffer, which are
			// temporary and re-usable.
			wrapper := protocol.ServerStatusResponse{}
			return wrapper
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
	// start monitoring root path
	diskMonitor.Start("/")
	// load info bytes
	loadInfoBytes()
}

func LoadIndexConstants() {
	// load constants
	IndexWelcomeJson = `{"name":"etherniti-public-api","description":"High Performance Web3 REST Proxy","cluster_name":"apollo-api","version":"` + banner.Version + `","commit":"` + banner.Commit + `","edition":"` + banner.Edition + `","env":"` + config.Env() + `","tagline":"dapps everywhere"}`
	indexWelcomeBytes = []byte(IndexWelcomeJson)
	indexWelcomeHtmlBytes, _ = ioutil.ReadFile(config.ResourcesIndexHtml)
	// reload info bytes to update banner.Version value
	loadInfoBytes()
}

func loadInfoBytes() {
	serverInfo.Architecture = runtime.GOARCH
	serverInfo.Os = runtime.GOOS
	serverInfo.Runtime.Compiler = runtime.Compiler

	serverInfo.Version.Etherniti = banner.Version
	serverInfo.Version.Go = runtime.Version()
	serverInfo.Version.HTTP = echo.Version

	serverInfo.Cpus.Cores = runtime.NumCPU()
	// load cpu features
	serverInfo.Cpus.Features = cpuid.GetCpuFeatures()
}

func NewIndexController() *IndexController {
	return new(IndexController)
}

func (ctl *IndexController) Index(c *echo.Context) error {
	c.OnSuccessCachePolicy = constants.CacheInfinite
	if c.IsJsonRequest() {
		return c.JSONBlob(protocol.StatusOK, indexWelcomeBytes)
	}
	return c.HTMLBlob(protocol.StatusOK, indexWelcomeHtmlBytes)
}

func (ctl *IndexController) Info(c *echo.Context) error {
	c.OnSuccessCachePolicy = constants.CacheInfinite
	return api.SendSuccess(c, []byte("info"), serverInfo)
}

func (ctl *IndexController) Status(c *echo.Context) error {
	data := ctl.status()
	c.OnSuccessCachePolicy = 5
	return c.JSONBlob(protocol.StatusOK, data)
}

func (ctl *IndexController) status() []byte {
	return statusTicker.Bytes()
}

// return server side integrity message signed with private ecdsa key
// concurrency safe
func (ctl *IndexController) Integrity(c *echo.Context) error {
	data := ctl.integrity()
	c.OnSuccessCachePolicy = constants.CacheOneDay
	return c.JSONBlob(protocol.StatusOK, data)
}

func (ctl *IndexController) integrity() []byte {
	return integrityTicker.Bytes()
}

func (ctl *IndexController) Ping(c *echo.Context) error {
	targetIp := c.QueryParam("ip")
	if !ip.IsIpv4(targetIp) {
		logger.Error("failed to read request target ip: ", data.ErrInvalidIpv4)
		return api.Error(c, data.ErrInvalidIpv4)
	}
	response, err := ctl.ping(targetIp)
	if err != nil {
		return api.Error(c, err)
	} else {
		return api.SendSuccess(c, []byte("icmp_ping"), response)
	}
}

func (ctl *IndexController) ping(addr string) (*ping.Statistics, error) {

	pinger, err := ping.NewPinger(addr)
	pinger.Count = 5
	pinger.Interval = time.Second * 1
	pinger.Timeout = time.Second * 2 // max: count * interval
	pinger.SetPrivileged(false)

	if err != nil {
		logger.Error("failed to create new ping tester: ", err.Error())
		return nil, err
	}

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v\n",
			pkt.Nbytes,
			pkt.IPAddr,
			pkt.Seq,
			pkt.Rtt,
			pkt.Ttl,
		)
	}

	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n",
			stats.Addr,
		)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent,
			stats.PacketsRecv,
			stats.PacketLoss,
		)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt,
			stats.AvgRtt,
			stats.MaxRtt,
			stats.StdDevRtt,
		)
	}

	pinger.Run()
	s := pinger.Statistics()
	return s, nil
}

// todo optimize struct creation. it should be created once, not every time is called by http clients. smae goes for byte array
func (ctl *IndexController) score(c *echo.Context) error {
	scoreWrapper := struct {
		Time  time.Duration `json:"time"`
		Score int64         `json:"score"`
	}{
		Time:  bench.GetBenchTime(),
		Score: bench.GetScore(),
	}
	return api.SendSuccess(c, []byte("bench_score"), scoreWrapper)
}

// implemented method from interface RouterRegistrable
func (ctl *IndexController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing index controller methods")
	router.GET("/", ctl.Index)
	router.GET("/info", ctl.Info)
	router.GET("/score", ctl.score)
	router.GET("/metrics", ctl.Status)
	router.GET("/integrity", ctl.Integrity)
	router.GET("/ping", ctl.Ping)
}
