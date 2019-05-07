// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package clientcache

import (
	"github.com/zerjioang/etherniti/shared/protocol"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

const (
	CacheOneDay   = 86400    // seconds in a day
	CacheInfinite = 31104000 // 86400 * 30 *12 // seconds in a year
	NoExpiration  = -1
)

func Cached(c *echo.Context, cacheHit bool, seconds uint) (int, *echo.Context) {
	return protocol.StatusOK, c
}

func CachedJsonBlob(c *echo.Context, cacheHit bool, seconds uint, data []byte) error {
	var code int
	code, c = Cached(c, cacheHit, seconds)
	return c.JSONBlob(code, data)
}
