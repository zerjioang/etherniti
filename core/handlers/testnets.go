// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

type EthereumPublicController struct {
	Web3Controller
}

type EthereumPrivateController struct {
	Web3Controller
}

const (
	ropsten       = "ropsten"
	ropstenInfura = "https://ropsten.infura.io/4f61378203ca4da4a6b6601bc16a22ad"

	rinkeby       = "rinkeby"
	rinkebyInfura = "https://rinkeby.infura.io/4f61378203ca4da4a6b6601bc16a22ad"

	kovan       = "kovan"
	kovanInfura = "https://kovan.infura.io/4f61378203ca4da4a6b6601bc16a22ad"

	mainnet       = "mainnet"
	mainnetInfura = "https://mainnet.infura.io/4f61378203ca4da4a6b6601bc16a22ad"
)

// constructor like function
func NewRopstenController() EthereumPublicController {
	ctl := EthereumPublicController{}
	ctl.Web3Controller = NewWeb3Controller()
	ctl.SetPeer(ropstenInfura)
	ctl.SetTargetName(ropsten)
	return ctl
}

// constructor like function
func NewRinkebyController() EthereumPublicController {
	ctl := EthereumPublicController{}
	ctl.Web3Controller = NewWeb3Controller()
	ctl.SetPeer(rinkebyInfura)
	ctl.SetTargetName(rinkeby)
	return ctl
}

// constructor like function
func NewKovanController() EthereumPublicController {
	ctl := EthereumPublicController{}
	ctl.Web3Controller = NewWeb3Controller()
	ctl.SetPeer(kovanInfura)
	ctl.SetTargetName(kovan)
	return ctl
}

// constructor like function
func NewMainNetController() EthereumPublicController {
	ctl := EthereumPublicController{}
	ctl.Web3Controller = NewWeb3Controller()
	ctl.SetPeer(mainnetInfura)
	ctl.SetTargetName(mainnet)
	return ctl
}

// constructor like function
func NewPrivateNetController() EthereumPrivateController {
	ctl := EthereumPrivateController{}
	ctl.Web3Controller = NewWeb3Controller()
	return ctl
}
