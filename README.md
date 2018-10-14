<h1 align="center">
  <br />
  <!--
  <img src="https://user-images.githubusercontent.com/168240/39501128-e66e2a18-4d6d-11e8-9e16-88655102da6c.png" alt="go-coinmarketcap" width="600" />
  -->
  <img src="https://user-images.githubusercontent.com/168240/41822669-34e92094-77a8-11e8-831e-67d38e686c21.png" alt="go-coinmarketcap" width="600" />
  <br />
  <br />
  <br />
</h1>

> The Unofficial [CoinMarketCap](https://coinmarketcap.com/) API client for [Go](https://golang.org/).

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/miguelmota/go-coinmarketcap/master/LICENSE.md) [![Build Status](https://travis-ci.org/miguelmota/go-coinmarketcap.svg?branch=master)](https://travis-ci.org/miguelmota/go-coinmarketcap) [![Go Report Card](https://goreportcard.com/badge/github.com/miguelmota/go-coinmarketcap?)](https://goreportcard.com/report/github.com/miguelmota/go-coinmarketcap) [![GoDoc](https://godoc.org/github.com/miguelmota/go-coinmarketcap?status.svg)](https://godoc.org/github.com/miguelmota/go-coinmarketcap)

Supports the CoinMarketCap API Pro Version, [V2](https://coinmarketcap.com/api) and V1 Public API

## Documentation

[https://godoc.org/github.com/miguelmota/go-coinmarketcap](https://godoc.org/github.com/miguelmota/go-coinmarketcap)

## Install

```bash
go get -u github.com/miguelmota/go-coinmarketcap
```

## Pro V1 (latest)

| Type           | Endpoint                               | Implemented? |
|----------------|----------------------------------------|--------------|
| Cryptocurrency | /v1/cryptocurrency/info                | Yes          |
| Cryptocurrency | /v1/cryptocurrency/map                 | Not yet      |
| Cryptocurrency | /v1/cryptocurrency/listings/latest     | Yes          |
| Cryptocurrency | /v1/cryptocurrency/market-pairs/latest | Not yet      |
| Cryptocurrency | /v1/cryptocurrency/ohlcv/historical    | Not yet      |
| Cryptocurrency | /v1/cryptocurrency/quotes/latest       | Not yet      |
| Cryptocurrency | /v1/cryptocurrency/quotes/historical   | Not yet      |
| Exchange       | /v1/exchange/info                      | Not yet      |
| Exchange       | /v1/exchange/map                       | Not yet      |
| Exchange       | /v1/exchange/listings/latest           | Not yet      |
| Exchange       | /v1/exchange/market-pairs/latest       | Not yet      |
| Exchange       | /v1/exchange/quotes/latest             | Not yet      |
| Exchange       | /v1/exchange/quotes/historical         | Not yet      |
| Global Metrics | /v1/global-metrics/quotes/latest       | Not yet      |
| Global Metrics | /v1/global-metrics/quotes/historical   | Not yet      |
| Tools          | /v1/tools/price-conversion             | Not yet      |

### Getting started

```go
package main

import (
	"fmt"
	"log"

	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

func main() {
	client := cmc.NewClient(&Config{
		ProAPIKey: "01585d6d-123-456-789-3146576cbc70",
	})

	listings, err := client.CryptocurrencyListingsLatests(&cmc.CryptocurrencyListingsLatestsOptions{
		Limit: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(listings[0].Name)
	fmt.Println(listings[0].Quote["USD"].Price)
}
```

### Examples

For more examples, check out the [`./pro/v1/example`](./pro/v1/example) directory and [documentation](https://godoc.org/github.com/miguelmota/go-coinmarketcap/pro/v1)

---

## V2

Note: will be deprecated December 2018

### Getting started

```go
package main

import (
	"fmt"
	"log"

	cmc "github.com/miguelmota/go-coinmarketcap"
)

func main() {
	tickers, err := cmc.Tickers(&cmc.TickersOptions{
		Start:   0,
		Limit:   100,
		Convert: "USD",
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, ticker := range tickers {
		fmt.Println(ticker.Symbol, ticker.Quotes["USD"].Price)
	}
}
```

### Examples

For more examples, check out the [`./v2/example`](./v2/example) directory and [documentation](https://godoc.org/github.com/miguelmota/go-coinmarketcap/v2)

---

## V1

Note: will be deprecated November 2018

### Getting started

```go
package main

import (
	"fmt"
	"log"

	cmc "github.com/miguelmota/go-coinmarketcap"
)

func main() {
	tickers, err := cmc.Tickers(&cmc.TickersOptions{
		Start:   0,
		Limit:   100,
		Convert: "USD",
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, ticker := range tickers {
		fmt.Println(ticker.Symbol, ticker.Quotes["USD"].Price)
	}
}
```

### Examples

For more examples, check out the [`./v1/example`](./v1/example) directory and [documentation](https://godoc.org/github.com/miguelmota/go-coinmarketcap/v1)

---

## License

MIT
