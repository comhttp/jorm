package exchange

import (
	"encoding/json"
	"fmt"
	"github.com/comhttp/jorm/pkg/utl"

	"github.com/comhttp/jorm/app/jdb"
)

type Exchange struct {
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Url         string  `json:"url"`
	Logo        string  `json:"logo"`
	Description string  `json:"description"`
	Established string  `json:"established"`
	Country     string  `json:"country"`
	Volume      float64 `json:"volume"`
	Markets     Markets `json:"markets"`
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
	Symbol     string              `json:"symbol"`
	Currencies map[string]Currency `json:"currencies"`
}
type Markets map[string]Market

type CoinMarket struct {
	Exchange     string   `json:"exchange"`
	ExchangeSlug string   `json:"exslug"`
	Market       string   `json:"market"`
	Ticker       Currency `json:"ticker"`
}
type CoinMarkets map[string]CoinMarket

// ReadAllExchanges reads in all of the data about all coins in the database
func ReadAllExchanges() map[string]map[string]map[string]float64 {
	e := getExchanges()
	ex := make(map[string]map[string]map[string]float64)
	for i := range e {
		m := make(map[string]map[string]float64)
		for marketSymbol, market := range e[i].Markets {
			c := make(map[string]float64)
			for currencySymbol, currency := range market.Currencies {
				c[currencySymbol] = currency.Volume
			}
			m[marketSymbol] = c
		}
		ex[e[i].Slug] = m
	}
	jdb.JDB.Write("jorm/info", "exchanges", map[string]interface{}{
		"n": len(e),
		"e": ex,
	})
	return ex
}

func getExchanges() []Exchange {
	data, err := jdb.JDB.ReadAll("jorm/exchanges")
	utl.ErrorLog(err)
	exchanges := make([][]byte, len(data))
	for i := range data {
		exchanges[i] = []byte(data[i])
	}
	ex := make([]Exchange, len(exchanges))
	for i := range exchanges {
		if err := json.Unmarshal(exchanges[i], &ex[i]); err != nil {
			fmt.Println("Error", err)
		}
	}
	return ex
}
