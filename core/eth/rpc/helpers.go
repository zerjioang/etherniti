// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ethrpc

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/zerjioang/etherniti/core/eth/fixtures"
)

// ParseInt parse hex string value to int
func ParseInt(value string) (int, error) {
	i, err := strconv.ParseInt(strings.TrimPrefix(value, "0x"), 16, 64)
	if err != nil {
		return 0, err
	}

	return int(i), nil
}

// ParseBigInt parse hex string value to big.Int
func ParseBigInt(value string) (*big.Int, error) {
	i := new(big.Int)
	i, _ = i.SetString(value, 10)
	return i, nil
}

// Int64ToHex convert int64 to hexadecimal representation
func Int64ToHex(i int64) string {
	return "0x" + strconv.FormatInt(i, 16)
}

// IntToHex convert int to hexadecimal representation
func IntToHex(i int) string {
	return Int64ToHex(int64(i))
}

// BigToHex covert big.Int to hexadecimal representation
func BigToHex(bigInt big.Int) string {
	if bigInt.BitLen() == 0 {
		return "0x0"
	}

	return fixtures.Encode(bigInt.Bytes())
}
