// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cns

import (
	"sync"

	"github.com/zerjioang/etherniti/core/trycatch"
)

// solc compiler wrapper using executable from image: stable-alpine 5 MB
// in order to call solidity from FS, we use the wrapper provided by go-ethereum located at
// https://github.com/ethereum/go-ethereum/blob/master/common/compiler/solidity.go

//in order to compile multiple .sol files, we need to do:
// https://gist.github.com/inovizz/1fdc2af0182584b90008e0cf2895554c

// useful 'ethereum rest api' results in google
// https://ethereum.stackexchange.com/search?q=ethereum+rest+api

// some interesting concurrent map implementations
// https://github.com/orcaman/concurrent-map/blob/master/concurrent_map.go

// hashing functions:
// * fnv:      builting golang
// * xxhash64: https://github.com/cespare/xxhash
// * xxh3:     https://github.com/dgryski/go-xxh3

// contract Name system service
type ContractNameSystem struct {
	data *sync.Map
}

type ContractInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Version     string `json:"version"`
}

func (c ContractInfo) Id() string {
	return c.Name + "-" + c.Version
}

func (c ContractInfo) Validate() trycatch.Error {
	if c.Name == "" {
		return trycatch.New("you must supply a valid contract name")
	}
	if c.Address == "" {
		return trycatch.New("you must supply a valid contract address starting with 0x")
	}
	if c.Version == "" {
		return trycatch.New("you must supply a valid contract version")
	}
	return trycatch.Nil()
}

func NewContractNameSystem() ContractNameSystem {
	cns := ContractNameSystem{}
	cns.data = new(sync.Map)
	return cns
}

func NewContractNameSystemPtr() *ContractNameSystem {
	cns := NewContractNameSystem()
	return &cns
}

func (ns *ContractNameSystem) Register(info ContractInfo) {
	//ns.data[contractName] = info
	key := info.Id()
	ns.data.Store(key, info)
}

func (ns *ContractNameSystem) Unregister(id string) {
	//delete(ns.data, id)
	ns.data.Delete(id)
}

func (ns *ContractNameSystem) Resolve(id string) (ContractInfo, bool) {
	//Address, found := ns.data[id]
	contractInfo, found := ns.data.Load(id)
	if found {
		return contractInfo.(ContractInfo), found
	}
	return ContractInfo{}, found
}