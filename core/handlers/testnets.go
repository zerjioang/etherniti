// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

type EthereumPublicController struct {
	Web3Controller
}

// constructor like function
func NewRopstenController() EthereumPublicController {
	ctl := EthereumPublicController{}
	ctl.Web3Controller = NewWeb3Controller()
	ctl.SetPeer("https://ropsten.infura.io/4f61378203ca4da4a6b6601bc16a22ad")
	ctl.SetTargetName("ropsten")
	return ctl
}

// constructor like function
func NewRinkebyController() EthereumPublicController {
	ctl := EthereumPublicController{}
	ctl.Web3Controller = NewWeb3Controller()
	ctl.SetPeer("https://rinkeby.infura.io/4f61378203ca4da4a6b6601bc16a22ad")
	ctl.SetTargetName("rinkeby")
	return ctl
}

// constructor like function
func NewKovanController() EthereumPublicController {
	ctl := EthereumPublicController{}
	ctl.Web3Controller = NewWeb3Controller()
	ctl.SetPeer("https://kovan.infura.io/4f61378203ca4da4a6b6601bc16a22ad")
	ctl.SetTargetName("kovan")
	return ctl
}

// constructor like function
func NewMainNetController() EthereumPublicController {
	ctl := EthereumPublicController{}
	ctl.Web3Controller = NewWeb3Controller()
	ctl.SetPeer("https://mainnet.infura.io/4f61378203ca4da4a6b6601bc16a22ad")
	ctl.SetTargetName("mainnet")
	return ctl
}
