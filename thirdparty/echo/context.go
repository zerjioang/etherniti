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
	"strings"
	"sync"

	"github.com/zerjioang/etherniti/core/modules/concurrentmap"
)

var (
	errInvalidateCache = errors.New("failed to get item from internal cache, cache invalidation issues around")
)

type (
	// ContextInterface represents the Context of the current HTTP request. It holds request and
	// response objects, path, path parameters, data and registered handler.
	ContextInterface interface {
		// Request returns `*http.Request`.
		Request() *http.Request

		// SetRequest sets `*http.Request`.
		SetRequest(r *http.Request)

		// Response returns `*Response`.
		Response() *Response

		// IsTLS returns true if HTTP connection is TLS otherwise false.
		IsTLS() bool

		// IsWebSocket returns true if HTTP connection is WebSocket otherwise false.
		IsWebSocket() bool

		// Scheme returns the HTTP protocol scheme, `http` or `https`.
		Scheme() string

		// RealIP returns the client's network address based on `X-Forwarded-For`
		// or `X-Real-IP` request header.
		RealIP() string

		// Path returns the registered path for the handler.
		Path() string

		// SetPath sets the registered path for the handler.
		SetPath(p string)

		// Param returns path parameter by name.
		Param(name string) string

		// ParamNames returns path parameter names.
		ParamNames() []string

		// SetParamNames sets path parameter names.
		SetParamNames(names ...string)

		// ParamValues returns path parameter values.
		ParamValues() []string

		// SetParamValues sets path parameter values.
		SetParamValues(values ...string)

		// QueryParam returns the query param for the provided name.
		QueryParam(name string) string

		// QueryParams returns the query parameters as `url.Values`.
		QueryParams() url.Values

		// QueryString returns the URL query string.
		QueryString() string

		// FormValue returns the form field value for the provided name.
		FormValue(name string) string

		// FormParams returns the form parameters as `url.Values`.
		FormParams() (url.Values, error)

		// FormFile returns the multipart form file for the provided name.
		FormFile(name string) (*multipart.FileHeader, error)

		// MultipartForm returns the multipart form.
		MultipartForm() (*multipart.Form, error)

		// Cookie returns the named cookie provided in the request.
		Cookie(name string) (*http.Cookie, error)

		// SetCookie adds a `Set-Cookie` header in HTTP response.
		SetCookie(cookie *http.Cookie)

		// Cookies returns the HTTP cookies sent with the request.
		Cookies() []*http.Cookie

		// Get retrieves data from the Context.
		Get(key string) interface{}

		// Set saves data in the Context.
		Set(key string, val interface{})

		// Bind binds the request body into provided type `i`. The default binder
		// does it based on Content-Type header.
		Bind(i interface{}) error

		// Validate validates provided `i`. It is usually called after `ContextInterface#Bind()`.
		// Validator must be registered using `Echo#Validator`.
		Validate(i interface{}) error

		// Render renders a template with data and sends a text/html response with status
		// code. Renderer must be registered using `Echo.Renderer`.
		Render(code int, name string, data interface{}) error

		// HTML sends an HTTP response with status code.
		HTML(code int, html string) error

		// HTMLBlob sends an HTTP blob response with status code.
		HTMLBlob(code int, b []byte) error

		// String sends a string response with status code.
		String(code int, s string) error

		// JSON sends a JSON response with status code.
		JSON(code int, i interface{}) error

		// JSONPretty sends a pretty-print JSON with status code.
		JSONPretty(code int, i interface{}, indent string) error

		// JSONBlob sends a JSON blob response with status code.
		JSONBlob(code int, b []byte) error

		// JSONP sends a JSONP response with status code. It uses `callback` to construct
		// the JSONP payload.
		JSONP(code int, callback string, i interface{}) error

		// JSONPBlob sends a JSONP blob response with status code. It uses `callback`
		// to construct the JSONP payload.
		JSONPBlob(code int, callback string, b []byte) error

		// Blob sends a blob response with status code and content type.
		Blob(code int, contentType string, b []byte) error

		// Blob sends a blob response with status code and content type.
		FastBlob(code int, contentType string, b []byte) error

		// Stream sends a streaming response with status code and content type.
		Stream(code int, contentType string, r io.Reader) error

		// File sends a response with the content of the file.
		File(file string) error

		// Attachment sends a response as attachment, prompting client to save the
		// file.
		Attachment(file string, name string) error

		// Inline sends a response as inline, opening the file in the browser.
		Inline(file string, name string) error

		// NoContent sends a response with no body and a status code.
		NoContent(code int) error

		// Redirect redirects the request to a provided URL with status code.
		Redirect(code int, url string) error

		// Error invokes the registered HTTP error handler. Generally used by middleware.
		Error(err error)

		// Handler returns the matched handler by router.
		Handler() HandlerFunc

		// SetHandler sets the matched handler by router.
		SetHandler(h HandlerFunc)

		// Logger returns the `Logger` instance.
		Logger() Logger

		// Echo returns the `Echo` instance.
		Echo() *Echo

		// Reset resets the Context after request completes. It must be called along
		// with `Echo#AcquireContext()` and `Echo#ReleaseContext()`.
		// See `Echo#ServeHTTP()`
		Reset(r *http.Request, w http.ResponseWriter)
	}

	Context struct {
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
	}
)

var (
	fileCache concurrentmap.ConcurrentMap
)

func init() {
	fileCache = concurrentmap.New()
}

const (
	defaultMemory = 32 << 20 // 32 MB
	indexPage     = "index.html"
	defaultIndent = "  "
)

func (c *Context) writeContentType(value string) {
	header := c.Response().Header()
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
	return c.request.TLS != nil
}

func (c *Context) IsWebSocket() bool {
	upgrade := c.request.Header.Get(HeaderUpgrade)
	return strings.ToLower(upgrade) == "websocket"
}

func (c *Context) Scheme() string {
	// Can't use `r.Request.URL.Scheme`
	// See: https://groups.google.com/forum/#!topic/golang-nuts/pMUkBlQBDF0
	if c.IsTLS() {
		return "https"
	}
	if scheme := c.request.Header.Get(HeaderXForwardedProto); scheme != "" {
		return scheme
	}
	if scheme := c.request.Header.Get(HeaderXForwardedProtocol); scheme != "" {
		return scheme
	}
	if ssl := c.request.Header.Get(HeaderXForwardedSsl); ssl == "on" {
		return "https"
	}
	if scheme := c.request.Header.Get(HeaderXUrlScheme); scheme != "" {
		return scheme
	}
	return "http"
}

func (c *Context) RealIP() string {
	if ip := c.request.Header.Get(HeaderXForwardedFor); ip != "" {
		return strings.Split(ip, ", ")[0]
	}
	if ip := c.request.Header.Get(HeaderXRealIP); ip != "" {
		return ip
	}
	ra, _, _ := net.SplitHostPort(c.request.RemoteAddr)
	return ra
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
	http.SetCookie(c.Response(), cookie)
}

func (c *Context) Cookies() []*http.Cookie {
	return c.request.Cookies()
}

func (c *Context) Get(key string) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.store[key]
}

func (c *Context) Set(key string, val interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.store == nil {
		c.store = make(Map)
	}
	c.store[key] = val
}

func (c *Context) Bind(i interface{}) error {
	return c.echo.Binder.Bind(i, c)
}

func (c *Context) Validate(i interface{}) error {
	if c.echo.Validator == nil {
		return ErrValidatorNotRegistered
	}
	return c.echo.Validator.Validate(i)
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
	return c.Blob(code, MIMETextHTMLCharsetUTF8, b)
}

func (c *Context) String(code int, s string) (err error) {
	return c.Blob(code, MIMETextPlainCharsetUTF8, []byte(s))
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

func (c *Context) JSON(code int, i interface{}) (err error) {
	indent := ""
	if _, pretty := c.QueryParams()["pretty"]; c.echo.Debug || pretty {
		indent = defaultIndent
	}
	return c.json(code, i, indent)
}

func (c *Context) JSONPretty(code int, i interface{}, indent string) (err error) {
	return c.json(code, i, indent)
}

func (c *Context) JSONBlob(code int, b []byte) (err error) {
	return c.Blob(code, MIMEApplicationJSONCharsetUTF8, b)
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
			c.Response().Header().Set("Cache-Control", "public, max-age=86400") // 24h cache = 86400
			http.ServeContent(c.Response(), c.Request(), buffer.name, buffer.time, buffer)
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
		c.Response().Header().Set("Cache-Control", "public, max-age=86400") // 24h cache = 86400
		http.ServeContent(c.Response(), c.Request(), fi.Name(), fi.ModTime(), f)
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
	// NOTE: Don't reset because it has to have length c.echo.maxParam at all times
	// c.pvalues = nil
}
