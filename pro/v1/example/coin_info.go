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

	info, err := client.Cryptocurrency.Info(&cmc.InfoOptions{
		Symbol: "BTC,ETH",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(info["BTC"].Name) // "Bitcoin"
	fmt.Println(info["ETH"].Name) // "Ethereum"
}
