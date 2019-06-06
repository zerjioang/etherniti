// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package constants

const (
	// set system pointer size
	PointerSize    = 32 + int(^uintptr(0)>>63<<5)
	ownNodeCommand = `geth --fast --cache=1048 --testnet --unlock "0xmyaddress" --rpc --rpcapi "eth,net,web3" --rpccorsdomain '*' --rpcaddr localhost --rpcport 8545`
)
