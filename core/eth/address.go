// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package eth

import (
	"context"
	"regexp"

	"github.com/zerjioang/etherniti/core/eth/fixtures"
)

// an ethereum address is represented as 20 bytes
// in 1Gb (1000000000 bytes), 50.000.000 addresses can be stored
// as individual addresses.
// using a database or radix tree, b-tree or similar this space can be reduced
var (
	addressRegex = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	ctx          = context.Background()
)

// converts an string ethereum address to eth address
// addr: 0x71c7656ec7ab88b098defb751b7401b5f6d8976f
func ConvertAddress(addr string) fixtures.Address {
	return fixtures.HexToAddress(addr)
}

// check if an address is syntactically valid or not
func IsValidAddress(addr string) bool {
	return addressRegex.MatchString(addr)
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(addr string) bool {
	return addr == "0x0000000000000000000000000000000000000000"
}
