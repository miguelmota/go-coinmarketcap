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

	listings, err := client.Cryptocurrency.Map(&cmc.MapOptions{
		ListingStatus: "active",
		Limit:         1,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(listings[0].Name)     // "Bitcoin"
	fmt.Println(listings[0].IsActive) // 1
}
