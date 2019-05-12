// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cyber

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/db"
	"github.com/zerjioang/etherniti/core/modules/fastime"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

var (
	collection *db.Db
	pool       *sync.Pool
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
}

// check if http request host value is allowed or not
func Analytics(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		go processAnalytics(
			c.RealIP(),
			c.Request(),
		)
		// fordward request to next middleware
		return next(c)
	}
}

func processAnalytics(ip string, r *http.Request) {
	if !config.IsDevelopment() && collection != nil {
		// save request analytics data:
		n := fastime.Now()
		// Get item from instance
		record := pool.Get().(map[string]string)
		record["time"] = strconv.Itoa(int(n.Unix()))
		record["ip"] = ip
		record["host"] = r.Host
		record["method"] = r.Method
		record["ref"] = r.Referer()
		record["remote"] = r.RemoteAddr
		record["ua"] = r.UserAgent()
		record["uri"] = r.URL.RequestURI()
		raw := str.GetJsonBytes(record)
		// return the item to the pool
		pool.Put(record)
		// store on disk
		collection.PutKeyValue(n.NanosByte(), raw)
	}
}
