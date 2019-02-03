package coinmarketcap

import (
	"os"
	"testing"
)

var proAPIKey = os.Getenv("CMC_PRO_API_KEY")
var client = NewClient(&Config{
	ProAPIKey: proAPIKey,
})

func TestCryptocurrencyInfo(t *testing.T) {
	info, err := client.Cryptocurrency.Info(&InfoOptions{
		Symbol: "BTC",
	})
	if err != nil {
		t.Error(err)
	}

	if info["BTC"].Name != "Bitcoin" {
		t.FailNow()
	}
}

func TestCryptocurrencyLatestListings(t *testing.T) {
	listings, err := client.Cryptocurrency.LatestListings(&ListingOptions{
		Limit: 1,
	})
	if err != nil {
		t.Error(err)
	}

	if len(listings) == 0 {
		t.FailNow()
	}
	if listings[0].Name != "Bitcoin" {
		t.FailNow()
	}
	if listings[0].Quote["USD"].Price <= 0 {
		t.FailNow()
	}
}

func TestCryptocurrencyLatestQuotes(t *testing.T) {
	quotes, err := client.Cryptocurrency.LatestQuotes(&QuoteOptions{
		Symbol:  "BTC",
		Convert: "CHF",
	})
	if err != nil {
		t.Error(err)
	}

	if len(quotes) == 0 {
		t.FailNow()
	}
	if quotes[0].Name != "Bitcoin" {
		t.FailNow()
	}
	if quotes[0].Quote["CHF"].Price == 0 {
		t.FailNow()
	}
}

func TestSortOptions(t *testing.T) {
	if SortOptions.MarketCap != "market_cap" {
		t.FailNow()
	}
}
