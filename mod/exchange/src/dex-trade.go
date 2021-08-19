package xsrc

import (
	"fmt"
	"github.com/comhttp/jorm/mod/exchange"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/rs/zerolog/log"
)

func GetDexTradeExchange(eq *exchange.ExchangeQueries) {
	t := exchange.ExchangeTicker{
		Ask:    "lowestAsk",
		Bid:    "highestBid",
		High24: "high24Hr",
		Last:   "last",
		Low24:  "low24Hr",
		Vol:    "baseVolume",
	}
	e := exchange.ExchangeSrc{
		Name:   "Dex Trade",
		Slug:   "dex-trade",
		APIUrl: "https://api.dex-trade.com/v1/public/",
		Ticker: t,
	}
	//e.GetExchange()
	//var ex exchange.Exchange

	fmt.Println("Get " + e.Name + " Exchange Start")
	var ex exchange.Exchange
	ex.Name = e.Name
	ex.Slug = e.Slug
	marketsSrc := make(map[string]interface{})
	err := utl.GetSource(e.APIUrl+"symbols", &marketsSrc)
	utl.ErrorLog(err)
	e.Markets = make(map[string]exchange.MarketSrc)
	if marketsSrc != nil {
		if marketsSrc["data"] != nil {
			for _, marketSrc := range marketsSrc["data"].([]interface{}) {
				//m := marketSrc.(map[string]interface{})
				//if q := m["quote"]; q != nil {
				//if nq := q.(string); nq != e.Markets[nq].Symbol {
				//	e.Markets[nq] = exchange.MarketSrc{
				//		Symbol:     nq,
				//		Currencies: make(map[string]exchange.Currency),
				//	}
				//}
				//if m["pair"] != nil {
				dexTradePair(&e, marketSrc)
				//tickerRaw := make(map[string]interface{})
				//err := utl.GetSource(e.APIUrl+"ticker?pair=" + m["pair"].(string),&tickerRaw)
				//utl.ErrorLog(err)
				//fmt.Println("tickerRawtickerRawtickerRaw No:" ,tickerRaw)
				//if tickerRaw["data"] != nil {
				//	ticker := tickerRaw["data"].(map[string]interface{})
				//	//log.Print("Get ticker",ticker)
				//	e.SetCurrencyMarket(
				//		q.(string),
				//		m["base"].(string),
				//		ticker["open"],
				//		ticker["close"],
				//		ticker["high"],
				//		ticker["last"],
				//		ticker["low"],
				//		ticker["volume_24H"])
				//}
				//}
				//}
			}
		}
		//jdb.JDB.Write(cfg.C.Out+"/exchanges", e.Slug, e)
		//ex.WriteExchange(j, e)
		eq.SetMarkets(e.Slug, e.Markets)
		log.Print("Get Dex Trade Exchange Done")
	} else {
		log.Print("Get Gate Exchange Fail")
	}
}
func dexTradePair(e *exchange.ExchangeSrc, marketSrc interface{}) {
	m := marketSrc.(map[string]interface{})
	if q := m["quote"]; q != nil {
		if nq := q; nq != e.Markets[nq.(string)].Symbol {
			e.Markets[nq.(string)] = exchange.MarketSrc{
				Symbol:     nq.(string),
				Currencies: make(map[string]exchange.Currency),
			}
		}
		if m["pair"] != nil {
			tickerRaw := make(map[string]interface{})
			err := utl.GetSource(e.APIUrl+"ticker?pair="+m["pair"].(string), &tickerRaw)
			utl.ErrorLog(err)
			if tickerRaw["data"] != nil {
				ticker := tickerRaw["data"].(map[string]interface{})
				e.SetCurrencyMarket(
					q.(string),
					m["base"].(string),
					ticker["open"],
					ticker["close"],
					ticker["high"],
					ticker["last"],
					ticker["low"],
					ticker["volume_24H"])
			}
		}
	}
}
