// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package eth

import (
	"regexp"

	"github.com/zerjioang/etherniti/core/eth/fixtures/common"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/eth/fixtures"
)

const (
	// length of an ethereum address
	AddressLengthBytes = 20
	// length of ethereum address hexadecimal encoded
	AddressLengthHexEncoded = AddressLengthBytes * 2
	// length of ethereum address hexadecimal encoded and prefixed with '0x'
	AddressLengthHexEncodedWithPrefix = 2 + AddressLengthHexEncoded
)

// an ethereum address is represented as 20 bytes
// in 1Gb (1000000000 bytes), 50.000.000 addresses can be stored
// as individual addresses.
// using a database or radix tree, b-tree or similar this space can be reduced
var (
	addressRegex = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
)

// converts an string ethereum address to eth address
// addr: 0x71c7656ec7ab88b098defb751b7401b5f6d8976f
func ConvertAddress(addr string) fixtures.Address {
	return fixtures.HexToAddress(addr)
}

// check if an address is syntactically valid or not
// example address: 0x71c7656ec7ab88b098defb751b7401b5f6d8976f
func IsValidAddress(addr string) bool {
	return addressRegex.MatchString(addr)
}

// check if an address is syntactically valid or not
// example address: 0x71c7656ec7ab88b098defb751b7401b5f6d8976f
func IsValidAddressLow(addr string) bool {
	if len(addr) == 42 {
		raw := str.UnsafeBytes(addr)
		// for bound checks speed up
		_ = raw[41]
		// check adress begin (0x)
		x := raw[1]
		zero := raw[0]
		if zero == '0' && (x == 'x' || x == 'X') {
			//check address body
			for i := 2; i < 40; i++ {
				c := raw[i]
				if !((c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F') || (c >= '0' && c <= '9')) {
					return false
				}
			}
			return true
		}
	}
	return false
}

// check if the input string represents a valid block number
// allowed block numbers are:
// The following options are possible for the defaultBlock parameter:
//    HEX String - an integer block number
//    String "earliest" for the earliest/genesis block
//    String "latest" - for the latest mined block
//    String "pending" - for the pending state/transactions
func IsValidBlockNumber(blkStr string) bool {
	if blkStr == "" {
		return false
	} else if blkStr == "earliest" {
		return true
	} else if blkStr == "latest" {
		return true
	} else if blkStr == "pending" {
		return true
	} else {
		//validate input is valid hex string
		hasHexPrefix := len(blkStr) >= 2 && blkStr[0] == '0' && (blkStr[1] == 'x' || blkStr[1] == 'X')
		if hasHexPrefix {
			return common.IsHex(blkStr[2:])
		}
		return false
	}
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(addr string) bool {
	if len(addr) == AddressLengthHexEncodedWithPrefix {
		return addr == "0x0000000000000000000000000000000000000000"
	}
	return false
}

// check whether given byte value is hexadecimal or not
func IsXdigit(c byte) bool {
	return (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F') || (c >= '0' && c <= '9')
}
