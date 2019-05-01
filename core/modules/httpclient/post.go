// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package httpclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/thirdparty/gommon/log"
)

var (
	ApplicationJson = "application/json"
)

func MakePost(client *http.Client, url string, headers http.Header, data string) (json.RawMessage, error) {
	log.Info("sending request: ", data)
	if client == nil {
		client = &http.Client{
			Timeout: time.Second * 3,
			Transport: &http.Transport{
				TLSHandshakeTimeout: 3 * time.Second,
			},
		}
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header = headers
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

	log.Info("response received", str.UnsafeString(responseData))
	return responseData, nil
}
