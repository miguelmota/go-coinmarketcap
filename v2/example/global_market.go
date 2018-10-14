package main

import (
	"fmt"
	"log"

	cmc "github.com/miguelmota/go-coinmarketcap/v2"
)

func main() {
	// Get global market data
	market, err := cmc.GlobalMarket(&cmc.GlobalMarketOptions{
		Convert: "USD",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(market.ActiveCurrencies)
}
