// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package e2e

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/core/cmd"
	"github.com/zerjioang/etherniti/shared/constants"
)

const (
	/*
		{
		  "endpoint": "HTTP://127.0.0.1:7545",
		  "address": "0x3DEB1894DC2d3e1B4a073f520e516C2DF6f45B88"
		}
	*/
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbmRwb2ludCI6IkhUVFA6Ly8xMjcuMC4wLjE6NzU0NSIsImFkZHJlc3MiOiIweDNERUIxODk0REMyZDNlMUI0YTA3M2Y1MjBlNTE2QzJERjZmNDVCODgiLCJzb3VyY2UiOjIxMzA3MDY0MzMsInZlcnNpb24iOiIwLjAuNiIsImV4cCI6MTU1NjUwMjE1MCwianRpIjoiZmQ0NjQzZmItNjk4My00MzI1LWIzNzctMTJmOWRmZDY2M2IxIiwiaWF0IjoxNTU2MTQyMTUwLCJpc3MiOiJldGhlcm5pdGkub3JnIiwibmJmIjoxNTU2MTQyMTUwLCJ2YWxpZGl0eSI6ZmFsc2V9.bwOFdtZBJ6oLhQtNwo_IQTPnOMf2edQGQfDeKQEhNuI"
)

func TestWeb3ControllerEndToEnd(t *testing.T) {
	t.Run("getBalance", func(t *testing.T) {
		// run server
		go cmd.RunServer()

		//wait 1 second
		time.Sleep(time.Second * 1)

		//make http request
		client := &http.Client{}
		req, _ := http.NewRequest("GET", "http://localhost:8080/v1/private/balance/0x3DEB1894DC2d3e1B4a073f520e516C2DF6f45B88", nil)
		req.Header.Set(constants.HttpProfileHeaderkey, token)
		resp, err := client.Do(req)

		assert.Nil(t, err)
		assert.NotNil(t, resp)

		responseData, readErr := ioutil.ReadAll(resp.Body)
		assert.Nil(t, readErr)
		assert.NotNil(t, responseData)

		//responseData readed, close body
		_ = resp.Body.Close()

		t.Log(string(responseData))
	})
}
