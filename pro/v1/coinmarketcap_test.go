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
	info, err := client.CryptocurrencyInfo(&CryptocurrencyInfoOptions{
		Symbol: "BTC",
	})
	if err != nil {
		t.Error(err)
	}

	if info["BTC"].Name != "Bitcoin" {
		t.FailNow()
	}
}

func TestCryptocurrencyListingsLatests(t *testing.T) {
	listings, err := client.CryptocurrencyListingsLatests(&CryptocurrencyListingsLatestsOptions{
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

func TestSortOptions(t *testing.T) {
	if SortOptions.MarketCap != "market_cap" {
		t.FailNow()
	}
}
