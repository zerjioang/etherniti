// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cns

import (
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/stack"
	"github.com/zerjioang/etherniti/core/util/id"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/shared/mixed"
	"github.com/zerjioang/etherniti/thirdparty/echo"
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

// contract name system service
type ContractInfo struct {
	// implement interface to be a rest-db-crud able struct
	mixed.DatabaseObjectInterface `json:"_,omitempty"`
	// unique project identifier used for database storage
	Uuid string `json:"uuid"`

	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Version     string `json:"version"`
}

// implementation of interface DatabaseObjectInterface
func (model ContractInfo) Key() []byte {
	return str.UnsafeBytes(model.Uuid)
}
func (model ContractInfo) Value() []byte {
	return str.GetJsonBytes(model)
}
func (model ContractInfo) New() mixed.DatabaseObjectInterface {
	return NewEmptyContractInfo()
}

// custom validation logic for read operation
// return nil if everyone can read
func (model ContractInfo) CanRead(context *echo.Context, key string) error {
	// todo check if current ContractInfo id belongs to current user
	return nil
}

// custom validation logic for update operation
// return nil if everyone can update
func (model ContractInfo) CanUpdate(context *echo.Context, key string) error {
	// todo check if current ContractInfo id belongs to current user
	return nil
}

// custom validation logic for delete operation
// return nil if everyone can delete
func (model ContractInfo) CanDelete(context *echo.Context, key string) error {
	// todo check if current ContractInfo id belongs to current user
	return nil
}

// custom validation logic for write operation
// return nil if everyone can write
func (model ContractInfo) CanWrite(context *echo.Context) error {
	return nil
}

// custom validation logic for list operation
// return nil if everyone can list
func (model ContractInfo) CanList(context *echo.Context) error {
	// todo check if current ContractInfo id belongs to current user
	return nil
}

func (model ContractInfo) Bind(context *echo.Context) (mixed.DatabaseObjectInterface, stack.Error) {
	if err := context.Bind(&model); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return nil, stack.Ret(err)
	}
	return nil, data.ErrBind
}

func NewEmptyContractInfo() ContractInfo {
	return ContractInfo{}
}

func NewDBContractInfo() mixed.DatabaseObjectInterface {
	return NewEmptyContractInfo()
}

func NewContractInfo(name string) *ContractInfo {
	p := new(ContractInfo)
	p.Uuid = id.GenerateIDString().String()
	p.Name = name
	return p
}
