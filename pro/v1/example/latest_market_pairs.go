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

	info, err := client.Cryptocurrency.LatestMarketPairs(&cmc.MarketPairOptions{
		Symbol: "BTC",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(info.Name)                                  // "Bitcoin"
	fmt.Println(info.MarketPairs[0].Quote["USD"].Price)     // 8127.80075640354
	fmt.Println(info.MarketPairs[0].ExchangeReported.Price) // 8136.96
}
