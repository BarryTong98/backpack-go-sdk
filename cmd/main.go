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

	//_, err = client.ExecuteOrder(0, "Limit", false, "0.0002", "100000", "", "", "Bid", "WEN_USDC", "", "")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//orders, err := client.GetOpenOrders("WEN_USDC")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, order := range orders {
	//	_, err = client.CancelOpenOrders(order.Symbol)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	address, err := client.GetMarkets()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", address)
}
