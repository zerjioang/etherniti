// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"errors"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/zerjioang/etherniti/core/modules/httpclient"

	"github.com/zerjioang/etherniti/core/data"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

const (
	ethPriceApi = "https://api.coinmarketcap.com/v1/ticker/ethereum/"
)

var (
	errNoResponse = errors.New("could not get eth price right now")
)

type coinMarketCapEthPriceResponse []struct {
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	Symbol           string      `json:"symbol"`
	Rank             string      `json:"rank"`
	PriceUsd         string      `json:"price_usd"`
	PriceBtc         string      `json:"price_btc"`
	Two4HVolumeUsd   string      `json:"24h_volume_usd"`
	MarketCapUsd     string      `json:"market_cap_usd"`
	AvailableSupply  string      `json:"available_supply"`
	TotalSupply      string      `json:"total_supply"`
	MaxSupply        interface{} `json:"max_supply"`
	PercentChange1H  string      `json:"percent_change_1h"`
	PercentChange24H string      `json:"percent_change_24h"`
	PercentChange7D  string      `json:"percent_change_7d"`
	LastUpdated      string      `json:"last_updated"`
}

// token controller
type ExternalController struct {
	// http client
	client http.Client
	//cached value. concurrent safe that stores []byte
	priceCache atomic.Value
}

// constructor like function
func NewExternalController() ExternalController {
	ctl := ExternalController{}
	ctl.client = http.Client{
		Timeout: time.Second * 3,
		Transport: &http.Transport{
			TLSHandshakeTimeout: 3 * time.Second,
		},
	}
	return ctl
}

func (ctl *ExternalController) coinMarketCapEthPrice(c *echo.Context) error {
	v := ctl.priceCache.Load()
	if v == nil {
		// value not set. generate and store in cache
		// generate value
		clientHeaders := c.Request().Header
		/*coinMarketheaders := http.Header{
			"Host":clientHeaders["Host"],
			"Cookie":clientHeaders["Cookie"],
		}*/
		raw, err := httpclient.MakeCall(&ctl.client, "GET", ethPriceApi, clientHeaders, "")
		if err != nil {
			return err
		} else if raw == nil {
			return errNoResponse
		} else {
			// store in cache
			ethPriceResponse := api.ToSuccess(data.EthPrice, raw)
			// cache response for next request
			ctl.priceCache.Store(ethPriceResponse)
			// return response to client
			return api.SendSuccessBlob(c, ethPriceResponse)
		}
	} else {
		//value already set and stored in memory cache
		versionResponse := v.([]byte)
		// return response to client
		return api.SendSuccessBlob(c, versionResponse)
	}
}

// implemented method from interface RouterRegistrable
func (ctl ExternalController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing external controller methods")
	router.GET("/eth/price", ctl.coinMarketCapEthPrice)
}
