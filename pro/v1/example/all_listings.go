package main

import (
	"fmt"
	"log"

	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

func main() {
	client := cmc.NewClient(&Config{
		ProAPIKey: "01585d6d-123-456-789-3146576cbc70",
	})

	listings, err := client.CryptocurrencyListingsLatests(&cmc.CryptocurrencyListingsLatestsOptions{
		Limit: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(listings[0].Name)
	fmt.Println(listings[0].Quote["USD"].Price)
}
