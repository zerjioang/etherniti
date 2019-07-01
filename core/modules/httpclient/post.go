// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package httpclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
	fallbackClient  = &http.Client{
		Timeout: time.Second * 3,
		Transport: &http.Transport{
			TLSHandshakeTimeout: 3 * time.Second,
		},
	}
	br = bytes.NewReader(nil)
)

func MakePost(client *http.Client, url string, headers http.Header, content []byte) (json.RawMessage, error) {
	return MakeCall(client, postMethod, url, headers, content)
}

func MakeCall(client *http.Client, method string, url string, headers http.Header, content []byte) (json.RawMessage, error) {
	br.Reset(content)
	req, err := http.NewRequest(method, url, br)
	if err != nil {
		return nil, err
	}
	req.Header = headers
	//prepare the client to be used for requests
	if client == nil {
		client = fallbackClient
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	//responseData readed, close body
	_ = response.Body.Close()

	log.Info("response received: ", str.UnsafeString(responseData))
	return responseData, nil
}
