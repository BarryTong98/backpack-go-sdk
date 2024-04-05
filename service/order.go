package service

import (
	"encoding/json"
	"strconv"

	"backpack-trade-bot/types"
)

func (c *BackpackClient) GetOpenOrder(clientId uint32, orderId, symbol string) (*types.Order, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	if clientId != 0 {
		params["clientId"] = strconv.Itoa(int(clientId))
	}

	if orderId != "" {
		params["orderId"] = orderId
	}

	resp, err := c.MakeAuthenticatedAPIRequest("GET", "api/v1/order", "orderQuery", params)
	if err != nil {
		return nil, err
	}

	var order *types.Order
	err = json.Unmarshal(resp, &order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// ExecuteOrder
// clientId
// integer <uint32>
// Custom order id.
//
// orderType
// required
// string
// Enum: "Market" "Limit"
// The type of order.
//
// postOnly
// boolean
// Whether the order is post only or not
//
// price
// string <decimal>
// The order price if this is a limit order
//
// quantity
// string <decimal>
// The order quantity. Market orders must specify either a quantity or quoteQuantity. All other order types must specify a quantity.
//
// quoteQuantity
// string <decimal>
// The maximum amount of the quote asset to spend (Ask) or receive (Bid) for market orders. This is used for reverse market orders. The order book will execute a quantity as close as possible to the notional value of quote_quantity.
//
// selfTradePrevention
// string
// Enum: "RejectTaker" "RejectMaker" "RejectBoth" "Allow"
// Self trade prevention describes what should happen if the order attempts to fill against another order from the same account or trade group.
//
// side
// required
// string
// Enum: "Bid" "Ask"
// Bid -> Buy token, Ask -> Sell token
//
// symbol
// required
// string
// The market for the order.
//
// timeInForce
// string
// Enum: "GTC" "IOC" "FOK"
// The time in force setting for an order.
//
// triggerPrice
// string <decimal>
// Trigger price if this is a conditional order.
//
// Example: ExecuteOrder(0, "Limit", false, "0.0002", "100000", "", "", "Bid", "WEN_USDC", "", "")
func (c *BackpackClient) ExecuteOrder(clientId uint32, orderType string, postOnly bool, price, quantity, quoteQuantity, selfTradePrevention, side, symbol, timeInForce, triggerPrice string) (*types.ExecuteOrder, error) {
	params := map[string]interface{}{
		"orderType": orderType,
		"side":      side,
		"symbol":    symbol,
		"postOnly":  postOnly,
	}

	if clientId != 0 {
		params["clientId"] = clientId
	}

	if price != "" {
		params["price"] = price
	}

	if quantity != "" {
		params["quantity"] = quantity
	}

	if quoteQuantity != "" {
		params["quoteQuantity"] = quoteQuantity
	}

	if selfTradePrevention != "" {
		params["selfTradePrevention"] = selfTradePrevention
	}

	if timeInForce != "" {
		params["timeInForce"] = timeInForce
	}

	if triggerPrice != "" {
		params["triggerPrice"] = triggerPrice
	}

	resp, err := c.MakeAuthenticatedAPIInterfaceRequest("POST", "api/v1/order", "orderExecute", params)
	if err != nil {
		return nil, err
	}

	var order *types.ExecuteOrder
	err = json.Unmarshal(resp, &order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (c *BackpackClient) CancelOpenOrder(clientId uint32, orderId, symbol string) (*types.Order, error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}

	if clientId != 0 {
		params["clientId"] = clientId
	}

	if orderId != "" {
		params["orderId"] = orderId
	}

	resp, err := c.MakeAuthenticatedAPIInterfaceRequest("DELETE", "api/v1/order", "orderCancel", params)
	if err != nil {
		return nil, err
	}
	var order *types.Order
	err = json.Unmarshal(resp, &order)
	if err != nil {
		return nil, err
	}

	return order, nil

}

func (c *BackpackClient) GetOpenOrders(symbol string) ([]*types.Order, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	resp, err := c.MakeAuthenticatedAPIRequest("GET", "api/v1/orders", "orderQueryAll", params)
	if err != nil {
		return nil, err
	}

	var orders []*types.Order
	err = json.Unmarshal(resp, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (c *BackpackClient) CancelOpenOrders(symbol string) ([]*types.Order, error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}

	resp, err := c.MakeAuthenticatedAPIInterfaceRequest("DELETE", "api/v1/orders", "orderCancelAll", params)
	if err != nil {
		return nil, err
	}

	var orders []*types.Order
	err = json.Unmarshal(resp, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
