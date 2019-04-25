// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network_test

import (
	"github.com/zerjioang/etherniti/core/cmd"
	"github.com/zerjioang/etherniti/shared/constants"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
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

func BenchmarkWeb3Controller(b *testing.B) {
	b.Run("get-balance", func(b *testing.B) {

		// run server
		go cmd.RunServer()

		//wait 1 second
		time.Sleep(time.Second * 1)

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = makeRequest()
		}
	})
}

func makeRequest() []byte {
	//make http request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8080/v1/private/balance/0x3DEB1894DC2d3e1B4a073f520e516C2DF6f45B88", nil)
	req.Header.Set(constants.HttpProfileHeaderkey, token)
	resp, _ := client.Do(req)

	responseData, _ := ioutil.ReadAll(resp.Body)

	//responseData readed, close body
	_ = resp.Body.Close()

	return responseData
}