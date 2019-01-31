package main

import (
	"fmt"

	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

func main() {
	client := cmc.NewClient(&cmc.Config{
		ProAPIKey: "01585d6d-123-456-789-3146576cbc70",
	})

	listings, err := client.CryptocurrencyListingsLatest(&cmc.CryptocurrencyListingsLatestOptions{
		Limit: 1,
	})
	if err != nil {
		panic(err)
	}

	for _, listing := range listings {
		fmt.Println(listing.Name)               // "Bitcoin"
		fmt.Println(listing.Quote["USD"].Price) // 6316.75736886
	}
}
