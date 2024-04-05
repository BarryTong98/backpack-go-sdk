package service

import (
	"encoding/json"
	"strconv"

	"backpack-trade-bot/types"
)

func (c *BackpackClient) GetBalances() ([]*types.Balance, error) {
	params := make(map[string]string)
	resp, err := c.MakeAuthenticatedAPIRequest("GET", "api/v1/capital", "balanceQuery", params)
	if err != nil {
		return nil, err
	}

	var balances []*types.Balance
	err = json.Unmarshal(resp, &balances)
	if err != nil {
		return nil, err
	}

	return balances, nil
}

// GetDeposits
// Limit Maximum number to return. Default 100, maximum 1000.
// Offset Default 0.
func (c *BackpackClient) GetDeposits(limit uint16, offset int64) ([]*types.Deposit, error) {
	params := make(map[string]string)

	if limit != 0 {
		params["limit"] = strconv.Itoa(int(limit))
	}

	if offset != 0 {
		params["offset"] = strconv.FormatInt(offset, 10)
	}

	resp, err := c.MakeAuthenticatedAPIRequest("GET", "wapi/v1/capital/deposits", "depositQueryAll", params)
	if err != nil {
		return nil, err
	}

	var deposits []*types.Deposit
	err = json.Unmarshal(resp, &deposits)
	if err != nil {
		return nil, err
	}

	return deposits, nil
}

// GetDepositAddress
// string (Blockchain)
// Enum: "Solana" "Ethereum" "Polygon" "Bitcoin"
// Blockchain symbol to get a deposit address for.
func (c *BackpackClient) GetDepositAddress(blockchain string) (*types.DepositAddress, error) {
	params := map[string]string{
		"blockchain": blockchain,
	}

	resp, err := c.MakeAuthenticatedAPIRequest("GET", "wapi/v1/capital/deposit/address", "depositAddressQuery", params)
	if err != nil {
		return nil, err
	}

	var address *types.DepositAddress
	err = json.Unmarshal(resp, &address)
	if err != nil {
		return nil, err
	}

	return address, nil
}

// GetWithdrawals
// Limit Maximum number to return. Default 100, maximum 1000.
// Offset Default 0.
func (c *BackpackClient) GetWithdrawals(limit uint16, offset int64) ([]*types.Withdrawal, error) {
	params := make(map[string]string)

	if limit != 0 {
		params["limit"] = strconv.Itoa(int(limit))
	}

	if offset != 0 {
		params["offset"] = strconv.FormatInt(offset, 10)
	}

	resp, err := c.MakeAuthenticatedAPIRequest("GET", "wapi/v1/capital/withdrawals", "withdrawalQueryAll", params)
	if err != nil {
		return nil, err
	}

	var withdrawals []*types.Withdrawal
	err = json.Unmarshal(resp, &withdrawals)
	if err != nil {
		return nil, err
	}

	return withdrawals, nil
}

// RequestWithdrawals
// address: Address to withdraw to.
// blockchain: Enum: "Solana" "Ethereum" "Polygon" "Bitcoin"
// clientId: Custom client id
// quantity: Quantity to withdraw.
// symbol: Symbol of the asset to withdraw.
// twoFactorToken: Issued two factor token.
func (c *BackpackClient) RequestWithdrawals(address, blockchain, clientId, quantity, symbol, twoFactorToken string) ([]*types.Withdrawal, error) {
	params := map[string]interface{}{
		"address":    address,
		"blockchain": blockchain,
		"quantity":   quantity,
		"symbol":     symbol,
	}

	if clientId != "" {
		params["clientId"] = clientId
	}

	if twoFactorToken != "" {
		params["twoFactorToken"] = twoFactorToken
	}

	resp, err := c.MakeAuthenticatedAPIInterfaceRequest("POST", "wapi/v1/capital/withdrawals", "withdraw", params)
	if err != nil {
		return nil, err
	}

	var withdrawals []*types.Withdrawal
	err = json.Unmarshal(resp, &withdrawals)
	if err != nil {
		return nil, err
	}

	return withdrawals, nil
}
