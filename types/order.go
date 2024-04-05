package types

type Order struct {
	OrderType             string `json:"orderType"`
	Id                    string `json:"id"`
	ClientId              uint32 `json:"clientId"`
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
	CreatedAt             int64  `json:"createdAt"`
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
	TradeId   uint32 `json:"tradeId"`
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
