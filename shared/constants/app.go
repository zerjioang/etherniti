// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package constants

import "github.com/zerjioang/go-hpc/thirdparty/echo/protocol/encoding/json"

const (
	//new line character
	NewLine = "\n"
	// set system pointer size
	PointerSize    = 32 + int(^uintptr(0)>>63<<5)
	ownNodeCommand = `geth --fast --cache=1048 --testnet --unlock "0xmyaddress" --rpc --rpcapi "eth,net,web3" --rpccorsdomain '*' --rpcaddr localhost --rpcport 8545`
)

var (
	// fallback data serialization method
	FallbackSerializer = json.Serialize
)
