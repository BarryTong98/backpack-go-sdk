package types

type Balance struct {
	Available string `json:"available"`
	Locked    string `json:"locked"`
	Staked    string `json:"staked"`
}

type Deposit struct {
	Id                      int    `json:"id"`
	ToAddress               string `json:"toAddress"`
	FromAddress             string `json:"fromAddress"`
	ConfirmationBlockNumber int    `json:"confirmationBlockNumber"`
	ProviderId              string `json:"providerId"`
	Source                  string `json:"source"`
	Status                  string `json:"status"`
	TransactionHash         string `json:"transactionHash"`
	SubaccountId            int    `json:"subaccountId"`
	Symbol                  string `json:"symbol"`
	Quantity                string `json:"quantity"`
	CreatedAt               string `json:"createdAt"`
}

type DepositAddress struct {
	Address string `json:"address"`
}

type Withdrawal struct {
	Id              int    `json:"id"`
	Blockchain      string `json:"blockchain"`
	ClientId        string `json:"clientId"`
	Identifier      string `json:"identifier"`
	Quantity        string `json:"quantity"`
	Fee             string `json:"fee"`
	Symbol          string `json:"symbol"`
	Status          string `json:"status"`
	SubaccountId    int    `json:"subaccountId"`
	ToAddress       string `json:"toAddress"`
	TransactionHash string `json:"transactionHash"`
	CreatedAt       string `json:"createdAt"`
}
