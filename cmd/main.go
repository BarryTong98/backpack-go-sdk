package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"

	"backpack-trade-bot/service"
)

func main() {
	var cfg service.Config
	if _, err := toml.DecodeFile("config/config.toml", &cfg); err != nil {
		log.Fatal(err)
	}

	// Initialize client
	client, err := service.NewBackpackClient(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Execute open order
	_, err = client.ExecuteOrder(0, "Limit", false, "0.0002", "100000", "", "", "Bid", "WEN_USDC", "", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Order have been successfully executed")

	// Get all open orders
	orders, err := client.GetOpenOrders("WEN_USDC")
	if err != nil {
		log.Fatal(err)
	}
	for _, order := range orders {
		fmt.Printf("ClientId: %s, Id: %s, Symbol: %s \n", order.ClientId, order.Id, order.Symbol)
	}

	//Cancel all open orders
	_, err = client.CancelOpenOrder(orders[0].ClientId, orders[0].Id, orders[0].Symbol)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Order have been successfully canceled")

	// Get Markets coins
	markets, err := client.GetMarkets()
	if err != nil {
		log.Fatal(err)
	}
	for _, market := range markets {
		fmt.Printf("Symbol: %s, BaseSymbol: %s, QuoteSymbol: %s \n", market.Symbol, market.BaseSymbol, market.QuoteSymbol)
	}
}
