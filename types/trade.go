package types

type Trade struct {
	Id            int    `json:"id"`
	Price         string `json:"price"`
	Quantity      string `json:"quantity"`
	QuoteQuantity string `json:"quoteQuantity"`
	Timestamp     int    `json:"timestamp"`
	IsBuyerMaker  bool   `json:"isBuyerMaker"`
}
