package metadata

import (
	"github.com/zerjioang/etherniti/core/modules/fastime"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type Metadata struct {
	// item creation date
	CreationDate int64 `json:"created"`
	// owner id of the creator
	Owner string `json:"owner"`
	// ip address who created the item
	Ip uint32 `json:"ip"`
}

func NewMetadata(ctx *echo.Context) *Metadata {
	mtdt := new(Metadata)
	mtdt.Ip = ctx.IntIp()
	mtdt.Owner = ctx.UserId
	mtdt.CreationDate = fastime.Now().Unix()
	return mtdt
}
