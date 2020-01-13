// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package index

import (
	"fmt"
	"io/ioutil"
	"runtime"

	"github.com/zerjioang/go-hpc/lib/metrics/model"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/controllers/wrap"
	"github.com/zerjioang/go-hpc/thirdparty/echo/protocol"

	"github.com/zerjioang/go-hpc/lib/codes"
	"github.com/zerjioang/go-hpc/lib/metrics/disk"
	"github.com/zerjioang/go-hpc/lib/metrics/mem"

	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/shared"
	"github.com/zerjioang/go-hpc/util/ip"
	"github.com/zerjioang/go-hpc/util/net/ping"

	"github.com/zerjioang/go-hpc/lib/cpuid"

	"github.com/zerjioang/go-hpc/lib/interval"

	"github.com/zerjioang/etherniti/shared/constants"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/util/banner"

	"sync"
	"time"

	"github.com/zerjioang/etherniti/shared/dto"

	"github.com/zerjioang/etherniti/core/logger"

	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

type IndexController struct {
	// use channels: https://talks.golang.org/2012/concurrency.slide#25
}

const (
	//bytes of welcome message in xml
	IndeWelcomeXml = `<?xml version="1.0" encoding="UTF-8"?>
<root><cluster_name>apollo-api</cluster_name><commit>latest</commit><description>High Performance Web3 REST Proxy</description><edition>oss</edition><env>development</env><name>etherniti-public-api</name><tagline>dapps everywhere</tagline><version>latest</version></root>`
)

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
	serverInfo            dto.ServerInfoResponse
	// internally used struct pools to reduce GC
	statusPool = sync.Pool{
		// New creates an object when the pool has nothing available to return.
		// New must return an interface{} to make it flexible. You have to cast
		// your type after getting it.
		New: func() interface{} {
			// Pools often contain things like *bytes.Buffer, which are
			// temporary and re-usable.
			wrapper := model.ServerStatusResponse{}
			return wrapper
		},
	}
	infoBytes = []byte("info")
)

func init() {
	// monitor disk usage and get basic stats
	diskMonitor = disk.DiskUsagePtr()
	// monitor memory usage and get basic stats
	memMonitor = mem.MemStatusMonitor()
	integrityTicker = interval.NewTask("integrity", 24*time.Hour, interval.Loop, true, onNewIntegrityData).Do()
	// status ticker update each 10s
	statusTicker = interval.NewTask("status", 10*time.Second, interval.Loop, true, onNewStatusData).Do()
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

func (ctl *IndexController) Index(c *shared.EthernitiContext) error {
	c.OnSuccessCachePolicy = constants.CacheInfinite
	if c.IsJsonRequest() {
		return c.JSONBlob(codes.StatusOK, indexWelcomeBytes)
	} else if c.IsXmlRequest() {
		return c.FastBlob(codes.StatusOK, protocol.ModeXML, []byte(IndeWelcomeXml))
	}
	return c.HTMLBlob(codes.StatusOK, indexWelcomeHtmlBytes)
}

func (ctl *IndexController) Info(c *shared.EthernitiContext) error {
	c.OnSuccessCachePolicy = constants.CacheInfinite
	return api.SendSuccess(c, infoBytes, serverInfo)
}

func (ctl *IndexController) Status(c *shared.EthernitiContext) error {
	d := ctl.status()
	c.OnSuccessCachePolicy = 5
	return c.JSONBlob(codes.StatusOK, d)
}

func (ctl *IndexController) status() []byte {
	return statusTicker.Bytes()
}

// return server side integrity message signed with private ecdsa key
// concurrency safe
func (ctl *IndexController) Integrity(c *shared.EthernitiContext) error {
	d := ctl.integrity()
	c.OnSuccessCachePolicy = constants.CacheOneDay
	return c.JSONBlob(codes.StatusOK, d)
}

func (ctl *IndexController) integrity() []byte {
	return integrityTicker.Bytes()
}

func (ctl *IndexController) Ping(c *shared.EthernitiContext) error {
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

// This library attempts to send an "unprivileged" ping via UDP. On linux, this must be enabled by setting
//
// sudo sysctl -w net.ipv4.ping_group_range="0   2147483647"
func (ctl *IndexController) ping(addr string) (*ping.Statistics, error) {

	pinger, err := ping.NewPinger(addr)
	if err != nil {
		logger.Error("failed to create new ping tester: ", err.Error())
		return nil, err
	}
	pinger.Count = 5
	pinger.Interval = time.Second * 1
	pinger.Timeout = time.Second * 2 // max: count * interval
	pinger.SetPrivileged(false)
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

// implemented method from interface RouterRegistrable
func (ctl *IndexController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing index controller methods")
	router.GET("/", wrap.Call(ctl.Index))
	router.GET("/info", wrap.Call(ctl.Info))
	router.GET("/metrics", wrap.Call(ctl.Status))
	router.GET("/integrity", wrap.Call(ctl.Integrity))
	router.GET("/ping", wrap.Call(ctl.Ping))
}
