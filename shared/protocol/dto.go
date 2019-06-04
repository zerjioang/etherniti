// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package protocol

import (
	"bytes"
	"encoding/json"

	"github.com/zerjioang/etherniti/core/util/str"
)

// profile model dto
type Profile struct {
	Address    string `json:"address"`
	PrivateKey string `json:"key"`
	Node       string `json:"node"`
}

// new profile request dto
type ProfileRequest struct {

	// user account sid
	AccountId string `json:"sid,omitempty" form:"sid" query:"sid"`

	// address of the connection node: ip, domain, infura, etc
	RpcEndpoint string `json:"endpoint" form:"endpoint" query:"endpoint"`

	// default ethereum account for transactioning
	Address string `json:"address,omitempty" form:"address" query:"address"`
	Key     string `json:"key,omitempty" form:"key" query:"key"`
	// source IP of the profile requester
	Source uint32 `json:"source,omitempty" form:"source" query:"source"`
	Ip     string `json:"ip,omitempty"`
}

// new entropy request dto
type EntropyRequest struct {
	// size of initial entropy: 128 to 256 bits (for BIP39)
	Size uint16 `json:"size" form:"size" query:"size"`
}

// new deploy request dto
type DeployRequest struct {
	Contract string `json:"contract"`
	Registry struct {
		Register    string `json:"register"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Version     string `json:"version"`
	} `json:"registry"`
}

// account creation response
type AccountResponse struct {
	Address string `json:"address" form:"address" query:"address"`
	Key     string `json:"key" form:"key" query:"key"`
}

// new hd wallet response dto
type EntropyResponse struct {
	Raw []byte `json:"entropy" form:"entropy" query:"entropy"`
}

// mnemonic request dto
type MnemonicRequest struct {
	EntropyRequest

	// language
	Language string `json:"language" form:"language" query:"language"`
	// Mnemonic secret
	Secret string `json:"secret" form:"secret" query:"secret"`
}

// new mnemonic request dto
type MnemonicResponse struct {
	EntropyRequest

	// language
	Language string `json:"language" form:"language" query:"language"`
	// mnemonic data
	Mnemonic string `json:"mnemonic" form:"mnemonic" query:"mnemonic"`
	// encrypted status
	IsEncrypted bool `json:"isEncrypted" form:"isEncrypted" query:"isEncrypted"`
	// Mnemonic hashed seed
	EncryptedSeed string `json:"encSeed" form:"encSeed" query:"encSeed"`
}

// hd wallet request dto
type NewHdWalletRequest struct {
	MnemonicRequest
}

// hd wallet response dto
type HdWalletResponse struct {
	MasterPrivateKey string `json:"private" form:"private" query:"private"`
	MasterPublicKey  string `json:"public" form:"public" query:"public"`
	Mnemonic         string `json:"mnemonic" form:"mnemonic" query:"mnemonic"`
}

// abi link request dto
type AbiLinkRequest json.RawMessage

// contract compilation options
type ContractCompilationOpts struct {
	// Select desired EVM version. Either homestead,
	//  tangerineWhistle, spuriousDragon, byzantium (default) or
	//  constantinople.
	Optimize           bool
	EstimateGas        bool
	GenerateAsm        bool
	GenerateOpcodes    bool
	GenerateBin        bool
	GenerateRuntimeBin bool
	GenerateAbi        bool
	GenerateHashes     bool
	OptimizeRuns       int
	EvmVersion         string
}

// contract compilation request dto
type SingleFileContractCompileRequest struct {
	ContractCompilationOpts `json:"opts" form:"opts" query:"opts"`
	Contract                string `json:"contract" form:"contract" query:"contract"`
}

// contract compilation request dto
type MultiFileContractCompileRequest struct {
	ContractCompilationOpts `json:"opts" form:"opts" query:"opts"`
	Contract                []string `json:"contract" form:"contract" query:"contract"`
}

// contract compilation response dto
type ContractCompileResponse struct {
	Code            string      `json:"code"`
	RuntimeCode     string      `json:"runtime"`
	Language        string      `json:"language"`
	LanguageVersion string      `json:"languageVersion"`
	CompilerVersion string      `json:"compilerVersion"`
	CompilerOptions string      `json:"compilerOptions"`
	AbiDefinition   interface{} `json:"abiDefinition"`
}

// api stack model dto
type ApiError struct {
	Desc []byte `json:"desc,omitempty"`
	Err  []byte `json:"error,omitempty"`
}

var (
	jsonBegin   = str.UnsafeBytes(`{`)
	jsonEnd     = str.UnsafeBytes(`}`)
	doubleQuote = str.UnsafeBytes(`"`)
	descId      = str.UnsafeBytes(`"desc":"`)
	errId       = str.UnsafeBytes(`,"error":"`)
)

func (e ApiError) Bytes(b *bytes.Buffer) []byte {
	b.Reset()
	// {"code":400,"msg":"test-error","details":""}
	b.Write(jsonBegin)
	if len(e.Desc) > 0 {
		b.Write(descId)
		b.Write(e.Desc)
		b.Write(doubleQuote)
	}
	if len(e.Err) > 0 {
		b.Write(errId)
		b.Write(e.Err)
		b.Write(doubleQuote)
	}
	b.Write(jsonEnd)
	return b.Bytes()
}

// api error constructor like function
func NewApiError(code int, details []byte) *ApiError {
	ae := ApiError{}
	ae.Desc = str.UnsafeBytes(StatusText(code))
	ae.Err = details
	return &ae
}

// api response model dto
type ApiResponse struct {
	Message []byte `json:"msg,omitempty"`
	Data    []byte `json:"data,omitempty"`
}

// api response constructor like function
func NewApiResponse(message []byte, payload []byte) *ApiResponse {
	ae := ApiResponse{}
	ae.Message = message
	ae.Data = payload
	return &ae
}

func (e ApiResponse) Bytes(b *bytes.Buffer) []byte {
	b.Reset()
	b.WriteString(`{"msg":"`)
	b.Write(e.Message)
	b.Write(doubleQuote)
	if e.Data != nil && len(e.Data) > 0 {
		b.WriteString(`,"data":`)
		b.Write(e.Data)
	}
	b.Write(jsonEnd)
	return b.Bytes()
}
