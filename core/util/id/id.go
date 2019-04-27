// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package id

import (
	"crypto/rand"
	"encoding/hex"
	"io"

	"github.com/zerjioang/etherniti/core/logger"
)

// Size of a UUID in bytes.
const Size = 16

// GenerateBytesUUID returns a UUID based on RFC 4122 returning the generated bytes
func GenerateUUIDFromEntropy() string {
	var uuidData [Size]byte
	_, err := io.ReadAtLeast(rand.Reader, uuidData[:], Size)
	if err != nil {
		logger.Error("error generating UUID. caused by: ", err)
		return ""
	}

	// variant bits; see section 4.1.1
	uuidData[8] = uuidData[8]&^0xc0 | 0x80

	// version 4 (pseudo-random); see section 4.1.3
	uuidData[6] = uuidData[6]&^0xf0 | 0x40

	var buf [36]byte

	hex.Encode(buf[0:8], uuidData[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], uuidData[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], uuidData[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], uuidData[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], uuidData[10:])

	return string(buf[:])
}