// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package str

import (
	"encoding/json"
	"unsafe"
)

var (
	empty []byte
)

func UnsafeBytes(data string) []byte {
	return *(*[]byte)(unsafe.Pointer(&data))
}

func UnsafeString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

func StdMarshal(data interface{}) ([]byte, error) {
	return json.Marshal(&data)
}
func GetJsonBytes(data interface{}) []byte {
	if data != nil {
		raw, _ := StdMarshal(data)
		return raw
	}
	return empty
}

// converts ascii chars of a given string in lowercase
// this function is at least
// twice as fast as standard to lower function of go standard library
func ToLowerAscii(src string) string {
	rawBytes := []byte(src)
	s := len(rawBytes)
	for i := 0; i < s; i++ {
		c := &rawBytes[i]
		if *c >= 'A' && *c <= 'Z' {
			*c = *c + 32
		}
	}
	return string(rawBytes)
}
