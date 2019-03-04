// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package eth

import "testing"

func TestInfuraConnectivity(t *testing.T) {
	t.Run("get-latest-block", func(t *testing.T) {
		infuraEndpoint := "https://ropsten.infura.io/4f61378203ca4da4a6b6601bc16a22ad"
		// define the client usgin infura address
		infuraClient, err := GetEthereumClient(HttpClient, infuraEndpoint)
		if err != nil {
			t.Error("failed to get the client", err)
		} else if infuraClient == nil {
			t.Error("failed to get a valid client")
		}
		blockData, bErr := infuraClient.BlockByNumber(ctx, nil)
		if bErr != nil {
			t.Error("failed to get latest block", bErr)
		} else {
			b := blockData.NumberU64()
			// 5112131 is a known block number
			if b > 5112131 {
				t.Log("block count value", b)
			} else {
				t.Error("failed to get latest block")
			}
		}
	})
}
