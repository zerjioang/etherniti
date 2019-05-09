// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package counter

import (
	"encoding/json"
	"sync/atomic"
)

type Count32 uint32

func (c *Count32) Increment() uint32 {
	return atomic.AddUint32((*uint32)(c), 1)
}

func (c *Count32) Get() uint32 {
	return atomic.LoadUint32((*uint32)(c))
}

func (c *Count32) Bytes() []byte {
	v := atomic.LoadUint32((*uint32)(c))
	raw, err := json.Marshal(v)
	if err !=nil {
		return []byte{}
	}
	return raw
}

// constructor like function
func NewCounter32() Count32 {
	var c Count32
	return c
}
