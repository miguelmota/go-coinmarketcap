package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// CryptocurrencyListingsLatests gets the all cryptocurrencies (latest)
func (s *Client) CryptocurrencyListingsLatests(options *CryptocurrencyListingsLatestsOptions) ([]*Listing, error) {
	var params []string
	if options == nil {
		options = new(CryptocurrencyListingsLatestsOptions)
	}
	if options.Start != 0 {
		params = append(params, fmt.Sprintf("start=%v", options.Start))
	}
	if options.Limit != 0 {
		params = append(params, fmt.Sprintf("limit=%v", options.Limit))
	}
	if options.Convert != "" {
		params = append(params, fmt.Sprintf("convert=%s", options.Convert))
	}
	if options.Sort != "" {
		params = append(params, fmt.Sprintf("sort=%s", options.Sort))
	}

	url := fmt.Sprintf("%s/cryptocurrency/listings/latest?%s", baseURL, strings.Join(params, "&"))

	body, err := s.makeReq(url)
	resp := new(Response)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	var listings []*Listing
	ifcs, ok := resp.Data.([]interface{})
	if ok {
		for i := range ifcs {
			ifc := ifcs[i]
			listing := new(Listing)
			b, err := json.Marshal(ifc)
			if err != nil {
				log.Fatal(err)
			}
			err = json.Unmarshal(b, listing)
			if err != nil {
				log.Fatal(err)
			}
			listings = append(listings, listing)
		}
	}

	return listings, nil
}
