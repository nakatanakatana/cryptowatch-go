package cryptowatch

import (
	"fmt"
	"testing"
)

func TestCryptwatchClient_Markets(t *testing.T) {
	c := new(CryptwatchClient)
	resp, err := c.Markets()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("%v\n%v", resp, c.Allowance)
}

func TestCryptwatchClient_Price(t *testing.T) {
	c := new(CryptwatchClient)
	m := Market{
		Exchange:     "bitflyer",
		CurrencyPair: "btcjpy",
	}
	resp, err := c.Price(m)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("%v\n%v", resp, c.Allowance)
}

func TestCryptwatchClient_Summary(t *testing.T) {
	c := new(CryptwatchClient)
	m := Market{
		Exchange:     "bitflyer",
		CurrencyPair: "btcjpy",
	}
	resp, err := c.Summary(m)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("%v\n%v", resp, c.Allowance)
}

func TestCryptwatchClient_Trades(t *testing.T) {
	c := new(CryptwatchClient)
	m := Market{
		Exchange:     "bitflyer",
		CurrencyPair: "btcjpy",
	}
	params := TradesParams{}
	resp, err := c.Trades(m, params)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("%v\n%v", resp, c.Allowance)
}

func TestCryptwatchClient_OrderBook(t *testing.T) {
	c := new(CryptwatchClient)
	m := Market{
		Exchange:     "bitflyer",
		CurrencyPair: "btcjpy",
	}
	resp, err := c.OrderBook(m)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("%v\n%v", resp, c.Allowance)
}

func TestCryptwatchClient_OHLC(t *testing.T) {
	c := new(CryptwatchClient)
	m := Market{
		Exchange:     "bitflyer",
		CurrencyPair: "btcjpy",
	}
	params := OHLCParams{}
	resp, err := c.OHLC(m, params)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("%v\n%v", resp, c.Allowance)
}

func TestCryptwatchClient_MarketsPrices(t *testing.T) {
	c := new(CryptwatchClient)
	resp, err := c.MarketsPrices()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("%v\n%v", resp, c.Allowance)
}

func TestCryptwatchClient_MarketsSummaries(t *testing.T) {
	c := new(CryptwatchClient)
	resp, err := c.MarketsSummaries()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("%v\n%v", resp, c.Allowance)
}
