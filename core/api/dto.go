// Copyright MethW
// SPDX-License-Identifier: Apache License 2.0

package api

import "net/http"

type Profile struct {
	Address string `json:"address"`
	Node    string `json:"node"`
}

type NewProfileRequest struct {
	Address string `json:"address" form:"address" query:"address"`
	Node    string `json:"node" form:"node" query:"node"`
}

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Details string `json:"details"`
}

func NewApiError(code int, details string) ApiError {
	ae := ApiError{}
	ae.Code = code
	ae.Message = http.StatusText(code)
	ae.Details = details
	return ae
}

type ApiResponse struct {
	Id   int `json:"id"`
	Code int `json:"code"`
	//Error error
	Message string      `json:"msg"`
	Result  interface{} `json:"result"`
}

func NewApiResponse(message string, payload interface{}) ApiResponse {
	ae := ApiResponse{}
	ae.Id = 0
	ae.Code = 200
	ae.Message = message
	ae.Result = payload
	return ae
}
