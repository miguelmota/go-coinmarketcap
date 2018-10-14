package coinmarketcap

import (
	"testing"
)

var proAPIKey = "02585d6d-fa6a-460c-833c-3146576cbc70"
var client = NewClient(&Config{
	ProAPIKey: proAPIKey,
})

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

/*
func TestTickers(t *testing.T) {
	tickers, err := Tickers(&TickersOptions{
		Start:   0,
		Limit:   10,
		Convert: "USD",
	})
	if err != nil {
		t.Error(err)
	}

	if len(tickers) != 10 {
		t.FailNow()
	}
}

func TestTicker(t *testing.T) {
	ticker, err := Ticker(&TickerOptions{
		Symbol:  "ETH",
		Convert: "USD",
	})
	if err != nil {
		t.FailNow()
	}
	if ticker.ID <= 0 {
		t.FailNow()
	}
	if ticker.Name != "Ethereum" {
		t.FailNow()
	}
	if ticker.Symbol != "ETH" {
		t.FailNow()
	}
	if ticker.Slug != "ethereum" {
		t.FailNow()
	}
	if ticker.Rank <= 0 {
		t.FailNow()
	}
	if ticker.CirculatingSupply <= 0 {
		t.FailNow()
	}
	if ticker.TotalSupply <= 0 {
		t.FailNow()
	}
	if ticker.MaxSupply <= -1 {
		t.FailNow()
	}
	if ticker.LastUpdated <= 0 {
		t.FailNow()
	}
	if ticker.Quotes["USD"].MarketCap <= 0 {
		t.FailNow()
	}
	if ticker.Quotes["USD"].Volume24H <= 0 {
		t.FailNow()
	}
	if ticker.Quotes["USD"].MarketCap <= 0 {
		t.FailNow()
	}
	if ticker.Quotes["USD"].PercentChange1H == 0 {
		t.FailNow()
	}
	if ticker.Quotes["USD"].PercentChange24H == 0 {
		t.FailNow()
	}
	if ticker.Quotes["USD"].PercentChange7D == 0 {
		t.FailNow()
	}
}

func TestTickerGraph(t *testing.T) {
	var threeMonths int64 = (60 * 60 * 24 * 90)
	end := time.Now().Unix()
	start := end - threeMonths

	graph, err := TickerGraph(&TickerGraphOptions{
		Symbol: "ETH",
		Start:  start,
		End:    end,
	})
	if err != nil {
		t.FailNow()
	}

	if graph.MarketCapByAvailableSupply[0][0] == 0 {
		t.FailNow()
	}
	if graph.PriceBTC[0][0] == 0 {
		t.FailNow()
	}
	if graph.PriceUSD[0][0] == 0 {
		t.FailNow()
	}
	if graph.VolumeUSD[0][0] == 0 {
		t.FailNow()
	}
}

func TestGlobalMarket(t *testing.T) {
	market, err := GlobalMarket(&GlobalMarketOptions{
		Convert: "USD",
	})
	if err != nil {
		t.FailNow()
	}
	if market.ActiveCurrencies <= 0 {
		t.FailNow()
	}
	if market.ActiveMarkets <= 0 {
		t.FailNow()
	}
	if market.LastUpdated <= 0 {
		t.FailNow()
	}
	if market.BitcoinPercentageOfMarketCap == 0 {
		t.FailNow()
	}
	if market.Quotes["USD"].TotalVolume24H == 0 {
		t.FailNow()
	}
	if market.Quotes["USD"].TotalMarketCap == 0 {
		t.FailNow()
	}
}

func TestGlobalMarketGraph(t *testing.T) {
	var threeMonths int64 = (60 * 60 * 24 * 90)
	end := time.Now().Unix()
	start := end - threeMonths
	graph, err := GlobalMarketGraph(&GlobalMarketGraphOptions{
		Start: start,
		End:   end,
	})
	if err != nil {
		t.FailNow()
	}
	if graph.MarketCapByAvailableSupply[0][0] == 0 {
		t.FailNow()
	}
	if graph.VolumeUSD[0][0] == 0 {
		t.FailNow()
	}
}

func TestGlobalAltcoinMarketGraph(t *testing.T) {
	var threeMonths int64 = (60 * 60 * 24 * 90)
	end := time.Now().Unix()
	start := end - threeMonths
	graph, err := GlobalAltcoinMarketGraph(&GlobalAltcoinMarketGraphOptions{
		Start: start,
		End:   end,
	})
	if err != nil {
		t.FailNow()
	}
	if graph.MarketCapByAvailableSupply[0][0] == 0 {
		t.FailNow()
	}
	if graph.VolumeUSD[0][0] == 0 {
		t.FailNow()
	}
}

func TestMarkets(t *testing.T) {
	markets, err := Markets(&MarketsOptions{
		Symbol: "ETH",
	})
	if err != nil {
		t.FailNow()
	}
	if len(markets) == 0 {
		t.FailNow()
	}

	market := markets[0]
	if market.Rank == 0 {
		t.FailNow()
	}
	if market.Exchange == "" {
		t.FailNow()
	}
	if market.Pair == "" {
		t.FailNow()
	}
	if market.VolumeUSD == 0 {
		t.FailNow()
	}
	if market.Price == 0 {
		t.FailNow()
	}
	if market.VolumePercent == 0 {
		t.FailNow()
	}
	if market.Updated == "" {
	}
}

func TestPrice(t *testing.T) {
	price, err := Price(&PriceOptions{
		Symbol:  "ETH",
		Convert: "USD",
	})
	if err != nil {
		t.FailNow()
	}
	if price <= 0 {
		t.FailNow()
	}
}

func TestDoReq(t *testing.T) {
	// TODO
}

func TestMakeReq(t *testing.T) {
	// TODO
}

func TestToInt(t *testing.T) {
	v := toInt("5")
	if v != 5 {
		t.FailNow()
	}
}

func TestToFloat(t *testing.T) {
	v := toFloat("5.2")
	if v != 5.2 {
		t.FailNow()
	}
}
*/
