package main

import (
	"fmt"
	"log"

	cmc "github.com/miguelmota/go-coinmarketcap/v1"
)

func main() {
	tickers, err := cmc.Tickers(10, "USD")
	if err != nil {
		log.Fatal(err)
	}

	for _, ticker := range tickers {
		fmt.Println(ticker.Symbol, ticker.PriceUSD)
	}
}
