package service

import (
	"encoding/json"
	"strconv"

	"backpack-trade-bot/types"
)

// GetRecentTrades
// Limit the number of fills returned. Default 100, maximum 1000.
func (c *BackpackClient) GetRecentTrades(symbol string, limit uint16) ([]*types.Trade, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	addOptionalParam := func(key string, value uint16) {
		if value != 0 {
			params[key] = strconv.Itoa(int(value))
		}
	}
	addOptionalParam("limit", limit)

	resp, err := c.MakePublicAPIRequest("GET", "api/v1/trades", params)
	if err != nil {
		return nil, err
	}

	var trades []*types.Trade
	err = json.Unmarshal(resp, &trades)
	if err != nil {
		return nil, err
	}

	return trades, nil
}

// GetHistoricalTrades
// Limit the number of fills returned. Default 100, maximum 1000.
// Offset. Default 0.
func (c *BackpackClient) GetHistoricalTrades(symbol string, limit uint16, offset int64) ([]*types.Trade, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	if limit != 0 {
		params["limit"] = strconv.Itoa(int(limit))
	}

	if offset != 0 {
		params["offset"] = strconv.FormatInt(offset, 10)
	}

	resp, err := c.MakePublicAPIRequest("GET", "api/v1/trades/history", params)
	if err != nil {
		return nil, err
	}

	var trades []*types.Trade
	err = json.Unmarshal(resp, &trades)
	if err != nil {
		return nil, err
	}

	return trades, nil
}
