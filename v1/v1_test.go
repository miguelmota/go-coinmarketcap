package coinmarketcap

import "testing"

func TestTickers(t *testing.T) {
	tickers, err := Tickers(10, "EUR")
	if err != nil {
		t.FailNow()
	}

	if len(tickers) != 10 {
		t.FailNow()
	}
}
