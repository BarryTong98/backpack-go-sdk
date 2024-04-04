package types

type Order struct {
	OrderType             string `json:"orderType"`
	Id                    string `json:"id"`
	ClientId              int    `json:"clientId"`
	Symbol                string `json:"symbol"`
	Side                  string `json:"side"`
	Quantity              string `json:"quantity"`
	ExecutedQuantity      string `json:"executedQuantity"`
	QuoteQuantity         string `json:"quoteQuantity"`
	ExecutedQuoteQuantity string `json:"executedQuoteQuantity"`
	TriggerPrice          string `json:"triggerPrice"`
	TimeInForce           string `json:"timeInForce"`
	SelfTradePrevention   string `json:"selfTradePrevention"`
	Status                string `json:"status"`
	CreatedAt             int    `json:"createdAt"`
}

type ExecuteOrder struct {
	ClientId            int    `json:"clientId"`
	OrderType           string `json:"orderType"`
	PostOnly            bool   `json:"postOnly"`
	Price               string `json:"price"`
	Quantity            string `json:"quantity"`
	QuoteQuantity       string `json:"quoteQuantity"`
	SelfTradePrevention string `json:"selfTradePrevention"`
	Side                string `json:"side"`
	Symbol              string `json:"symbol"`
	TimeInForce         string `json:"timeInForce"`
	TriggerPrice        string `json:"triggerPrice"`
}

type OrderHistory struct {
	Id                  string `json:"id"`
	OrderType           string `json:"orderType"`
	Symbol              string `json:"symbol"`
	Side                string `json:"side"`
	Price               string `json:"price"`
	TriggerPrice        string `json:"triggerPrice"`
	Quantity            string `json:"quantity"`
	QuoteQuantity       string `json:"quoteQuantity"`
	TimeInForce         string `json:"timeInForce"`
	SelfTradePrevention string `json:"selfTradePrevention"`
	PostOnly            bool   `json:"postOnly"`
	Status              string `json:"status"`
}

type HistoricalFill struct {
	TradeId   int    `json:"tradeId"`
	OrderId   string `json:"orderId"`
	Symbol    string `json:"symbol"`
	Side      string `json:"side"`
	Price     string `json:"price"`
	Quantity  string `json:"quantity"`
	Fee       string `json:"fee"`
	FeeSymbol string `json:"feeSymbol"`
	IsMaker   bool   `json:"isMaker"`
	Timestamp string `json:"timestamp"`
}
