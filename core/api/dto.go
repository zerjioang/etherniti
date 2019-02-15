// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

import "net/http"

// profile model dto
type Profile struct {
	Address    string `json:"address"`
	PrivateKey string `json:"key"`
	Node       string `json:"node"`
}

// new profile request dto
type NewProfileRequest struct {
	// address of the connection node: ip, domain, infura, etc
	NodeAddress string `json:"peer" form:"peer" query:"peer"`

	//connection mode: ipc,http,rpc
	Mode string `json:"mode" form:"mode" query:"mode"`

	//connection por if required
	Port int `json:"port" form:"port" query:"port"`

	// default ethereum account for transactioning
	Address string `json:"address" form:"address" query:"address"`
}

// api error model dto
type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Details string `json:"details"`
}

// api error constructor like function
func NewApiError(code int, details string) ApiError {
	ae := ApiError{}
	ae.Code = code
	ae.Message = http.StatusText(code)
	ae.Details = details
	return ae
}

// api response model dto
type ApiResponse struct {
	Id   int `json:"id"`
	Code int `json:"code"`
	//Error error
	Message string      `json:"msg"`
	Result  interface{} `json:"result"`
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
