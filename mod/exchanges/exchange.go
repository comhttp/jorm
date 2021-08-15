package exchanges

import (
	"encoding/json"
	"fmt"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strings"
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

func (e *Exchange) WriteExchange(j *jdb.JDB, ex ExchangeSrc) {
	for _, exs := range ex.Markets {
		mSrc := Market{}
		for _, cur := range exs.Currencies {
			mSrc.Symbol = exs.Symbol
			mSrc.Currencies = append(mSrc.Currencies, cur)
		}
		e.Markets = append(e.Markets, mSrc)
	}
	j.Write("exchanges", "ex_"+e.Slug, e)
}

func (es *ExchangeSrc) SetExchange(j *jdb.JDB) {
	log.Print("Get " + es.Name + " Exchange Start")
	var e Exchange
	e.Name = es.Name
	e.Slug = es.Slug
	marketsSrc := utl.GetSource(es.Url).(map[string]interface{})
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
		log.Print("Get " + e.Name + " Exchange Done")
	} else {
		log.Print("Get " + e.Name + " Exchange Fail")
	}
}
