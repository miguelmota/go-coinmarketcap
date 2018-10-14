package coinmarketcap

// Interface interface
type Interface interface {
	//Listings() ([]*types.Listing, error)
	/*
		Tickers(options *TickersOptions) ([]*types.Ticker, error)
		Ticker(options *TickerOptions) (*types.Ticker, error)
		TickerGraph(options *TickerGraphOptions) (*types.TickerGraph, error)
		GlobalMarket(options *GlobalMarketOptions) (*types.GlobalMarket, error)
		GlobalMarketGraph(options *GlobalMarketGraphOptions) (*types.MarketGraph, error)
		GlobalAltcoinMarketGraph(options *GlobalAltcoinMarketGraphOptions) (*types.MarketGraph, error)
		Markets(options *MarketsOptions) ([]*types.Market, error)
		Price(options *PriceOptions) (float64, error)
		CoinID(symbol string) (int, error)
		CoinSlug(symbol string) (string, error)
	*/
}

// listingsMedia listings response media
type listingsMedia struct {
	//Data []*types.Listing `json:"data"`
}

// Status is the status structure
type Status struct {
	Timestamp    string  `json:"timestamp"`
	ErrorCode    int     `json:"error_code"`
	ErrorMessage *string `json:"error_message"`
	Elapsed      int     `json:"elapsed"`
	CreditCount  int     `json:"credit_count"`
}

// Response is the response structure
type Response struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

// Listing is the listing structure
type Listing struct {
	ID                float64           `json:"id"`
	Name              string            `json:"name"`
	Symbol            string            `json:"symbol"`
	Slug              string            `json:"slug"`
	CirculatingSupply float64           `json:"circulating_supply"`
	TotalSupply       float64           `json:"total_supply"`
	MaxSupply         float64           `json:"max_supply"`
	DateAdded         string            `json:"date_added"`
	NumMarketPairs    float64           `json:"num_market_pairs"`
	CMCRank           float64           `json:"cmc_rank"`
	LastUpdated       string            `json:"last_updated"`
	Quote             map[string]*Quote `json:"quote"`
}

// Quote is the quote structure
type Quote struct {
	Price            float64 `json:"price"`
	Volume24H        float64 `json:"volume_24h"`
	PercentChange1H  float64 `json:"percent_change_1h"`
	PercentChange24H float64 `json:"percent_change_24h"`
	PercentChange7D  float64 `json:"percent_change_7d"`
	MarketCap        float64 `json:"market_cap"`
	LastUpdated      string  `json:"last_updated"`
}

// CryptocurrencyListingsLatestsOptions options
type CryptocurrencyListingsLatestsOptions struct {
	Start   int
	Limit   int
	Convert string
	Sort    string
}

// SortOptions sort options
var SortOptions sortOptions

type sortOptions struct {
	Name              string
	Symbol            string
	DateAdded         string
	MarketCap         string
	Price             string
	CirculatingSupply string
	TotalSupply       string
	MaxSupply         string
	NumMarketPairs    string
	Volume24H         string
	PercentChange1H   string
	PercentChange24H  string
	PercentChange7D   string
}

// Client the CoinMarketCap client
type Client struct {
	proAPIKey string
}

// Config the client config structure
type Config struct {
	ProAPIKey string
}

func init() {
	SortOptions = sortOptions{
		Name:              "name",
		Symbol:            "symbol",
		DateAdded:         "date_added",
		MarketCap:         "market_cap",
		Price:             "price",
		CirculatingSupply: "circulating_supply",
		TotalSupply:       "total_supply",
		MaxSupply:         "max_supply",
		NumMarketPairs:    "num_market_pairs",
		Volume24H:         "volume_24h",
		PercentChange1H:   "percent_change_1h",
		PercentChange24H:  "percent_change_24h",
		PercentChange7D:   "percent_change_7d",
	}
}
