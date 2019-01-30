package main

import (
	"fmt"
	"log"

	cmc ".."
)

func main() {
	client := cmc.NewClient(&cmc.Config{
		ProAPIKey: "01585d6d-123-456-789-3146576cbc70",
	})

	quotesLatest, err := client.CryptocurrencyQuotesLatest(&cmc.CryptocurrencyQuotesLatestOptions{
		Symbol:  "BTC", //add comma and other coin symbol e.g.: "BTC,ETH,XRP"
		Convert: "BRL", //add comma and other coin symbol e.g.: "BRL,USD"
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, quote := range quotesLatest {
		fmt.Println(quote.Name)               // "Bitcoin"
		fmt.Println(quote.Quote["BRL"].Price) // 12880.958109581403 // BTC price converted to BRL
	}
}
