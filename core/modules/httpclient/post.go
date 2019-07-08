// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package httpclient

import (
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"

	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/thirdparty/gommon/log"
)

const (
	postMethod = "POST"
)

var (
	ApplicationJSON = "application/json"
	fallbackClient  = &fasthttp.Client{
		ReadTimeout: time.Second * 3,
		WriteTimeout: time.Second * 3,
		WriteBufferSize: 2048,
		ReadBufferSize: 2048,
	}
	br = bytes.NewReader(nil)
)

func MakePost(client *fasthttp.Client, url string, headers http.Header, content []byte) (json.RawMessage, error) {
	return MakeCall(client, postMethod, url, headers, content)
}

func MakeCall(client *fasthttp.Client, method string, url string, headers http.Header, content []byte) (json.RawMessage, error) {
	br.Reset(content)
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url) //set URL
	req.Header.SetMethodBytes([]byte(method)) //set method mode
	req.SetBody(content) //set body

	res := fasthttp.AcquireResponse()

	//prepare the client to be used for requests
	if client == nil {
		client = fallbackClient
	}
	err := client.Do(req, res)
	if err != nil {
		return nil, err
	}

	responseData := res.Body()

	log.Info("response received: ", str.UnsafeString(responseData))
	return responseData, nil
}
