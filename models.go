package cryptowatch

import (
	"time"
)

type Allowance struct {
	Cost      int64
	Remaining int64
}
type Market struct {
	Exchange     string `json:"exchange"`
	CurrencyPair string `json:"currencyPair"`
}

func (m *Market) toString() string {
	return m.Exchange + ":" + m.CurrencyPair
}

type MarketsPrices map[string]float64

type MarketsSummaries map[string]Summary

type Price struct {
	Last   float64 `json:"last"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Change Change `json:"change"`
}

type Change struct {
	Percentage float64 `json:"percentage"`
	Absolute   float64 `json:"absolute"`
}

type Summary struct {
	Price  Price `json:"price"`
	Volume float64        `json:"volume"`
}

type Trade struct {
	ID        int64 `json:"id"`
	Timestamp int64 `json:"timestamp"`
	Time      time.Time
	Price     float64 `json:"price"`
	Amount    float64`json:"amout"`
}

type Book struct {
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}

type OrderBook struct {
	Asks []Book `json:"asks"`
	Bids []Book `json:"bids"`
}

type OHLC map[string][]OHLCSummary

type OHLCSummary struct {
	CloseTimestamp int64
	CloseTime      time.Time
	OpenPrice      float64
	HighPrice      float64
	LowPrice       float64
	ClosePrice     float64
	Volume         float64
}
