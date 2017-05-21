package cryptowatch

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/mitchellh/mapstructure"
)

const (
	URIBase = "https://api.cryptowat.ch/"
)

type CryptwatchClient struct {
	Allowance Allowance
	http.Client
}

// Get: base function to access cryptwatch
func (c *CryptwatchClient) Get(resourcePath string, params url.Values) (interface{}, error) {
	URI := URIBase + resourcePath + params.Encode()
	resp, err := c.Client.Get(URI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	err = mapstructure.Decode(result["allowance"], &c.Allowance)
	if err != nil {
		return nil, err
	}
	return result["result"], nil
}

// Markets: Returns a list of all supported markets.
// Example: https://api.cryptowat.ch/markets
func (c *CryptwatchClient) Markets() (*[]Market, error) {
	result, err := c.Get("markets", nil)
	if err != nil {
		return nil, err
	}
	markets := &[]Market{}
	err = mapstructure.Decode(result, markets)
	if err != nil {
		return nil, err
	}
	return markets, nil
}

// Price: Returns a market’s current price.
// Example: https://api.cryptowat.ch/markets/coinbase/btcusd/price
func (c *CryptwatchClient) Price(m Market) (float64, error) {
	URI := "markets/" + m.Exchange + "/" + m.CurrencyPair + "/price"
	result, err := c.Get(URI, nil)
	if err != nil {
		return -1, err
	}
	price := result.(map[string]interface{})["price"].(float64)
	return price, nil
}

// Summary: Returns a market’s current price as well as other stats based on a 24-hour sliding window.
// Example: https://api.cryptowat.ch/markets/coinbase/btcusd/summary
func (c *CryptwatchClient) Summary(m Market) (*Summary, error) {
	URI := "markets/" + m.Exchange + "/" + m.CurrencyPair + "/summary"
	result, err := c.Get(URI, nil)
	if err != nil {
		return nil, err
	}
	summary := &Summary{}
	err = mapstructure.Decode(result, summary)
	if err != nil {
		return nil, err
	}
	return summary, nil

}

// Trades: Returns a market’s most recent trades, incrementing chronologically.
// Example: https://api.cryptowat.ch/markets/coinbase/btcusd/trades
func (c *CryptwatchClient) Trades(m Market, params TradesParams) (*[]Trade, error) {
	URI := "markets/" + m.Exchange + "/" + m.CurrencyPair + "/trades"
	result, err := c.Get(URI, params.toValues())
	if err != nil {
		return nil, err
	}
	results := result.([]interface{})
	trades := make([]Trade, len(results))
	var tmp []interface{}
	for i, r := range results {
		tmp = r.([]interface{})
		trades[i] = Trade{
			ID:        int64(tmp[0].(float64)),
			Timestamp: int64(tmp[1].(float64)),
			Time:      time.Unix(int64(tmp[1].(float64)), 0),
			Price:     tmp[2].(float64),
			Amount:    tmp[3].(float64),
		}
	}
	return &trades, nil
}

// Order Book: Returns a market’s order book.
// Example: https://api.cryptowat.ch/markets/coinbase/btcusd/orderbook
func (c *CryptwatchClient) OrderBook(m Market) (*OrderBook, error) {
	URI := "markets/" + m.Exchange + "/" + m.CurrencyPair + "/orderbook"
	result, err := c.Get(URI, nil)
	if err != nil {
		return nil, err
	}
	results := result.(map[string]interface{})
	asks, err := parseBook(results["asks"])
	if err != nil {
		return nil, err
	}
	bids, err := parseBook(results["bids"])
	if err != nil {
		return nil, err
	}
	orderBook := &OrderBook{
		Asks: asks,
		Bids: bids,
	}
	return orderBook, nil
}

func parseBook(input interface{}) ([]Book, error) {
	var tmp []interface{}
	inputs := input.([]interface{})
	books := make([]Book, len(inputs))
	for i, r := range inputs {
		tmp = r.([]interface{})
		books[i] = Book{
			Price:  tmp[0].(float64),
			Amount: tmp[1].(float64),
		}
	}

	return books, nil
}

// OHLC: Returns a market’s OHLC candlestick data. Returns data as lists of lists of numbers for each time period integer.
// Example: https://api.cryptowat.ch/markets/coinbase/btcusd/ohlc
func (c *CryptwatchClient) OHLC(m Market, params OHLCParams) (*OHLC, error) {
	URI := "markets/" + m.Exchange + "/" + m.CurrencyPair + "/ohlc"
	result, err := c.Get(URI, params.toValues())
	if err != nil {
		return nil, err
	}
	results := result.(map[string]interface{})
	OHLC := OHLC{}
	for key, value := range results {
		values := value.([]interface{})
		summaries := make([]OHLCSummary, len(values))
		for i, v := range values {
			vs := v.([]interface{})
			summaries[i] = OHLCSummary{
				CloseTimestamp: int64(vs[0].(float64)),
				CloseTime:      time.Unix(int64(vs[0].(float64)), 0),
				OpenPrice:      vs[1].(float64),
				HighPrice:      vs[2].(float64),
				LowPrice:       vs[3].(float64),
				ClosePrice:     vs[4].(float64),
				Volume:         vs[5].(float64),
			}
		}
		OHLC[key] = summaries
	}
	return &OHLC, nil
}

// Prices: Returns the current price for all supported markets. Some values may be out of date by a few seconds.
// Example: https://api.cryptowat.ch/markets/prices
func (c *CryptwatchClient) MarketsPrices() (*MarketsPrices, error) {
	result, err := c.Get("markets/prices", nil)
	if err != nil {
		return nil, err
	}
	marketsPrices := &MarketsPrices{}
	err = mapstructure.Decode(result, marketsPrices)
	if err != nil {
		return nil, err
	}
	return marketsPrices, nil
}

// Summaries: Returns the market summary for all supported markets. Some values may be out of date by a few seconds.
// Example: https://api.cryptowat.ch/markets/summaries
func (c *CryptwatchClient) MarketsSummaries() (*MarketsSummaries, error) {
	result, err := c.Get("markets/summaries", nil)
	if err != nil {
		return nil, err
	}
	marketsSummaries := &MarketsSummaries{}
	err = mapstructure.Decode(result, marketsSummaries)
	if err != nil {
		return nil, err
	}
	return marketsSummaries, nil
}
