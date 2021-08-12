package xsrc

import (
	"fmt"
	"github.com/comhttp/jorm/mod/exchanges"
	"github.com/comhttp/jorm/pkg/jdb"
)

func getDexTradeExchange(j *jdb.JDB) {
	t := exchanges.ExchangeTicker{
		Ask:    "lowestAsk",
		Bid:    "highestBid",
		High24: "high24Hr",
		Last:   "last",
		Low24:  "low24Hr",
		Vol:    "baseVolume",
	}
	e := exchanges.ExchangeSrc{
		Name:        "Dex Trade",
		Slug:        "dex-trade",
		Url:         "https://api.dex-trade.com/v1/public/symbols",
		Logo:        "",
		Description: "",
		Established: "",
		Country:     "",
		Ticker:      t,
	}
	//e.GetExchange()
	//var ex exchange.Exchange

	log.Println("Get " + e.Name + " Exchange Start")
	var ex exchanges.Exchange
	ex.Name = e.Name
	ex.Slug = e.Slug
	marketsSrc := exchanges.GetSource(e.Url).(map[string]interface{})
	e.Markets = make(map[string]exchanges.MarketSrc)

	if marketsSrc != nil {
		if marketsSrc["data"] != nil {
			for _, marketSrc := range marketsSrc["data"].([]interface{}) {
				m := marketSrc.(map[string]interface{})
				if q := m["quote"]; q != nil {
					if nq := q.(string); nq != e.Markets[nq].Symbol {
						e.Markets[nq] = exchanges.MarketSrc{
							Symbol:     nq,
							Currencies: make(map[string]exchanges.Currency),
						}
					}
					if m["pair"] != nil {
						tickerRaw := exchanges.GetSource("https://api.dex-trade.com/v1/public/ticker?pair=" + m["pair"].(string)).(map[string]interface{})
						if tickerRaw["data"] != nil {
							ticker := tickerRaw["data"].(map[string]interface{})
							//log.Println("Get ticker",ticker)
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
			//jdb.JDB.Write(cfg.C.Out+"/exchanges", e.Slug, e)
			ex.WriteExchange(j, e)
			log.Println("Get Dex Trade Exchange Done")
		}
	} else {
		log.Println("Get Gate Exchange Fail")
	}
}
