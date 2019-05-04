// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"bytes"
	"strconv"
	"time"

	"github.com/zerjioang/etherniti/core/eth/fastime"
	"github.com/zerjioang/etherniti/core/modules/integrity"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/shared/protocol"
)

func onNewStatusData() []byte {

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
	buffer := bufferBool.Get().(*bytes.Buffer)
	data := wrapper.Bytes(buffer)

	// Then put it back
	buffer.Reset()
	statusPool.Put(wrapper)
	bufferBool.Put(buffer)

	return data
}

func onNewIntegrityData() []byte {
	// get current date time
	millis := fastime.Now().Unix()
	timeStr := time.Unix(millis, 0).Format(time.RFC3339)
	millisStr := strconv.FormatInt(millis, 10)

	//sign message
	signMessage := "Hello from Etherniti Proxy. Today message generated at " + timeStr
	hash, signature := integrity.SignMsgWithIntegrity(signMessage)

	var wrapper protocol.IntegrityResponse
	wrapper.Message = signMessage
	wrapper.Millis = millisStr
	wrapper.Hash = hash
	wrapper.Signature = signature

	data := str.GetJsonBytes(wrapper)
	return data
}
