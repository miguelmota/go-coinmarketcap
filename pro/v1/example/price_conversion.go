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

	listing, err := client.Tools.PriceConversion(&cmc.ConvertOptions{
		Amount:  100,
		Symbol:  "BTC",
		Convert: "USD",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(listing.Quote["USD"].Price)
}
