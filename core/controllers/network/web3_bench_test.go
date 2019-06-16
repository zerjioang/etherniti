// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network_test

import (
	"github.com/zerjioang/etherniti/core/eth/rpc"
	"github.com/zerjioang/etherniti/core/logger"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

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

func BenchmarkIdResolver(b *testing.B) {
	b.Run("resolve-network-id", func(b *testing.B) {
		b.Run("id-42", func(b *testing.B) {
			logger.Enabled(false)
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				ethrpc.ResolveNetworkId("42")
			}
		})
		b.Run("id-61717561", func(b *testing.B) {
			logger.Enabled(false)
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				ethrpc.ResolveNetworkId("61717561")
			}
		})
	})
	b.Run("resolve-network-id-2", func(b *testing.B) {
		b.Run("id-42", func(b *testing.B) {
			logger.Enabled(false)
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				ethrpc.ResolveNetworkId2("42")
			}
		})
		b.Run("id-61", func(b *testing.B) {
			logger.Enabled(false)
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				ethrpc.ResolveNetworkId2("61")
			}
		})
	})
}

func BenchmarkWeb3Controller(b *testing.B) {
	b.Run("get-balance", func(b *testing.B) {

		// run server
		notifier := make(chan error, 1)
		go cmd.RunServer(notifier)

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
