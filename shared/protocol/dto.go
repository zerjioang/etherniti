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

	// address of the connection node: ip, domain, infura, etc
	RpcEndpoint string `json:"endpoint" form:"endpoint" query:"endpoint"`

	// default ethereum account for transactioning
	Address string `json:"address,omitempty" form:"address" query:"address"`
	Key     string `json:"key,omitempty" form:"key" query:"key"`
	// source IP of the profile requester
	Source uint32 `json:"source,omitempty" form:"source" query:"source"`
	Ip     string `json:"ip,omitempty"`
}

// new login request dto
type LoginRequest struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"pwd" form:"pwd" query:"pwd"`
}

// new register request dto
type RegisterRequest struct {
	Username string `json:"user" form:"user" query:"user"`
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"pwd" form:"pwd" query:"pwd"`
}

// new recovery request dto
type RecoveryRequest struct {
	Email string `json:"email" form:"email" query:"email"`
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

// api trycatch model dto
type ApiError struct {
	Code    int    `json:"code"`
	Message []byte `json:"msg"`
	Details []byte `json:"details"`
}

func (e ApiError) Bytes(b *bytes.Buffer) []byte {
	b.Reset()
	// {"code":400,"msg":"test-error","details":""}
	b.Write(str.UnsafeBytes(`{"code":`))
	b.Write(str.IntToByte(e.Code))
	b.Write(str.UnsafeBytes(`,"msg":"`))
	b.Write(e.Message)
	b.Write(str.UnsafeBytes(`","details":"`))
	b.Write(e.Details)
	b.Write(str.UnsafeBytes(`"}`))
	return b.Bytes()
}

// api error constructor like function
func NewApiError(code int, details []byte) *ApiError {
	ae := ApiError{}
	ae.Code = code
	ae.Message = str.UnsafeBytes(StatusText(code))
	ae.Details = details
	return &ae
}

// api response model dto
type ApiResponse struct {
	Id   int `json:"id"`
	Code int `json:"code"`
	//Error trycatch
	Message []byte      `json:"msg,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

// api response constructor like function
func NewApiResponse(message []byte, payload interface{}) *ApiResponse {
	ae := ApiResponse{}
	ae.Id = 0
	ae.Code = 200
	ae.Message = message
	ae.Result = payload
	return &ae
}

func (e ApiResponse) Bytes(b *bytes.Buffer) []byte {
	b.Reset()
	// {"code":400,"msg":"test-error","details":""}
	b.WriteString(`{"code":`)
	b.Write(str.IntToByte(e.Code))
	b.WriteString(`,"id":`)
	b.Write(str.IntToByte(e.Id))
	b.WriteString(`,"msg":"`)
	b.Write(e.Message)
	b.WriteString(`","details":`)
	b.Write(str.GetJsonBytes(e.Result))
	b.WriteString(`}`)
	return b.Bytes()
}
