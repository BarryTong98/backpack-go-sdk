package service

import (
	"encoding/json"
	"strconv"

	"backpack-trade-bot/types"
)

func (c *BackpackClient) GetOrderHistory(orderId, symbol string, offset, limit int64) ([]*types.OrderHistory, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	if orderId != "" {
		params["orderId"] = orderId
	}

	if offset != 0 {
		params["offset"] = strconv.FormatInt(offset, 10)
	}

	if limit != 0 {
		params["limit"] = strconv.FormatInt(limit, 10)
	}

	resp, err := c.MakeAuthenticatedAPIRequest("GET", "wapi/v1/history/orders", "orderHistoryQueryAll", params)
	if err != nil {
		return nil, err
	}

	var orders []*types.OrderHistory
	err = json.Unmarshal(resp, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// GetFillHistory
// orderId: Filter to the given order.
//
// from: Filter to minimum time (milliseconds).
//
// to: Filter to maximum time (milliseconds).
//
// symbol: Filter to the given symbol.
//
// limit: Maximum number to return. Default 100, maximum 1000.
//
// offset: Offset. Default 0.
func (c *BackpackClient) GetFillHistory(orderId, symbol string, from, to, offset, limit int64) ([]*types.HistoricalFill, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	if orderId != "" {
		params["orderId"] = orderId
	}

	if from != 0 {
		params["from"] = strconv.FormatInt(from, 10)
	}

	if to != 0 {
		params["to"] = strconv.FormatInt(to, 10)
	}

	if offset != 0 {
		params["offset"] = strconv.FormatInt(offset, 10)
	}

	if limit != 0 {
		params["limit"] = strconv.FormatInt(limit, 10)
	}

	resp, err := c.MakeAuthenticatedAPIRequest("GET", "wapi/v1/history/fills", "fillHistoryQueryAll", params)
	if err != nil {
		return nil, err
	}

	var fills []*types.HistoricalFill
	err = json.Unmarshal(resp, &fills)
	if err != nil {
		return nil, err
	}

	return fills, nil
}
