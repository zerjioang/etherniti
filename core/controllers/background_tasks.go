// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package controllers

import (
	"strconv"
	"time"

	"github.com/zerjioang/etherniti/core/modules/encoding/ioproto"
	"github.com/zerjioang/etherniti/thirdparty/echo"

	"github.com/zerjioang/etherniti/core/modules/fastime"
	"github.com/zerjioang/etherniti/core/modules/integrity"
	"github.com/zerjioang/etherniti/shared/protocol"
)

func onNewStatusData(ctx *echo.Context) []byte {
	if ctx != nil {
		//get the wrapper from the pool, and cast it
		wrapper := statusPool.Get().(protocol.ServerStatusResponse)
		// force a new read memory
		memMonitor.ReadMemory()
		// read values
		memMonitor.ReadPtr(&wrapper)

		wrapper.Disk.All = diskMonitor.All()
		wrapper.Disk.Used = diskMonitor.Used()
		wrapper.Disk.Free = diskMonitor.Free()

		//get the buffer from the pool, and cast it
		raw := ioproto.GetBytesFromSerializer(ctx.ResponseSerializer(), wrapper)

		// Then put it back
		wrapper.Reset()
		statusPool.Put(wrapper)

		return raw
	}
	return nil
}

func onNewIntegrityData(ctx *echo.Context) []byte {
	if ctx != nil {
		// get current date time
		millis := fastime.Now().Unix()
		timeStr := time.Unix(millis, 0).Format(time.RFC3339)
		millisStr := strconv.FormatInt(millis, 10)

		//sign message
		signMessage := "Hello from Etherniti Proxy. Today message generated at " + timeStr
		hash, signature := integrity.SignMsgWithIntegrity(signMessage)

		wrapper := new(protocol.IntegrityResponse)
		wrapper.Message = signMessage
		wrapper.Millis = millisStr
		wrapper.Hash = hash
		wrapper.Signature = signature
		return ioproto.GetBytesFromSerializer(ctx.ResponseSerializer(), wrapper)
	}
	return nil
}
