package cyber

import (
	"github.com/zerjioang/etherniti/core/db"
	"github.com/zerjioang/etherniti/core/eth/fastime"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/thirdparty/echo"
	"strconv"
)

var(
	collection *db.Db
)

func init(){
	var err error
	collection, err = db.NewCollection("access")
	if err != nil {
		logger.Error("failed to initialize access analytics db collection: ", err)
	}
}
// check if http request host value is allowed or not
func Analytics(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.ContextInterface) error {
		if collection != nil {
			// save request analytics data:
			// time
			// ip
			// host
			// method
			// referer
			// uri
			// ua
			n := fastime.Now()
			r := c.Request()
			var record = map[string]string{
				"time":   strconv.Itoa(int(n.Unix())),
				"ip":     c.RealIP(),
				"host":   r.Host,
				"method": r.Method,
				"ref":    r.Referer(),
				"remote": r.RemoteAddr,
				"ua":    r.UserAgent(),
				"uri":    r.URL.RequestURI(),
			}
			collection.PutKeyValue(n.NanosByte(), str.GetJsonBytes(record))
		}
		// fordward request to next middleware
		return next(c)
	}
}
