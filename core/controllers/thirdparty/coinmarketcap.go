// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package thirdparty

import (
	"encoding/json"
	"errors"
	"sync/atomic"

	"github.com/zerjioang/etherniti/core/controllers/wrap"
	"github.com/zerjioang/etherniti/shared"
	"github.com/zerjioang/go-hpc/lib/eth/rpc/client"

	"github.com/zerjioang/etherniti/shared/constants"

	"github.com/zerjioang/go-hpc/lib/httpclient"

	"github.com/zerjioang/etherniti/core/data"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

const (
	ethPriceApi   = "https://api.coinmarketcap.com/v1/ticker/ethereum/"
	ethTickersApi = "https://api.coinmarketcap.com/v1/ticker/"
	get           = "GET"
)

var (
	/*
		coinMarketheaders = http.Header{
			"Host":clientHeaders["Host"],
			"Cookie":clientHeaders["Cookie"],
		}
	*/
	errNoResponse = errors.New("could not get eth price right now")
	none          []byte
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
	client *client.EthClient

	//cached value. concurrent safe that stores []byte
	priceCache atomic.Value
	//cached value. concurrent safe that stores []byte
	tickerCache atomic.Value
}

// constructor like function
func NewExternalController(client *client.EthClient) ExternalController {
	ctl := ExternalController{}
	ctl.client = client
	return ctl
}

func (ctl *ExternalController) coinMarketCapTickers(c *shared.EthernitiContext) error {
	v := ctl.tickerCache.Load()
	if v == nil {
		// value not set. generate and store in cache
		// generate value
		clientHeaders := c.Request().Header
		//overwrite http client configuration to send request without compression
		clientHeaders.Set("Accept-Encoding", "deflate")
		raw, err := httpclient.MakeCall(ctl.client, get, ethTickersApi, clientHeaders, none)
		if err != nil {
			return err
		} else if raw == nil {
			return errNoResponse
		} else {
			// store in cache
			// cache response for next request
			ctl.tickerCache.Store(raw)
			// return response to client
			c.OnSuccessCachePolicy = constants.CacheOneDay
			return api.SendSuccess(c, []byte("tickers"), raw)
		}
	} else {
		//value already set and stored in memory cache
		// return response to client
		return api.SendSuccess(c, data.EthTicker, v.(json.RawMessage))
	}
}
func (ctl *ExternalController) coinMarketCapEthPrice(c *shared.EthernitiContext) error {
	v := ctl.priceCache.Load()
	if v == nil {
		// value not set. generate and store in cache
		// generate value
		clientHeaders := c.Request().Header
		//overwrite http client configuration to send request without compression
		clientHeaders.Set("Accept-Encoding", "deflate")
		raw, err := httpclient.MakeCall(ctl.client, get, ethPriceApi, clientHeaders, none)
		if err != nil {
			return err
		} else if raw == nil {
			return errNoResponse
		} else {
			// store in cache
			// cache response for next request
			ctl.priceCache.Store(raw)
			// return response to client
			return api.SendSuccess(c, data.EthPrice, raw)
		}
	} else {
		//value already set and stored in memory cache
		// return response to client
		c.OnSuccessCachePolicy = constants.CacheOneDay
		return api.SendSuccess(c, data.EthPrice, v.(json.RawMessage))
	}
}

// implemented method from interface RouterRegistrable
func (ctl ExternalController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing external controller methods")
	router.GET("/eth/price", wrap.Call(ctl.coinMarketCapEthPrice))
	router.GET("/eth/ticker", wrap.Call(ctl.coinMarketCapTickers))
}
