package service

import (
	"backpack-trade-bot/types"
	"encoding/json"
	"strconv"
)

func (c *BackpackClient) GetAssets() ([]*types.Asset, error) {
	params := make(map[string]string)

	resp, err := c.MakePublicAPIRequest("GET", "api/v1/assets", params)
	if err != nil {
		return nil, err
	}
	var assets []*types.Asset
	err = json.Unmarshal(resp, &assets)
	if err != nil {
		return nil, err
	}

	return assets, nil
}

func (c *BackpackClient) GetMarkets() ([]*types.Market, error) {
	params := make(map[string]string)

	resp, err := c.MakePublicAPIRequest("GET", "api/v1/markets", params)
	if err != nil {
		return nil, err
	}
	var markets []*types.Market
	err = json.Unmarshal(resp, &markets)
	if err != nil {
		return nil, err
	}

	return markets, nil
}

func (c *BackpackClient) GetTicker(symbol string) (*types.Ticker, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	resp, err := c.MakePublicAPIRequest("GET", "api/v1/ticker", params)
	if err != nil {
		return nil, err
	}
	var ticker *types.Ticker
	err = json.Unmarshal(resp, &ticker)
	if err != nil {
		return nil, err
	}

	return ticker, nil
}

func (c *BackpackClient) GetTickers() ([]*types.Ticker, error) {
	params := make(map[string]string)

	resp, err := c.MakePublicAPIRequest("GET", "api/v1/tickers", params)
	if err != nil {
		return nil, err
	}
	var tickers []*types.Ticker
	err = json.Unmarshal(resp, &tickers)
	if err != nil {
		return nil, err
	}

	return tickers, nil
}

func (c *BackpackClient) GetDepth(symbol string) (*types.Depth, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	resp, err := c.MakePublicAPIRequest("GET", "api/v1/depth", params)
	if err != nil {
		return nil, err
	}
	var depth *types.Depth
	err = json.Unmarshal(resp, &depth)
	if err != nil {
		return nil, err
	}

	return depth, nil
}

// GetKLines
// Interval: Enum: "1m" "3m" "5m" "15m" "30m" "1h" "2h" "4h" "6h" "8h" "12h" "1d" "3d" "1w" "1month"
func (c *BackpackClient) GetKLines(symbol, interval string, startTime, endTime int64) ([]*types.KLines, error) {
	params := map[string]string{
		"symbol":   symbol,
		"interval": interval,
	}

	addOptionalParam := func(key string, value int64) {
		if value != -1 {
			params[key] = strconv.FormatInt(value, 10)
		}
	}

	addOptionalParam("startTime", startTime)
	addOptionalParam("endTime", endTime)

	resp, err := c.MakePublicAPIRequest("GET", "api/v1/klines", params)
	if err != nil {
		return nil, err
	}
	var kLines []*types.KLines
	err = json.Unmarshal(resp, &kLines)
	if err != nil {
		return nil, err
	}

	return kLines, nil

}
