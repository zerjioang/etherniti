// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ip

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIpToUint32(t *testing.T) {

	t.Run("ip-to-int", func(t *testing.T) {
		intVal := Ip2int("10.41.132.6")
		t.Log("str ip:", strconv.Itoa(int(intVal)))
		assert.Equal(t, int(intVal), 170492934, "failed to convert ip to numeric")
	})
	t.Run("convert-uint32-low", func(t *testing.T) {
		intVal := Ip2intLow("10.41.132.6")
		t.Log("str ip:", strconv.Itoa(int(intVal)))
		assert.Equal(t, int(intVal), 170492934, "failed to convert ip to numeric")
	})
}
