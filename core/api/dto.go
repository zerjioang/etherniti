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

// ne profile request dto
type NewProfileRequest struct {
	Address    string `json:"address" form:"address" query:"address"`
	PrivateKey string `json:"key" form:"key" query:"key"`
	Node       string `json:"node" form:"node" query:"node"`
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
