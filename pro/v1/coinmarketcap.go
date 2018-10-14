// Package coinmarketcap Coin Market Cap API client for Go
package coinmarketcap

var (
	siteURL               = "https://coinmarketcap.com"
	baseURL               = "https://pro-api.coinmarketcap.com/v1"
	coinGraphURL          = "https://graphs2.coinmarketcap.com/currencies"
	globalMarketGraphURL  = "https://graphs2.coinmarketcap.com/global/marketcap-total"
	altcoinMarketGraphURL = "https://graphs2.coinmarketcap.com/global/marketcap-altcoin"
)

// NewClient initializes a new client
func NewClient(cfg *Config) *Client {
	return &Client{
		proAPIKey: cfg.ProAPIKey,
	}
}
