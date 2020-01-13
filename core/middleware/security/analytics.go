// Copyright go-phc (https://github.com/zerjioang/go-hpc)
// SPDX-License-Identifier: Apache License 2.0

package security

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/zerjioang/go-hpc/thirdparty/echo/protocol"
	"github.com/zerjioang/go-hpc/thirdparty/echo/protocol/encoding"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/go-hpc/lib/db/badgerdb"

	"github.com/zerjioang/etherniti/shared/notifier"
	"github.com/zerjioang/go-hpc/lib/fastime"
	"github.com/zerjioang/go-hpc/lib/logger"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

var (
	collection         *badgerdb.BadgerStorage
	pool               *sync.Pool
	analyze            bool
	dbSerializer, mode = encoding.EncodingModeSelector(protocol.ModeJson)
)

func init() {
	var err error
	collection, err = badgerdb.NewCollection("", "access")
	if err != nil {
		logger.Error("failed to initialize access notifier db collection: ", err)
	}
	pool = &sync.Pool{
		New: func() interface{} {
			return map[string]string{}
		},
	}
	analyze = !config.IsDevelopment() && collection != nil
}

// add a background client http request notifier modules
func Analytics(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		go processAnalytics(
			c.RealIP(),
			c.Request(),
		)
		// forward request to next middleware
		return next(c)
	}
}

// Internal notifier
func InternalAnalytics(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// send new request event each time new http request is received
		notifier.NewProxyRequestEvent.Emit()
		// forward request to next middleware
		return next(c)
	}
}

// this method is called form a goroutine
func processAnalytics(ip string, r *http.Request) {
	if analyze {
		// save request notifier data
		n := fastime.Now()
		// Get item from instance
		record := pool.Get().(map[string]string)
		// populate the access notifier item
		record["time"] = strconv.Itoa(int(n.Unix()))
		record["ip"] = ip
		record["host"] = r.Host
		record["method"] = r.Method
		record["ref"] = r.Referer()
		record["remote"] = r.RemoteAddr
		record["ua"] = r.UserAgent()
		record["uri"] = r.URL.RequestURI()
		// serialize the item
		raw := encoding.GetBytesFromSerializer(dbSerializer, record)
		// return the item to the pool
		pool.Put(record)
		// store on disk
		storeErr := collection.SetRawKey(n.SafeBytes(), raw)
		if storeErr != nil {
			logger.Error("failed to store notifier information due to error: ", storeErr)
		}
	}
}
