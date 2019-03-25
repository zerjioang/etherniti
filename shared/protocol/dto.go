// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package protocol

// profile model dto
type Profile struct {
	Address    string `json:"address"`
	PrivateKey string `json:"key"`
	Node       string `json:"node"`
}

// new profile request dto
type NewProfileRequest struct {

	//network id of target connection
	NetworkId uint8 `json:"networkId" form:"networkId" query:"networkId"`

	// address of the connection node: ip, domain, infura, etc
	Peer string `json:"peer" form:"peer" query:"peer"`

	//connection mode: ipc,http,rpc
	Mode string `json:"mode" form:"mode" query:"mode"`

	//connection por if required
	Port int `json:"port" form:"port" query:"port"`

	// default ethereum account for transactioning
	Address string `json:"address" form:"address" query:"address"`
	Key     string `json:"key" form:"key" query:"key"`
}

// new entropy request dto
type EntropyRequest struct {
	// size of initial entropy: 128 to 256 bits (for BIP39)
	Size uint16 `json:"size" form:"size" query:"size"`
}

// new hd wallet response dto
type EntropyResponse struct {
	Raw []byte `json:"entropy" form:"entropy" query:"entropy"`
}

// new mnemonic request dto
type NewMnemonicRequest struct {
	EntropyRequest

	// language
	Language string `json:"language" form:"language" query:"language"`
}

// new hd wallet request dto
type NewHdWalletRequest struct {
	NewMnemonicRequest
}

// new hd wallet response dto
type HdWalletResponse struct {
	MasterPrivateKey string
	MasterPublicKey  string
	Mnemonic         string
}

// api trycatch model dto
type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Details string `json:"details"`
}

// api trycatch constructor like function
func NewApiError(code int, details string) ApiError {
	ae := ApiError{}
	ae.Code = code
	ae.Message = StatusText(code)
	ae.Details = details
	return ae
}

// api response model dto
type ApiResponse struct {
	Id   int `json:"id"`
	Code int `json:"code"`
	//Error trycatch
	Message string      `json:"msg,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

// api response constructor like function
func NewApiResponse(message string, payload interface{}) ApiResponse {
	ae := ApiResponse{}
	ae.Id = 0
	ae.Code = 200
	ae.Message = message
	ae.Result = payload
	return ae
}
