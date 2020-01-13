package metadata

import (
	"github.com/zerjioang/etherniti/shared"
	"github.com/zerjioang/go-hpc/lib/fastime"
)

type Metadata struct {
	// item creation date
	CreationDate int64 `json:"created"`
	// owner id of the creator
	Owner string `json:"owner"`
	// ip address who created the item
	Ip uint32 `json:"issued"`
}

func NewMetadata(ctx *shared.EthernitiContext) *Metadata {
	mtdt := new(Metadata)
	mtdt.Ip = ctx.IntIp()
	mtdt.Owner = ctx.UserId()
	mtdt.CreationDate = fastime.Now().Unix()
	return mtdt
}
