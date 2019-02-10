package util

import (
	"encoding/binary"
	"net"
	"strconv"
)

// converts an IP address to uint32 value
func Ip2int(ip string) uint32 {
	rawBytes := net.ParseIP(ip).To4()
	return binary.BigEndian.Uint32(rawBytes)
}

// converts an uint32 to IP address
func Int2ip(ipInt int64) string {
	// need to do two bit shifting and "0xff" masking

	b0 := strconv.FormatInt((ipInt>>24)&0xff, 10)
	b1 := strconv.FormatInt((ipInt>>16)&0xff, 10)
	b2 := strconv.FormatInt((ipInt>>8)&0xff, 10)
	b3 := strconv.FormatInt(ipInt&0xff, 10)

	return b0 + "." + b1 + "." + b2 + "." + b3
}
