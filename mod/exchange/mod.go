package exchange

import (
	"github.com/comhttp/jorm/pkg/jdb"
)

type ExchangeQueries struct {
	j   *jdb.JDB
	col string
}

type Exchange struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Url  string `json:"url"`

	Logo        string   `json:"logo"`
	Description string   `json:"description"`
	Established string   `json:"established"`
	Country     string   `json:"country"`
	Volume      float64  `json:"volume"`
	Markets     []Market `json:"markets"`
}
type ExchangeShort struct {
	Name   string  `json:"name"`
	Slug   string  `json:"slug"`
	Volume float64 `json:"volume"`
}

type Exchanges struct {
	N int      `json:"n"`
	E []string `json:"e"`
}
type Markets struct {
	N int               `json:"n"`
	M map[string]Market `json:"m"`
}

type Currency struct {
	Symbol string  `json:"symbol"`
	Ask    float64 `json:"ask"`
	Bid    float64 `json:"bid"`
	High   float64 `json:"high"`
	Last   float64 `json:"last"`
	Low    float64 `json:"low"`
	Volume float64 `json:"volume"`
}

type Market struct {
	Symbol     string     `json:"symbol"`
	Currencies []Currency `json:"currencies"`
}

type CoinMarket struct {
	Exchange     string   `json:"exchange"`
	ExchangeSlug string   `json:"exslug"`
	Market       string   `json:"market"`
	Ticker       Currency `json:"ticker"`
}
type CoinMarkets map[string]CoinMarket

type ExchangeSrc struct {
	Name    string
	Slug    string
	APIUrl  string
	Volume  float64
	Markets map[string]MarketSrc
	Ticker  ExchangeTicker
}
type ExchangeTicker struct {
	Ask, Bid, High24, Last, Low24, Vol string
}

type MarketSrc struct {
	Symbol     string
	Currencies map[string]Currency
}
