package main

import (
	"fmt"

	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

func main() {
	client := cmc.NewClient(&cmc.Config{
		ProAPIKey: "01585d6d-123-456-789-3146576cbc70",
	})

	quotesLatest, err := client.CryptocurrencyQuotesLatest(&cmc.CryptocurrencyQuotesLatestOptions{
		Symbol:  "BTC",
		Convert: "BRL",
	})
	if err != nil {
		panic(err)
	}

	for _, quote := range quotesLatest {
		fmt.Println(quote.Name) // "Bitcoin"

		// BTC price converted to BRL
		fmt.Println(quote.Quote["BRL"].Price) // 12880.958109581403
	}
}
