// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"bytes"

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
	"github.com/zerjioang/etherniti/core/integrity"
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
	indexWelcomeHtml      string
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
	  "description": "Ethereum REST API Proxy",
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
	IndexWelcomeJson = `{"name":"eth-wbapi","description":"High Performance Ethereum REST API Proxy","cluster_name":"eth-wbapi","version":"` + constants.Version + `","commit":"` + banner.Commit + `","env":"` + config.EnvironmentName + `","tagline":"dapps everywhere"}`
	indexWelcomeHtml = `<!DOCTYPE html><html> <head> <title>Etherniti - High Performance Ethereum REST API Proxy</title> <meta name="robots" content="noindex, nofollow"/> <meta name="viewport" content="width=device-width, initial-scale=1.0"> <style type="text/css"> @import url('https://fonts.googleapis.com/css?family=Lato:300,400,700'); html, body, div, span, applet, object, iframe, h1, h2, h3, h4, h5, h6, p, blockquote, pre, a, abbr, acronym, address, big, cite, code, del, dfn, em, img, ins, kbd, q, s, samp, small, strike, strong, sub, sup, tt, var, b, u, i, center, dl, dt, dd, ol, ul, li, fieldset, form, label, legend, table, caption, tbody, tfoot, thead, tr, th, td, article, aside, canvas, details, embed, figure, figcaption, footer, header, hgroup, menu, nav, output, ruby, section, summary, time, mark, audio, video{margin: 0; padding: 0; border: 0; font-size: 100%; font: inherit; vertical-align: baseline;}/* HTML5 display-role reset for older browsers */ article, aside, details, figcaption, figure, footer, header, hgroup, menu, nav, section{display: block;}a{text-decoration: none; text-transform: none; color: #4A90E2;}body{line-height: 1; font-family: lato, ubuntu,-apple-system,BlinkMacSystemFont,Segoe UI,Roboto,Oxygen,Ubuntu,Cantarell,Open Sans,Helvetica Neue,sans-serif; text-rendering : optimizeLegibility; -webkit-font-smoothing : antialiased; font-size: 19px; background-color: #FEFEFE; color: #04143A;}ol, ul{list-style: none;}blockquote, q{quotes: none;}blockquote:before, blockquote:after, q:before, q:after{content: ''; content: none;}table{border-collapse: collapse; border-spacing: 0;}p{color: #15171a; font-size: 17; line-height: 31px;}strong{font-weight: 600;}div , footer{box-sizing: border-box;}/* Reset ends */ /*Hero section*/ .container{max-width: 1100px; height: auto; margin: 60px auto;}.hero{margin: 50px auto; position: relative;}h1.name{font-size: 70px; font-weight: 300; display: inline-block;}.job-title{vertical-align: top; background-color: #D9E7F8; color: #4A90E2; font-weight: 600; margin-top: 5px; margin-left: 20px; border-radius: 5px; display: inline-block; padding: 15px 25px;}.email{display: block; font-size: 24px; font-weight: 300; color: #81899C; margin-top: 10px;}.lead{font-size: 44px; font-weight: 300; margin-top: 60px; line-height: 55px;}/*hero ends*/ /*skills & intrests*/ .sections{vertical-align: top; display: inline-block; width: 49.7%; height: 50px;}.section-title{font-size: 20px; font-weight: 600; margin-bottom: 15px;}.list-card{margin: 30px 0;}.list-card .exp , .list-card div{display: inline-block; vertical-align: top;}.list-card .exp{margin-right: 15px; color: #4A90E2; font-weight: 600; width: 100px;}.list-card div{width: 70%;}.list-card h3{font-size: 20px; font-weight: 600; color: #5B6A9A; line-height: 26px; margin-bottom: 8px;}.list-card div span{font-size: 16px; color: #81899C; line-height: 22px;}/*skill and intrests ends*/ /* Achievements */ .cards{max-width: 1120px; display: block; margin-top: 280px;}.card{width: 47.9%; height: 200px; background-color: #EEF0F7; display: inline-block; margin: 7px 5px; vertical-align: top; border-radius: 10px; text-align: center; padding-top: 50px}.card-active , .card:hover{transform: scale(1.02); transition: 0.5s; background-color: #fff; box-shadow: 0px 5px 50px -8px #ddd; cursor: pointer;}.skill-level{display: inline-block; max-width: 160px;}.skill-level span{font-size: 35px; font-weight: 300; color: #5B6A9A; vertical-align: top;}.skill-level h2{font-size: 95px; font-weight: 300; display: inline-block; vertical-align: top; color: #5B6A9A; letter-spacing: -5px;}.skill-meta{vertical-align: top; display: inline-block; max-width: 300px; text-align: left; margin-top: 15px; margin-left: 15px;}.skill-meta h3{font-size: 20px; font-weight: 800; color: #5B6A9A; margin-bottom: 5px;}.skill-meta span{color: #81899C; line-height: 20px; font-size: 16px;}/* Achievements ends */ /* Timeline styles*/ ol{position: relative; display: block; margin: 100px 0; height: 2px; background: #EEF0F7;}ol::before, ol::after{content: ""; position: absolute; top: -10px; display: block; width: 0; height: 0; border-radius: 10px; border: 0px solid #31708F;}ol::before{left: -5px;}ol::after{right: -10px; border: 0px solid transparent; border-right: 0; border-left: 20px solid #31708F; border-radius: 3px;}/* ---- Timeline elements ---- */ li{position: relative; display: inline-block; float: left; width: 25%; height: 50px;}li .line{position: absolute; top: -47px; left: 1%; font-size: 20px; font-weight: 600; color: #04143A;}li .point{content: ""; top: -7px; left: 0%; display: block; width: 8px; height: 8px; border: 4px solid #fff; border-radius: 10px; background: #4A90E2; position: absolute;}li .description{display: none; padding: 10px 0; margin-top: 20px; position: relative; font-weight: normal; z-index: 1; max-width: 95%; font-size: 18px; font-weight: 600; line-height: 25px; color: #5B6A9A;}.description::before{content: ''; width: 0; height: 0; border-left: 5px solid transparent; border-right: 5px solid transparent; border-bottom: 5px solid #f4f4f4; position: absolute; top: -5px; left: 43%;}.timeline .date{font-size: 14px; color: #81899C; font-weight: 300;}/* ---- Hover effects ---- */ li:hover{color: #48A4D2;}li .description{display: block;}/*timeline ends*/ /* Media queries*/ @media(max-width: 1024px){.container{padding: 15px; margin: 0px auto;}.cards{margin-top: 250px;}}@media(max-width: 768px){.container{padding: 15px; margin: 0px auto;}.cards{margin-top: 320px;}.card{padding: 15px; text-align: left;}.card h2{font-size: 70px;}.card , .sections{width: 100%; height: auto; margin: 10px 0; float: left;}.timeline{border: none; background-color: rgba(0,0,0,0);}.timeline li{margin-top: 70px; height: 150px;}}@media(max-width: 425px){h1.name{font-size: 40px;}.card , .sections{width: 100%; height: auto; margin: 10px 0; float: left;}.timeline{display: none;}.job-title{position: absolute; font-size: 15px; top: -40px; right: 20px; padding: 10px}.lead{margin-top: 15px; font-size: 20px; line-height: 28px;}.container{margin: 0px; padding: 0 15px;}footer{margin-top: 2050px;}}/* configure selection color */ /* css selector style */ ::-moz-selection{/* Code for Firefox */ color: #fff9f9; background: #2253bc;}::selection{color: #fff9f9; background: #2253bc;}</style> </head> <body> <div class="container" style="text-align: center;"> <div class="hero"> <h1 class="name"><strong>Etherniti</strong> REST API</h1> <span class="job-title">v` + constants.Version + ` (commit ` + banner.Commit + `) </span> <span class="email">` + config.EnvironmentName + ` build</span> <h2 class="lead">A High Performance, Multitenant RESTful proxy service for Ethereum</h2> </div></div><div class="container"> <div class="sections"> <h2 class="section-title">Features</h2> <div class="list-card"> <span class="exp">&#10003;</span> <div> <h3>Full Ethereum RPC support</h3> <span>100% drop-in replacement</span> </div></div><div class="list-card"> <span class="exp">&#10003;</span> <div> <h3>Web3 compatible</h3> <span>Interact even easier than using web3j libs</span> </div></div><div class="list-card"> <span class="exp">&#10003;</span> <div> <h3>ECDSA signature</h3> <span>offline + online signing supported</span> </div></div></div><div class="sections"> <h2 class="section-title">Features</h2> <div class="list-card"> <span class="exp">&#10003;</span> <div> <h3>High performance design</h3> <span>Deliver high performance results</span> </div></div><div class="list-card"> <span class="exp">&#10003;</span> <div> <h3>Designed for IIOT</h3> <span>connect your IIOT devices effortless</span> </div></div></div></div><br></br> <br></br> <br></br> <br></br> <br></br> <footer class="container" style="text-align: center;"> <span style="font-size: 16px; margin-top: ">Please refer to official API documentation for further details or visit <a href="http://dev-proxy.etherniti.org/swagger">http://dev-proxy.etherniti.org/swagger </a> </span> </footer> </body></html>`
	indexWelcomeBytes = []byte(IndexWelcomeJson)
	indexWelcomeHtmlBytes = []byte(indexWelcomeHtml)
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

func Index(c echo.Context) error {
	if c.Request().Header.Get("Accept") == "application/json" {
		return clientcache.CachedJsonBlob(c, true, clientcache.CacheInfinite, indexWelcomeBytes)
	}
	return clientcache.CachedHtml(c, true, clientcache.CacheInfinite, indexWelcomeHtmlBytes)
}

func (ctl *IndexController) Status(c echo.Context) error {
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
func (ctl *IndexController) Integrity(c echo.Context) error {
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
