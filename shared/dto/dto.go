// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package dto

import (
	"encoding/json"
	"errors"

	"github.com/zerjioang/go-hpc/lib/codes"

	"github.com/zerjioang/go-hpc/util/str"
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

// ethereum transaction required data model dto
type TransactionRequest struct {
	// the address who makes the tx
	From string `json:"from,omitempty"`
	Auth struct {
		// in case of being using a node account that needs to be unlocked
		UnlockPassword string `json:"unlockPassword,omitempty"`
		// in the case of signing the request in external application or device such as cold wallets
		OfflineSignature string `json:"offlineSignature"`
		// in case of not providing a signature and providing private key
		PrivateKey string `json:"privateKey"`
	} `json:"auth,omitempty"`
}

// new deploy request dto
type DeployRequest struct {
	// optional transaction data required for signing and validation
	Tx       TransactionRequest `json:"tx"`
	Contract string             `json:"contract"`
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

/*
{
	"evm-version": "evm",
	"optimize": {
		"enabled": true,
		"runs": 200
	},
	"optimize-yul": false,
	"gas": true,
	"assemble": false,
	"yul": false,
	"strict-assembly": false,
	"machine": "",
	"link": false,
	"metadata-literal": false,
	"allow-paths": false,
	"report": {
		"ast": false,
		"asm": false,
		"opcodes": false,
		"bin": true,
		"bin-runtime": true,
		"abi": false,
		"ir": false,
		"hashes": false,
		"userdoc": false,
		"devdoc": false,
		"metadata": false
	}
}
*/

type ContractCompilationOpts struct {
	// Select desired EVM version. Either homestead,
	//  tangerineWhistle, spuriousDragon, byzantium (default) or
	//  constantinople.
	EvmVersion      string       `json:"evm-version"`
	Optimize        OptimizeOpts `json:"optimize"`
	OptimizeYul     bool         `json:"optimize-yul"`
	Gas             bool         `json:"gas"`
	Assemble        bool         `json:"assemble"`
	Yul             bool         `json:"yul"`
	StrictAssembly  bool         `json:"strict-assembly"`
	Machine         string       `json:"machine"`
	Link            bool         `json:"link"`
	MetadataLiteral bool         `json:"metadata-literal"`
	AllowPaths      bool         `json:"allow-paths"`
	Report          ReportOpts   `json:"report"`
}
type OptimizeOpts struct {
	Enabled bool `json:"enabled"`
	Runs    int  `json:"runs"`
}
type ReportOpts struct {
	Ast        bool `json:"ast"`
	Asm        bool `json:"asm"`
	Opcodes    bool `json:"opcodes"`
	Bin        bool `json:"bin"`
	BinRuntime bool `json:"bin-runtime"`
	Abi        bool `json:"abi"`
	Ir         bool `json:"ir"`
	Hashes     bool `json:"hashes"`
	Userdoc    bool `json:"userdoc"`
	Devdoc     bool `json:"devdoc"`
	Metadata   bool `json:"metadata"`
}

// contract compilation request dto
type SingleFileContractCompileRequest struct {
	Opts     ContractCompilationOpts `json:"opts" form:"opts" query:"opts"`
	Contract string                  `json:"contract" form:"contract" query:"contract"`
}

// contract compilation request dto
type MultiFileContractCompileRequest struct {
	Opts     ContractCompilationOpts `json:"opts" form:"opts" query:"opts"`
	Contract []string                `json:"contract" form:"contract" query:"contract"`
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
	Desc string `json:"desc,omitempty" msg:"desc"`
	Err  string `json:"error,omitempty" msg:"error"`
}

// api error constructor like function
func NewApiError(code codes.HttpStatusCode, details string) *ApiError {
	ae := ApiError{}
	ae.Desc = code.Text()
	ae.Err = details
	return &ae
}

// api response model dto
type ApiResponse struct {
	Message string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// api response constructor like function
func NewApiResponse(message []byte, payload interface{}) *ApiResponse {
	ae := ApiResponse{}
	ae.Message = str.UnsafeString(message)
	ae.Data = payload
	return &ae
}

type EthSignatureParseRequest struct {
	//Eth message signature data encoded as hex string starting with 0x
	Signature string `json:"signature"`
}

func (req *EthSignatureParseRequest) Validate() error {
	if req.Signature == "" || len(req.Signature) == 0 {
		return errors.New("empty signature data provided")
	}
	//todo check if signature string is hex valid
	//todo check if signature string starts with 0x
	return nil
}

type EthSignRequest struct {
	From       string `json:"from"`
	To         string `json:"to"`
	PrivateKey string `json:"private"`
}

func (req *EthSignRequest) Validate() error {
	//todo check if signature string is hex valid
	//todo check if signature string starts with 0x
	if req.PrivateKey == "" {
		return errors.New("private key is missing in request")
	}
	//todo check if signature string is hex valid
	//todo check if signature string starts with 0x
	if req.To == "" {
		return errors.New("from field is missing in request")
	}
	return nil
}

type EthSignatureParseResponse struct {
	R string `json:"r"`
	S string `json:"s"`
	V string `json:"v"`
}

type EthSha3Request struct {
	// data to be signed with sha3
	Data string `json:"data"`
}

func (req *EthSha3Request) Validate() error {
	if req.Data == "" || len(req.Data) == 0 {
		return errors.New("empty data provided for sha3 generation")
	}
	//todo check if signature string is hex valid
	//todo check if signature string starts with 0x
	return nil
}

type DbStorageRequest struct {
	Database string `json:"database"` //Database name.
	Key      string `json:"key"`      // Key name.
	Value    string `json:"value"`    // String to store.
}

func (req *DbStorageRequest) Validate() error {
	if req.Database == "" {
		return errors.New("database name cannot be empty")
	}
	if req.Key == "" {
		return errors.New("item key cannot be empty")
	}
	if req.Value == "" {
		return errors.New("item value cannot be empty")
	}
	return nil
}

type DbStorageResponseRequest struct {
	Value string `json:"value"`
}
