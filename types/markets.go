package types

type Asset struct {
	Symbol string `json:"symbol"`
	Tokens []struct {
		Blockchain        string      `json:"blockchain"`
		DepositEnabled    bool        `json:"depositEnabled"`
		MaximumWithdrawal interface{} `json:"maximumWithdrawal"`
		MinimumDeposit    string      `json:"minimumDeposit"`
		MinimumWithdrawal string      `json:"minimumWithdrawal"`
		WithdrawEnabled   bool        `json:"withdrawEnabled"`
		WithdrawalFee     string      `json:"withdrawalFee"`
	} `json:"tokens"`
}

type Ticker struct {
	Symbol             string `json:"symbol"`
	FirstPrice         string `json:"firstPrice"`
	LastPrice          string `json:"lastPrice"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	High               string `json:"high"`
	Low                string `json:"low"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	Trades             int    `json:"trades"`
}

type Market struct {
	Symbol      string `json:"symbol"`
	BaseSymbol  string `json:"baseSymbol"`
	QuoteSymbol string `json:"quoteSymbol"`
	Filters     struct {
		Price struct {
			MinPrice string `json:"minPrice"`
			MaxPrice string `json:"maxPrice"`
			TickSize string `json:"tickSize"`
		} `json:"price"`
		Quantity struct {
			MinQuantity string `json:"minQuantity"`
			MaxQuantity string `json:"maxQuantity"`
			StepSize    string `json:"stepSize"`
		} `json:"quantity"`
		Leverage struct {
			MinLeverage string `json:"minLeverage"`
			MaxLeverage string `json:"maxLeverage"`
			StepSize    string `json:"stepSize"`
		} `json:"leverage"`
	} `json:"filters"`
}

type Depth struct {
	Asks         [][]string `json:"asks"`
	Bids         [][]string `json:"bids"`
	LastUpdateId string     `json:"lastUpdateId"`
}

type KLines struct {
	Start  string `json:"start"`
	Open   string `json:"open"`
	High   string `json:"high"`
	Low    string `json:"low"`
	Close  string `json:"close"`
	End    string `json:"end"`
	Volume string `json:"volume"`
	Trades string `json:"trades"`
}
