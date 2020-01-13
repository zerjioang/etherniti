package shared

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/zerjioang/go-hpc/common"

	"github.com/zerjioang/go-hpc/thirdparty/echo/protocol/encoding"
	"github.com/zerjioang/go-hpc/util/ip"

	"github.com/zerjioang/go-hpc/lib/codes"
	"github.com/zerjioang/go-hpc/thirdparty/echo/protocol"

	"github.com/zerjioang/go-hpc/lib/eth/rpc/client"

	"github.com/zerjioang/etherniti/core/eth/profile"
	"github.com/zerjioang/etherniti/shared/constants"
	ethrpc "github.com/zerjioang/go-hpc/lib/eth/rpc"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

type EthernitiContext struct {
	echo.Context

	// connection profile data for interaction
	profileLock *sync.Mutex
	profileData *profile.ConnectionProfile

	//boost data cache value
	isJson     bool
	isXml      bool
	isTls      bool
	isWs       bool
	SchemeType protocol.RequestScheme
	SchemeName string
	intIp      uint32

	//x-profile http header token
	connectionToken string

	// http client cache policy
	OnSuccessCachePolicy int
	userId               string

	//response serialization protocol
	// json, fastjson, xml, msgpack, gogoproto, etc
	serializer common.Serializer
	//request content type identification string
	responseEncoding protocol.ContentTypeMode
}

var (
	noPeerAddressError = errors.New("no peer address to connect defined")
	isDebug            = true
)

//middleware function for the server
func EthernitiContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := AdquireContext(c)
		ret := next(cc)
		ReleaseContext(cc)
		return ret
	}
}

func AdquireContext(c echo.Context) *EthernitiContext {
	return newEthernitiContext(c)
}

func ReleaseContext(c *EthernitiContext) {
}

func newEthernitiContext(c echo.Context) *EthernitiContext {
	ec := new(EthernitiContext)
	if c != nil {
		ec.Context = c
		r, w := c.Request(), c.Response()
		c.Preload(r, w)
		ec.Preload(r, w)
	}
	return ec
}

func (c *EthernitiContext) Preload(r *http.Request, w http.ResponseWriter) {
	mode := c.Context.RequestContentType()
	c.isJson = mode == echo.MIMEApplicationJSON
	c.isXml = mode == echo.MIMEApplicationXML
	c.isTls = c.Request().TLS != nil
	c.isWs = strings.ToLower(c.Request().Header.Get(echo.HeaderUpgrade)) == "websocket"
	c.SchemeName = c.resolveScheme()
	c.intIp = ip.Ip2intLow(c.RealIP())
	//detect user requested data serialization method
	c.serializer, c.responseEncoding = encoding.EncodingSelector(mode)
}

func (c *EthernitiContext) IsTLS() bool {
	return c.isTls
}

func (c *EthernitiContext) IsWebSocket() bool {
	return c.isWs
}

func (c *EthernitiContext) IsJsonRequest() bool {
	return c.isJson
}
func (c *EthernitiContext) IsXmlRequest() bool {
	return c.isXml
}

func (c *EthernitiContext) IntIp() uint32 {
	return c.intIp
}

// check if request is http
func (c *EthernitiContext) IsHttp() bool {
	return c.SchemeType == protocol.Http
}

// check if request is https
func (c *EthernitiContext) IsHttps() bool {
	return c.SchemeType == protocol.Https
}

func (c *EthernitiContext) resolveScheme() string {
	// Can't use `r.Request.URL.Scheme`
	// See: https://groups.google.com/forum/#!topic/golang-nuts/pMUkBlQBDF0
	if c.IsTLS() {
		c.SchemeType = protocol.Https
		return "https"
	}
	req := c.Request()
	if scheme := req.Header.Get(echo.HeaderXForwardedProto); scheme != "" {
		c.SchemeType = protocol.Other
		return scheme
	}
	if scheme := req.Header.Get(echo.HeaderXForwardedProtocol); scheme != "" {
		c.SchemeType = protocol.Other
		return scheme
	}
	if ssl := req.Header.Get(echo.HeaderXForwardedSsl); ssl == "on" {
		c.SchemeType = protocol.Https
		return "https"
	}
	if scheme := req.Header.Get(echo.HeaderXUrlScheme); scheme != "" {
		c.SchemeType = protocol.Other
		return scheme
	}
	c.SchemeType = protocol.Http
	return "http"
}

func (c *EthernitiContext) RequestHost() string {
	return c.Request().Host
}

func (c *EthernitiContext) RequestUrl() string {
	return c.Request().RequestURI
}

func (c *EthernitiContext) Scheme() protocol.RequestScheme {
	return c.SchemeType
}

//added new functions
// returns connection profile from token information
func (c *EthernitiContext) ConnectionProfileSetup() (*profile.ConnectionProfile, error) {
	c.connectionToken = c.ReadConnectionProfileToken()
	// read connection profile token if exists
	readedProfile, err := profile.ParseConnectionProfileToken(c.connectionToken)
	if err == nil {
		//save profile data
		c.profileLock.Lock()
		c.profileData = readedProfile
		c.profileLock.Unlock()
	}
	return readedProfile, err
}

// get caller eth address
func (c *EthernitiContext) CallerEthAddress() string {
	var from string
	c.profileLock.Lock()
	if c.profileData != nil {
		from = c.profileData.Address
	}
	c.profileLock.Unlock()
	return from
}

// it recovers the eth client linked to it
// if peer url is provided, this peer address is used to dial
// otherwise, token information is readed in order to custom peer dial
func (c *EthernitiContext) RecoverEthClientFromTokenOrPeerUrl(peerUrl string, httpClient *client.EthClient) (*ethrpc.EthRPC, string, error) {
	client := new(ethrpc.EthRPC)
	var contextId string
	// by default, peer url is used to dial
	if peerUrl == "" {
		//no peer url found, try to read from user token
		if c.profileData == nil || c.profileData.RpcEndpoint == "" {
			return client, "", noPeerAddressError
		}
		contextId = c.profileData.RpcEndpoint
	} else {
		// use peer url
		contextId = peerUrl
	}
	client = ethrpc.NewDefaultRPCPtr(contextId, true, httpClient)
	return client, contextId, nil
}

// reads connection profile token from allowed sources
func (c *EthernitiContext) ReadConnectionProfileToken() string {
	req := c.Request()
	var tokenDataStr string
	// read if token provided via header key
	tokenDataStr = req.Header.Get(constants.HttpProfileHeaderkey)
	if tokenDataStr == "" {
		//read if token provided via query param
		tokenDataStr = c.QueryParam("profile")
		if tokenDataStr == "" {
			//read from request cookie
			cok, err := c.Cookie("profile")
			if err == nil && cok != nil {
				tokenDataStr = cok.Value
			}
		}
	}
	return tokenDataStr
}

// reads connection profile token from allowed sources
func (c *EthernitiContext) ReadToken(tokenName string) string {
	req := c.Request()
	var tokenDataStr string
	// read if token provided via header key
	tokenDataStr = req.Header.Get(tokenName)
	if tokenDataStr == "" {
		//read if token provided via query param
		tokenDataStr = c.QueryParam(tokenName)
		if tokenDataStr == "" {
			//read from request cookie
			cok, err := c.Cookie(tokenName)
			if err == nil && cok != nil {
				tokenDataStr = cok.Value
			}
		}
	}
	return tokenDataStr
}

func (c *EthernitiContext) User() *profile.ConnectionProfile {
	return c.profileData
}

// read user email
// this value is in jwt token setn with each request
func (c *EthernitiContext) AuthUserUuid() string {
	return c.profileData.AccountId
}

func (c *EthernitiContext) UserId() string {
	return c.userId
}

func (c *EthernitiContext) FastBlob(code codes.HttpStatusCode, contentType protocol.ContentTypeMode, b []byte) (err error) {
	r := c.Response()
	if code == codes.StatusOK {
		// add http client cache headers
		/*
			The Cache-Control header is the most important header
			to set as it effectively ‘switches on’ caching in
			the browser. With this header in place, and set with
			a value that enables caching, the browser will cache
			the file for as long as specified. Without this header
			the browser will re-request the file on each
			subsequent request.

			public resources can be cached not only by the
			end-user’s browser but also by any intermediate
			proxies that may be serving many other users as well.

			private resources are bypassed by intermediate
			proxies and can only be cached by the end-client.

			The max-age value sets a timespan for how
			long to cache the resource (in seconds).
		*/
		h := r.Header()
		if c.OnSuccessCachePolicy > 0 {
			timeStr := strconv.Itoa(c.OnSuccessCachePolicy)
			h.Set("Cache-Control", "public, max-age="+timeStr) // 24h cache = 86400
			h.Set("X-Cache", "HIT")
		}
	}
	c.WriteContentType(contentType.String())
	r.WriteHeaderCode(code)
	_, err = r.Write(b)
	return
}

// read rate limit identifier value
// this value can be:
// * the ip address
// * the token value
// * the ip address + token value
func (c *EthernitiContext) RateLimitIdentifier() string {
	clientIdentifier := c.RealIP()
	return clientIdentifier
}

func (c *EthernitiContext) ConnectionToken() string {
	return c.connectionToken
}

// check whether current context has a jWT or not
func (c *EthernitiContext) HasJWT() bool {
	return c.connectionToken != ""
}

func (c *EthernitiContext) SetTokenData(id string) {
	// todo implement this
	c.userId = id
}

func (c *EthernitiContext) Reset() {
	c.Context.Reset()
	//boots data cache value
	c.isJson = false
	c.isTls = false
	c.isWs = false
	c.SchemeType = protocol.Other
	c.SchemeName = ""
	c.intIp = 0
	c.connectionToken = ""
}

func (c *EthernitiContext) ResponseContentType() protocol.ContentTypeMode {
	return c.responseEncoding
}

func (c *EthernitiContext) ResponseSerializer() common.Serializer {
	return c.serializer
}
