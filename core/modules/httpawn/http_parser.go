package httpawn

// Definitions of HTTP methods according to
// https://tools.ietf.org/html/rfc7231
type httpMethod uint8

const (
	GET httpMethod = iota
	HEAD
	POST
	PUT
	DELETE
	CONNECT
	OPTIONS
	TRACE
	PATCH
	UNKNOWN
)

const (
	httpReqMinSize = 12

	commonHttpResponsePreamble = `HTTP/1.1 200 OK
Date: Wed, 18 Sep 2019 17:31:26 GMT
Accept: */*
Content-Type: text/html; charset=utf-8
Server: httpawn
Last-Modified: Fri, 08 Mar 2019 10:50:51 GMT
Access-Control-Allow-Origin: *
Expires: Wed, 25 Sep 2019 17:31:26 GMT
Cache-Control: max-age=604800
X-Proxy-Cache: MISS
Vary: Negotiate,Accept-Encoding
Last-Modified: Tue, 10 Sep 2019 17:49:11 GMT
ETag: "3ce5a1-4c0cc-59236859b73c0;592d73264bbe0"
Accept-Ranges: bytes
Content-Encoding: plain
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block
X-Content-Type-Options: nosniff
Strict-Transport-Security: max-age=7776000; includeSubDomains; preload
Expect-Ct: max-age=604800, report-uri="https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct"
X-Firefox-Spdy: h2
`

	invalidHttpRequest = `Invalid HTTP request`
	invalidHttpMethod  = `Invalid HTTP Method in request`
)

var (
	commonHttpResponsePreambleRaw = []byte(commonHttpResponsePreamble)
	invalidHttpRequestRaw         = []byte(invalidHttpRequest)
	invalidHttpMethodRaw          = []byte(invalidHttpMethod)
)

// @return body content as []byte
// @return response headers as as []byte
func processHttpRequest(r *Router, req *socketRequest) ([]byte, []byte) {
	srcIp := req.client.RemoteAddr()
	raw := req.raw
	//0. validate the minimum length of the http request
	// 3chars (http method)
	// space
	// 1char ( / as url)
	// raw minimum size is 5 bytes
	// example: GET / HTTP/1.1
	// example: GET / HTTP/2
	size := len(raw)
	if raw == nil || size < httpReqMinSize {
		//return content no valid message
		return commonHttpResponsePreambleRaw, invalidHttpRequestRaw
	}
	//1. first we detect the type of method of the request
	// allowed methods are:
	// GET, HEAD, POST, PUT, DELETE, CONNECT, OPTIONS, TRACE, PATCH
	method, start := resolveHttpMethod(raw)
	if method != UNKNOWN {
		//2. extract defined path in request using provided start offset
		path := resolveHttpPath(raw, start)
		body := r.Execute(srcIp, method, path)
		return commonHttpResponsePreambleRaw, body
	} else {
		// send invalid http method message
		// specified method is not defined in RFC
		return commonHttpResponsePreambleRaw, invalidHttpMethodRaw
	}
}
