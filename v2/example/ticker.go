package main

import (
	"fmt"
	"log"

	cmc "github.com/miguelmota/go-coinmarketcap/v2"
)

func main() {
	// Get info about coin
	ticker, err := cmc.Ticker(&cmc.TickerOptions{
		Symbol:  "ETH",
		Convert: "EUR",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ticker.Name, ticker.Quotes["EUR"].Price)
}
