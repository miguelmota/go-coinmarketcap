package main

import (
	"fmt"
	"os"

	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

func main() {
	client := cmc.NewClient(&cmc.Config{
		ProAPIKey: os.Getenv("CMC_PRO_API_KEY"),
	})

	quotes, err := client.Cryptocurrency.LatestQuotes(&cmc.QuoteOptions{
		Symbol:  "BTC",
		Convert: "CHF",
	})
	if err != nil {
		panic(err)
	}

	for _, quote := range quotes {
		fmt.Println(quote.Name) // "Bitcoin"

		// BTC price converted to CHF
		fmt.Println(quote.Quote["CHF"].Price) // 3464.684951197137
	}
}
