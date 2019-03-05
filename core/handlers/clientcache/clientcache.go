// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package clientcache

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo"
)

const (
	CacheOneDay   = 86400    // seconds in a day
	CacheInfinite = 31104000 // 86400 * 30 *12 // seconds in a year
	NoExpiration  = -1
)

// cache middleware function.
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
func Cached(c echo.Context, cacheValid bool, seconds uint) (int, echo.Context) {
	// add cache headers
	edit := sync.Mutex{}
	edit.Lock()
	r := c.Response()
	h := r.Header()
	if seconds > 0 {
		timeStr := strconv.Itoa(int(seconds))
		h.Set("Cache-Control", "public, max-age="+timeStr) // 24h cache = 86400
		//h.Set("Cache-Control", "private")
		if cacheValid {
			//cached item is still valid, so return a not modified
			r.Status = http.StatusOK // http.StatusNotModified
		} else {
			// cached data set as invalid, return 200 ok in order to update
			r.Status = http.StatusOK
		}
	}
	edit.Unlock()
	return r.Status, c
}

func CachedHtml(c echo.Context, cacheValid bool, seconds uint, htmlContent []byte) error {
	var code int
	code, c = Cached(c, cacheValid, seconds)
	return c.HTMLBlob(code, htmlContent)
}

func CachedJsonBlob(c echo.Context, cacheValid bool, seconds uint, data []byte) error {
	var code int
	code, c = Cached(c, cacheValid, seconds)
	return c.JSONBlob(code, data)
}

func CachedJson(c echo.Context, cacheValid bool, seconds uint, data interface{}) error {
	var code int
	code, c = Cached(c, cacheValid, seconds)
	return c.JSON(code, data)
}
