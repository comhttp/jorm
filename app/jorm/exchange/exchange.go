package exchange

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/comhttp/jorm/app/jdb"
)

type Exchange struct {
	Name        string   `json:"name"`
	Slug        string   `json:"slug"`
	Url         string   `json:"url"`
	Logo        string   `json:"logo"`
	Description string   `json:"description"`
	Established string   `json:"established"`
	Country     string   `json:"country"`
	Volume      float64  `json:"volume"`
	Markets     []Market `json:"markets"`
}
type BaseExchange struct {
	Name   string  `json:"name"`
	Slug   string  `json:"slug"`
	Volume float64 `json:"volume"`
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
	Name        string
	Slug        string
	Url         string
	Logo        string
	Description string
	Established string
	Country     string
	Volume      float64
	Markets     map[string]MarketSrc
	Ticker      ExchangeTicker
}
type ExchangeTicker struct {
	Ask, Bid, High24, Last, Low24, Vol string
}

type MarketSrc struct {
	Symbol     string
	Currencies map[string]Currency
}

// ReadAllExchanges reads in all of the data about all coins in the database
func ReadAllExchanges() {
	e := getExchanges()
	baseExchanges := []BaseExchange{}
	exchanges := []Exchange{}
	//ex := make(map[string]map[string]map[string]float64)
	for i := range e {
		baseExchanges = append(baseExchanges, BaseExchange{
			Name:   e[i].Name,
			Slug:   e[i].Slug,
			Volume: e[i].Volume,
		})
		//m := make(map[string]map[string]float64)
		//for marketSymbol, market := range e[i].Markets {
		//	c := make(map[string]float64)
		//	for currencySymbol, currency := range market.Currencies {
		//		c[currencySymbol] = currency.Volume
		//	}
		//	m[marketSymbol] = c
		//}
		exchanges = append(exchanges, Exchange{
			Name:   e[i].Name,
			Slug:   e[i].Slug,
			Volume: e[i].Volume,
		})
	}
	//jdb.JDB.Write(cfg.C.Out+"/info", "exchanges", map[string]interface{}{
	//	"n": len(e),
	//	"e": exchanges,
	//})
	//jdb.JDB.Write(cfg.C.Out+"/info", "exc", map[string]interface{}{
	//	"n": len(e),
	//	"e": baseExchanges,
	//})
}

func getExchanges() []Exchange {
	//data, err := jdb.JDB.ReadAll(cfg.C.Out + "/exchanges")
	//utl.ErrorLog(err)
	//exchanges := make([][]byte, len(data))
	//for i := range data {
	//	exchanges[i] = []byte(data[i])
	//}
	//ex := make([]Exchange, len(exchanges))
	//for i := range exchanges {
	//	if err := json.Unmarshal(exchanges[i], &ex[i]); err != nil {
	//		fmt.Println("Error", err)
	//	}
	//}
	//return ex
	return nil
}

func GetSource(url string) interface{} {
	var marketsRaw interface{}
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer res.Body.Close()
	mapBody, err := ioutil.ReadAll(res.Body)
	if mapBody != nil {
		json.Unmarshal(mapBody, &marketsRaw)
	}
	return marketsRaw
}

func (e *Exchange) WriteExchange(j *jdb.JDB, ex ExchangeSrc) {
	for _, exs := range ex.Markets {
		mSrc := Market{}
		for _, cur := range exs.Currencies {
			mSrc.Symbol = exs.Symbol
			mSrc.Currencies = append(mSrc.Currencies, cur)
		}
		e.Markets = append(e.Markets, mSrc)
	}
	j.Write("exchanges", e.Slug, e)
}

func (es *ExchangeSrc) GetExchange(j *jdb.JDB) {
	fmt.Println("Get " + es.Name + " Exchange Start")
	var e Exchange
	e.Name = es.Name
	e.Slug = es.Slug
	marketsSrc := GetSource(es.Url).(map[string]interface{})
	if len(marketsSrc) > 0 {
		es.Markets = make(map[string]MarketSrc)

		for key, marketSrcRaw := range marketsSrc {
			marketSrc := marketSrcRaw.(map[string]interface{})
			m := strings.Split(key, "_")
			if nq := m[0]; nq != es.Markets[nq].Symbol {
				es.Markets[nq] = MarketSrc{
					Symbol:     nq,
					Currencies: make(map[string]Currency),
				}
			}
			es.SetCurrencyMarket(
				m[0],
				m[1],
				marketSrc[es.Ticker.Ask],
				marketSrc[es.Ticker.Bid],
				marketSrc[es.Ticker.High24],
				marketSrc[es.Ticker.Last],
				marketSrc[es.Ticker.Low24],
				marketSrc[es.Ticker.Vol])
		}
		e.WriteExchange(j, *es)
		fmt.Println("Get " + e.Name + " Exchange Done")
	} else {
		fmt.Println("Get " + e.Name + " Exchange Fail")
	}
}
