// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package util

import (
	"unsafe"

	"github.com/json-iterator/go"
)

var (
	empty    []byte
	json     = jsoniter.ConfigCompatibleWithStandardLibrary
	fastJson = jsoniter.ConfigFastest
)

func Bytes(data string) []byte {
	return *(*[]byte)(unsafe.Pointer(&data))
}

func ToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

func FastMarshal(data interface{}) ([]byte, error) {
	return fastJson.Marshal(&data)
}
func StdMarshal(data interface{}) ([]byte, error) {
	return json.Marshal(&data)
}
func GetJsonBytes(data interface{}) []byte {
	if data != nil {
		raw, _ := fastJson.Marshal(&data)
		return raw
	}
	return empty
}
