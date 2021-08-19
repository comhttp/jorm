package xsrc

import (
	"github.com/comhttp/jorm/mod/exchange"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/rs/zerolog/log"
	"strings"
)

func GetPoloniexExchange(eq *exchange.ExchangeQueries) {
	//e.Markets = make(map[string]exchange.Market)
	//for key, marketSrcRaw := range marketsRaw {
	//	marketSrc := marketSrcRaw.(map[string]interface{})
	//	m := strings.Split(key, "_")
	//	if nq := m[0]; nq != e.Markets[nq].Symbol {
	//		e.Markets[nq] = exchange.Market{
	//			Symbol:     nq,
	//			Currencies: make(map[string]exchange.Currency),
	//		}
	//	}
	//}
	t := exchange.ExchangeTicker{
		Ask:    "lowestAsk",
		Bid:    "highestBid",
		High24: "high24Hr",
		Last:   "last",
		Low24:  "low24Hr",
		Vol:    "baseVolume",
	}
	e := exchange.ExchangeSrc{
		Name:   "Poloniex",
		Slug:   "poloniex",
		APIUrl: "https://poloniex.com/public?command=returnTicker",
		Ticker: t,
	}

	log.Print("Get " + e.Name + " Exchange Start")
	//var e exchange.Exchange
	//e.Name = es.Name
	//e.Slug = es.Slug
	marketsSrc := make(map[string]interface{})
	err := utl.GetSource(e.APIUrl, &marketsSrc)
	utl.ErrorLog(err)
	if len(marketsSrc) > 0 {
		e.Markets = make(map[string]exchange.MarketSrc)
		for key, marketSrcRaw := range marketsSrc {
			marketSrc := marketSrcRaw.(map[string]interface{})
			m := strings.Split(key, "_")
			if nq := m[0]; nq != e.Markets[nq].Symbol {
				e.Markets[nq] = exchange.MarketSrc{
					Symbol:     nq,
					Currencies: make(map[string]exchange.Currency),
				}
			}
			e.SetCurrencyMarket(
				m[0],
				m[1],
				marketSrc[e.Ticker.Ask],
				marketSrc[e.Ticker.Bid],
				marketSrc[e.Ticker.High24],
				marketSrc[e.Ticker.Last],
				marketSrc[e.Ticker.Low24],
				marketSrc[e.Ticker.Vol])
		}
		eq.SetMarkets(e.Slug, e.Markets)
		log.Print("Get " + e.Name + " Exchange Done")
	} else {
		log.Print("Get " + e.Name + " Exchange Fail")
	}
	//fmt.Println(e)

}
