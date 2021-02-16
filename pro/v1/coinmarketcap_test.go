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

func TestCryptocurrencyInfoName(t *testing.T) {
	info, err := client.Cryptocurrency.Info(&InfoOptions{
		Slug: "bitcoin",
	})
	if err != nil {
		t.Error(err)
	}
	if len(info) == 0 {
		t.FailNow()
	}
	for k := range info {
		if info[k].Name != "Bitcoin" {
			t.FailNow()
		}
	}
}

func TestCryptocurrencyMap(t *testing.T) {
	listings, err := client.Cryptocurrency.Map(&MapOptions{
		ListingStatus: "active",
		Limit:         1,
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
	if listings[0].IsActive != 1 {
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

func TestCryptocurrencyLatestMarketPairs(t *testing.T) {
	t.Skip("requires paid plan for api")

	info, err := client.Cryptocurrency.LatestMarketPairs(&MarketPairOptions{
		Symbol: "BTC",
	})
	if err != nil {
		t.Error(err)
	}
	if info.Name != "Bitcoin" {
		t.FailNow()
	}
	if len(info.MarketPairs) == 0 {
		t.FailNow()
	}
	if info.MarketPairs[0].Quote["USD"].Price == 0 {
		t.FailNow()
	}
	if info.MarketPairs[0].ExchangeReported.Price == 0 {
		t.FailNow()
	}
}

func TestCryptocurrencyHistoricalOHLCV(t *testing.T) {
	// TODO
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

func TestCryptocurrencyHistoricalQuotes(t *testing.T) {
	// TODO
}

func TestExchangeInfo(t *testing.T) {
	// TODO
}

func TestExchangeMap(t *testing.T) {
	// TODO
}

func TestExchangeLatestListings(t *testing.T) {
	// TODO
}

func TestExchangeLatestMarketPairs(t *testing.T) {
	// TODO
}

func TestExchangeLatestQuotes(t *testing.T) {
	// TODO
}

func TestExchangeHistoricalQuotes(t *testing.T) {
	// TODO
}

func TestGlobalMetricsLatestQuotes(t *testing.T) {
	metrics, err := client.GlobalMetrics.LatestQuotes(&QuoteOptions{
		Convert: "USD",
	})
	if err != nil {
		t.Error(err)
	}
	if metrics.BTCDominance <= 0 {
		t.FailNow()
	}
	if metrics.Quote["USD"].TotalMarketCap <= 0 {
		t.FailNow()
	}
}

func TestGlobalMetricsHistoricalQuotes(t *testing.T) {
	// TODO
}

func TestToolsPriceConversion(t *testing.T) {
	t.Skip("requires paid plan for api")
	listing, err := client.Tools.PriceConversion(&ConvertOptions{
		Amount:  100,
		Symbol:  "BTC",
		Convert: "USD",
	})
	if err != nil {
		t.Error(err)
	}
	if listing.Quote["USD"].Price <= 0 {
		t.FailNow()
	}
}

func TestSortOptions(t *testing.T) {
	if SortOptions.MarketCap != "market_cap" {
		t.FailNow()
	}
}
