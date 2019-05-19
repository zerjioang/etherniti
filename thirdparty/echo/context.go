// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package echo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/json-iterator/go"
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/eth/profile"
	"github.com/zerjioang/etherniti/core/eth/rpc"
	"github.com/zerjioang/etherniti/shared/constants"

	"github.com/zerjioang/etherniti/core/modules/concurrentmap"
)

var (
	errInvalidateCache = errors.New("failed to get item from internal cache, cache invalidation issues around")
)

// Context represents the Context of the current HTTP request. It holds request and
// response objects, path, path parameters, data and registered handler.
type Context struct {
	request  *http.Request
	response *Response
	path     string
	pnames   []string
	pvalues  []string
	query    url.Values
	handler  HandlerFunc
	store    Map
	echo     *Echo
	lock     sync.RWMutex

	// connection profile data for interaction
	profileLock *sync.Mutex
	profileData profile.ConnectionProfile

	//boots data cache value
	isJson     bool
	isTls      bool
	isWs       bool
	SchemeType constants.RequestScheme
	SchemeName string
	ip         string
	// http client cache policy
	OnSuccessCachePolicy int
	UserId               string
}

var (
	jsonfast           = jsoniter.ConfigFastest
	noPeerAddressError = errors.New("no peer address to connect defined")
	isDebug            = config.IsDevelopment()
	fileCache          concurrentmap.ConcurrentMap
)

func init() {
	fileCache = concurrentmap.New()
}

const (
	defaultMemory = 32 << 20 // 32 MB
	indexPage     = "index.html"
)

func (c *Context) Preload() {
	c.isJson = strings.Contains(c.request.Header.Get("Accept"), "application/json")
	c.isTls = c.request.TLS != nil
	c.isWs = strings.ToLower(c.request.Header.Get(HeaderUpgrade)) == "websocket"
	c.SchemeName = c.resolveScheme()
	c.ip = c.resolveRealIP()
}

func (c *Context) writeContentType(value string) {
	header := c.response.Header()
	if header.Get(HeaderContentType) == "" {
		header.Set(HeaderContentType, value)
	}
}

func (c *Context) Request() *http.Request {
	return c.request
}

func (c *Context) SetRequest(r *http.Request) {
	c.request = r
}

func (c *Context) Response() *Response {
	return c.response
}

func (c *Context) IsTLS() bool {
	return c.isTls
}

func (c *Context) IsWebSocket() bool {
	return c.isWs
}

func (c *Context) resolveScheme() string {
	// Can't use `r.Request.URL.Scheme`
	// See: https://groups.google.com/forum/#!topic/golang-nuts/pMUkBlQBDF0
	if c.IsTLS() {
		c.SchemeType = constants.Https
		return "https"
	}
	if scheme := c.request.Header.Get(HeaderXForwardedProto); scheme != "" {
		c.SchemeType = constants.Other
		return scheme
	}
	if scheme := c.request.Header.Get(HeaderXForwardedProtocol); scheme != "" {
		c.SchemeType = constants.Other
		return scheme
	}
	if ssl := c.request.Header.Get(HeaderXForwardedSsl); ssl == "on" {
		c.SchemeType = constants.Https
		return "https"
	}
	if scheme := c.request.Header.Get(HeaderXUrlScheme); scheme != "" {
		c.SchemeType = constants.Other
		return scheme
	}
	c.SchemeType = constants.Http
	return "http"
}

func (c *Context) Scheme() constants.RequestScheme {
	return c.SchemeType
}

func (c *Context) resolveRealIP() string {
	if ip := c.request.Header.Get(HeaderXForwardedFor); ip != "" {
		return strings.Split(ip, ", ")[0]
	}
	if ip := c.request.Header.Get(HeaderXRealIP); ip != "" {
		return ip
	}
	ra, _, _ := net.SplitHostPort(c.request.RemoteAddr)
	return ra
}

func (c *Context) RealIP() string {
	return c.ip
}

func (c *Context) Path() string {
	return c.path
}

func (c *Context) SetPath(p string) {
	c.path = p
}

func (c *Context) Param(name string) string {
	for i, n := range c.pnames {
		if i < len(c.pvalues) {
			if n == name {
				return c.pvalues[i]
			}
		}
	}
	return ""
}

func (c *Context) ParamNames() []string {
	return c.pnames
}

func (c *Context) SetParamNames(names ...string) {
	c.pnames = names
}

func (c *Context) ParamValues() []string {
	return c.pvalues[:len(c.pnames)]
}

func (c *Context) SetParamValues(values ...string) {
	c.pvalues = values
}

func (c *Context) QueryParam(name string) string {
	if c.query == nil {
		c.query = c.request.URL.Query()
	}
	return c.query.Get(name)
}

func (c *Context) QueryParams() url.Values {
	if c.query == nil {
		c.query = c.request.URL.Query()
	}
	return c.query
}

func (c *Context) QueryString() string {
	return c.request.URL.RawQuery
}

func (c *Context) FormValue(name string) string {
	return c.request.FormValue(name)
}

func (c *Context) FormParams() (url.Values, error) {
	if strings.HasPrefix(c.request.Header.Get(HeaderContentType), MIMEMultipartForm) {
		if err := c.request.ParseMultipartForm(defaultMemory); err != nil {
			return nil, err
		}
	} else {
		if err := c.request.ParseForm(); err != nil {
			return nil, err
		}
	}
	return c.request.Form, nil
}

func (c *Context) FormFile(name string) (*multipart.FileHeader, error) {
	_, fh, err := c.request.FormFile(name)
	return fh, err
}

func (c *Context) MultipartForm() (*multipart.Form, error) {
	err := c.request.ParseMultipartForm(defaultMemory)
	return c.request.MultipartForm, err
}

func (c *Context) Cookie(name string) (*http.Cookie, error) {
	return c.request.Cookie(name)
}

func (c *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(c.response, cookie)
}

func (c *Context) Cookies() []*http.Cookie {
	return c.request.Cookies()
}

func (c *Context) Get(key string) interface{} {
	c.lock.RLock()
	v := c.store[key]
	c.lock.RUnlock()
	return v
}

func (c *Context) Set(key string, val interface{}) {
	c.lock.Lock()

	if c.store == nil {
		c.store = make(Map)
	}
	c.store[key] = val
	c.lock.Unlock()
}

func (c *Context) Bind(i interface{}) error {
	err := json.NewDecoder(c.request.Body).Decode(i)
	return err
}

func (c *Context) Render(code int, name string, data interface{}) (err error) {
	if c.echo.Renderer == nil {
		return ErrRendererNotRegistered
	}
	buf := new(bytes.Buffer)
	if err = c.echo.Renderer.Render(buf, name, data, c); err != nil {
		return
	}
	return c.HTMLBlob(code, buf.Bytes())
}

func (c *Context) HTML(code int, html string) (err error) {
	return c.HTMLBlob(code, []byte(html))
}

func (c *Context) HTMLBlob(code int, b []byte) (err error) {
	return c.FastBlob(code, MIMETextHTMLCharsetUTF8, b)
}

func (c *Context) String(code int, s string) (err error) {
	return c.FastBlob(code, MIMETextPlainCharsetUTF8, []byte(s))
}

func (c *Context) jsonPBlob(code int, callback string, i interface{}) (err error) {
	enc := json.NewEncoder(c.response)
	_, pretty := c.QueryParams()["pretty"]
	if c.echo.Debug || pretty {
		enc.SetIndent("", "  ")
	}
	c.writeContentType(MIMEApplicationJavaScriptCharsetUTF8)
	c.response.WriteHeader(code)
	if _, err = c.response.Write([]byte(callback + "(")); err != nil {
		return
	}
	if err = enc.Encode(i); err != nil {
		return
	}
	if _, err = c.response.Write([]byte(");")); err != nil {
		return
	}
	return
}

func (c *Context) json(code int, i interface{}, indent string) error {
	enc := json.NewEncoder(c.response)
	if indent != "" {
		enc.SetIndent("", indent)
	}
	c.writeContentType(MIMEApplicationJSONCharsetUTF8)
	c.response.WriteHeader(code)
	return enc.Encode(i)
}

//custom json encoder
func (c *Context) JSON(code int, i interface{}) (err error) {
	raw, encErr := jsonfast.Marshal(&i)
	if encErr != nil {
		return encErr
	}
	return c.Blob(code, MIMEApplicationJSONCharsetUTF8, raw)
}

func (c *Context) JSONBlob(code int, b []byte) (err error) {
	return c.FastBlob(code, MIMEApplicationJSONCharsetUTF8, b)
}

func (c *Context) JSONP(code int, callback string, i interface{}) (err error) {
	return c.jsonPBlob(code, callback, i)
}

func (c *Context) JSONPBlob(code int, callback string, b []byte) (err error) {
	c.writeContentType(MIMEApplicationJavaScriptCharsetUTF8)
	c.response.WriteHeader(code)
	if _, err = c.response.Write([]byte(callback + "(")); err != nil {
		return
	}
	if _, err = c.response.Write(b); err != nil {
		return
	}
	_, err = c.response.Write([]byte(");"))
	return
}

func (c *Context) Blob(code int, contentType string, b []byte) (err error) {
	c.writeContentType(contentType)
	c.response.WriteHeader(code)
	_, err = c.response.Write(b)
	return
}

func (c *Context) FastBlob(code int, contentType string, b []byte) (err error) {
	if code == protocol.StatusOK {
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
		r := c.Response()
		h := r.Header()
		if c.OnSuccessCachePolicy > 0 {
			timeStr := strconv.Itoa(c.OnSuccessCachePolicy)
			h.Set("Cache-Control", "public, max-age="+timeStr) // 24h cache = 86400
		}
	}
	c.writeContentType(contentType)
	c.response.WriteHeader(code)
	_, err = c.response.Write(b)
	return
}

func (c *Context) Stream(code int, contentType string, r io.Reader) (err error) {
	c.writeContentType(contentType)
	c.response.WriteHeader(code)
	_, err = io.Copy(c.response, r)
	return
}

func (c *Context) File(file string) (err error) {
	initialFilePath := file
	// check if file is cached
	// check if file was already readed before and saved in our cache
	// this avoid overhead on disk readings
	object, found := fileCache.Get(initialFilePath)
	if found && object != nil {
		// cast
		buffer, ok := object.(*FileBuffer)
		if ok {
			//casting was ok
			// add a http cache directive too
			c.response.Header().Set("Cache-Control", "public, max-age=86400") // 24h cache = 86400
			http.ServeContent(c.response, c.request, buffer.name, buffer.time, buffer)
			return nil
		} else {
			// some cache and data error occured.
			return errInvalidateCache
		}
	} else {
		// file not cached
		f, err := os.Open(file)
		if err != nil {
			return NotFoundHandler(c)
		}

		fi, _ := f.Stat()
		if fi.IsDir() {
			//append index.html if directory detected
			file = filepath.Join(file, indexPage)
		}
		f, err = os.Open(file)
		if err != nil {
			return NotFoundHandler(c)
		}
		if fi, err = f.Stat(); err != nil {
			return f.Close()
		}
		// before sending file data to the client, create a filebuffer  for caching purposes
		raw, _ := ioutil.ReadAll(f)
		b := bytes.Buffer{}
		_, _ = b.Write(raw)
		item := new(FileBuffer)
		item.name = fi.Name()
		item.time = fi.ModTime()
		item.Buffer = b
		item.Index = 0
		fileCache.Set(initialFilePath, item)
		// add a http cache directive too
		c.response.Header().Set("Cache-Control", "public, max-age=86400") // 24h cache = 86400
		http.ServeContent(c.response, c.request, fi.Name(), fi.ModTime(), f)
		return f.Close()
	}
}

func (c *Context) Attachment(file, name string) error {
	return c.contentDisposition(file, name, "attachment")
}

func (c *Context) Inline(file, name string) error {
	return c.contentDisposition(file, name, "inline")
}

func (c *Context) contentDisposition(file, name, dispositionType string) error {
	c.response.Header().Set(HeaderContentDisposition, fmt.Sprintf("%s; filename=%q", dispositionType, name))
	return c.File(file)
}

func (c *Context) NoContent(code int) error {
	c.response.WriteHeader(code)
	return nil
}

func (c *Context) Redirect(code int, url string) error {
	if code < 300 || code > 308 {
		return ErrInvalidRedirectCode
	}
	c.response.Header().Set(HeaderLocation, url)
	c.response.WriteHeader(code)
	return nil
}

func (c *Context) Error(err error) {
	c.echo.HTTPErrorHandler(err, c)
}

func (c *Context) Echo() *Echo {
	return c.echo
}

func (c *Context) Handler() HandlerFunc {
	return c.handler
}

func (c *Context) SetHandler(h HandlerFunc) {
	c.handler = h
}

func (c *Context) Logger() Logger {
	return c.echo.Logger
}

func (c *Context) Reset(r *http.Request, w http.ResponseWriter) {
	c.request = r
	c.response.reset(w)
	c.query = nil
	c.handler = NotFoundHandler
	c.store = nil
	c.path = ""
	c.pnames = nil

	//boots data cache value
	c.isJson = false
	c.isTls = false
	c.isWs = false
	c.SchemeType = constants.Other
	c.SchemeName = ""
	c.ip = ""
	c.OnSuccessCachePolicy = 0
	// NOTE: Don't reset because it has to have length c.echo.maxParam at all times
	// c.pvalues = nil
}

//added new functions
// returns connection profile from token information
func (c *Context) ConnectionProfileSetup() (profile.ConnectionProfile, error) {
	requestProfileKeyContent := c.ReadConnectionProfileToken()
	readedProfile, err := profile.ParseConnectionProfileToken(requestProfileKeyContent)
	if err == nil {
		//save profile data
		c.profileLock.Lock()
		c.profileData = readedProfile
		c.profileLock.Unlock()
	}
	return readedProfile, err
}

// get caller eth address
func (c *Context) CallerEthAddress() string {
	c.profileLock.Lock()
	from := c.profileData.Address
	c.profileLock.Unlock()
	return from
}

// it recovers the eth client linked to it
// if peer url is provided, this peer address is used to dial
// otherwise, token information is readed in order to custom peer dial
func (c *Context) RecoverEthClientFromTokenOrPeerUrl(peerUrl string) (*ethrpc.EthRPC, string, error) {
	client := new(ethrpc.EthRPC)
	var contextId string
	// by default, peer url is used to dial
	if peerUrl == "" {
		//no peer url found, try to read from user token
		if c.profileData.RpcEndpoint == "" {
			return client, "", noPeerAddressError
		}
		contextId = c.profileData.RpcEndpoint
	} else {
		// use peer url
		contextId = peerUrl
	}
	client = ethrpc.NewDefaultRPCPtr(contextId, isDebug)
	return client, contextId, nil
}

// reads connection profile token from allowed sources
func (c *Context) ReadConnectionProfileToken() string {
	req := c.request

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
func (c *Context) ReadToken(tokenName string) string {
	req := c.request

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

func (c *Context) IsJsonRequest() bool {
	return c.isJson
}

// read user email
// this value is in jwt token setn with each request
func (c *Context) UserUuid() string {
	return c.profileData.AccountId
}

func (c *Context) AuthenticatedUserUuid() string {
	return c.UserId
}

func (c *Context) User() profile.ConnectionProfile {
	return c.profileData
}
