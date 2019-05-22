// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cyber

import (
	"github.com/zerjioang/etherniti/core/config"
	"net/http"
	"strconv"
	"sync"

	"github.com/zerjioang/etherniti/core/db"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/fastime"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

var (
	collection *db.BadgerStorage
	pool       *sync.Pool
	analyze    bool
)

func init() {
	var err error
	collection, err = db.NewCollection("access")
	if err != nil {
		logger.Error("failed to initialize access analytics db collection: ", err)
	}
	pool = &sync.Pool{
		New: func() interface{} {
			return map[string]string{}
		},
	}
	analyze = !config.IsDevelopment() && collection != nil
}

// check if http request host value is allowed or not
func Analytics(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		go processAnalytics(
			c.RealIP(),
			c.Request(),
		)
		// forward request to next middleware
		return next(c)
	}
}

// this method is called form a goroutine
func processAnalytics(ip string, r *http.Request) {
	if analyze {
		// save request analytics data
		n := fastime.Now()
		// Get item from instance
		record := pool.Get().(map[string]string)
		// populate the access analytics item
		record["time"] = strconv.Itoa(int(n.Unix()))
		record["ip"] = ip
		record["host"] = r.Host
		record["method"] = r.Method
		record["ref"] = r.Referer()
		record["remote"] = r.RemoteAddr
		record["ua"] = r.UserAgent()
		record["uri"] = r.URL.RequestURI()
		// serialize the item
		raw := str.GetJsonBytes(record)
		// return the item to the pool
		pool.Put(record)
		// store on disk
		storeErr := collection.Set(n.SafeBytes(), raw)
		if storeErr != nil {
			logger.Error("failed to store analytics information due to error: ", storeErr)
		}
	}
}
