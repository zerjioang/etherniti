package util

import (
	"testing"
)

func TestIpToUint32(t *testing.T) {

	t.Run("convert-bytes", func(t *testing.T) {
		intVal := Ip2int("101.41.132.176")
		t.Log("uint32 ip:", intVal)
		if intVal != 1697219760 {
			t.Error("failed to convert ip to numeric")
		}
	})
	t.Run("convert-uint32", func(t *testing.T) {
		ipStr := Int2ip(1697219760)
		t.Log("str ip:", string(ipStr))
		if string(ipStr) != "101.41.132.176" {
			t.Error("failed to convert ip to numeric")
		}
	})
}
