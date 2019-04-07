// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ip

import (
	"encoding/binary"
	"net"
	"strconv"
)

const (
	asciiDot  uint8 = 46
	asciiZero uint8 = 48
)

// converts an IP address to uint32 value
func Ip2int(ip string) uint32 {
	rawBytes := net.ParseIP(ip).To4()
	return binary.BigEndian.Uint32(rawBytes)
}

func Ip2intLow(ip string) uint32 {
	var octets [4][3]byte
	var currentOctect uint8 = 0
	var currentOctectPos uint8 = 0
	for i := 0; i < len(ip); i++ {
		ipVal := ip[i]
		isDot := ipVal == asciiDot
		if isDot {
			//move to the next octect
			currentOctect++
			currentOctectPos = 0
		} else {
			// assign value to current octect
			octets[currentOctect][currentOctectPos] = ipVal
			currentOctectPos++
		}
	}
	// convert octects string bytes to decimal
	var octectsDecimal [4]byte
	octectsDecimal[0] = (octets[0][2]-asciiZero)*100 + (octets[0][1]-asciiZero)*10 + (octets[0][0] - asciiZero)
	octectsDecimal[1] = (octets[1][2]-asciiZero)*100 + (octets[1][1]-asciiZero)*10 + (octets[1][0] - asciiZero)
	octectsDecimal[2] = (octets[2][2]-asciiZero)*100 + (octets[2][1]-asciiZero)*10 + (octets[2][0] - asciiZero)
	octectsDecimal[3] = (octets[3][2]-asciiZero)*100 + (octets[3][1]-asciiZero)*10 + (octets[3][0] - asciiZero)
	// convert octects to uint32
	// octets[0]*256³ + octets[1]*256² + octets[2]*256¹ + octets[1]*256⁰
	var intIp uint32
	intIp = uint32(octectsDecimal[0])*16777216 + uint32(octectsDecimal[1])*65536*uint32(octectsDecimal[2])*256 + uint32(octectsDecimal[3])
	return intIp
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
