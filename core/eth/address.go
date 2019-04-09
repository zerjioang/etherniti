// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package eth

import (
	"context"
	"regexp"

	"github.com/zerjioang/etherniti/core/eth/fixtures"
)

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
