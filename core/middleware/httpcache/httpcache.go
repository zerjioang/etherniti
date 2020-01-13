package httpcache

import (
	"github.com/zerjioang/go-hpc/lib/codes"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

const (
	cacheControlServer = "Cache-Control"
	etagHeaderServer   = "ETag"
	etagHeaderClient   = "If-None-Match"
	ifModifiedSince    = "If-Modified-Since"
)

func HttpServerCache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ifModData := c.Request().Header.Get(ifModifiedSince)
		if ifModData != "" {
			//request contains If-Modified-Since information
			// todo parse and handle request successfully
			invalidateCache := false
			if invalidateCache {
				// resend new data to client
				// forward request to next middleware
				// inject etag header
				// c.Response().Header().Set(etagHeaderServer, "")
				return next(c)
			} else {
				// do not send data. keep cached version
				return c.JSONBlob(codes.StatusNotModified, nil)
			}
		} else {
			// no cache headers found.
			// forward request to next middleware
			return next(c)
		}
	}
}
