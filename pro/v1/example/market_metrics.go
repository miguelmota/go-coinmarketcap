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

	metrics, err := client.GlobalMetrics.LatestQuotes(&cmc.QuoteOptions{
		Convert: "USD",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(metrics.BTCDominance)                // 52.3318
	fmt.Println(metrics.Quote["USD"].TotalMarketCap) // 1.76486412385549e+11
}
