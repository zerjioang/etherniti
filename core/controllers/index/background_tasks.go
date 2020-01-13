// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package index

import (
	"strconv"
	"time"

	"github.com/zerjioang/etherniti/shared/dto"
	"github.com/zerjioang/go-hpc/lib/metrics/model"

	"github.com/zerjioang/go-hpc/thirdparty/echo/protocol/encoding"

	"github.com/zerjioang/go-hpc/common"

	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/go-hpc/lib/fastime"
	"github.com/zerjioang/go-hpc/lib/integrity"
)

var (
	fallbackSerializer = constants.FallbackSerializer
)

func onNewStatusData(serializer common.Serializer) []byte {
	if serializer == nil {
		serializer = fallbackSerializer
	}
	//get the wrapper from the pool, and cast it
	wrapper := statusPool.Get().(model.ServerStatusResponse)
	// force a new read memory
	memMonitor.ReadMemory()
	// read values
	memMonitor.ReadPtr(&wrapper)

	wrapper.Disk.All = diskMonitor.All()
	wrapper.Disk.Used = diskMonitor.Used()
	wrapper.Disk.Free = diskMonitor.Free()

	//get the buffer from the pool, and cast it
	raw := encoding.GetBytesFromSerializer(serializer, wrapper)

	// Then put it back
	wrapper.Reset()
	statusPool.Put(wrapper)

	return raw
}

func onNewIntegrityData(serializer common.Serializer) []byte {
	if serializer == nil {
		serializer = fallbackSerializer
	}
	// get current date time
	millis := fastime.Now().Unix()
	timeStr := time.Unix(millis, 0).Format(time.RFC3339)
	millisStr := strconv.FormatInt(millis, 10)

	//sign message
	signMessage := "Hello from Etherniti Proxy. Today message generated at " + timeStr
	hash, signature := integrity.SignMsgWithIntegrity(signMessage)

	wrapper := new(dto.IntegrityResponse)
	wrapper.Message = signMessage
	wrapper.Millis = millisStr
	wrapper.Hash = hash
	wrapper.Signature = signature
	return encoding.GetBytesFromSerializer(serializer, wrapper)
}
