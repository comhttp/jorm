package xsrc

import (
	"fmt"
	"github.com/comhttp/jorm/app/jdb"
	"github.com/comhttp/jorm/app/jorm/exchange"
)

func getDexTradeExchange(j *jdb.JDB) {
	t := exchange.ExchangeTicker{
		Ask:    "lowestAsk",
		Bid:    "highestBid",
		High24: "high24Hr",
		Last:   "last",
		Low24:  "low24Hr",
		Vol:    "baseVolume",
	}
	e := exchange.ExchangeSrc{
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

	fmt.Println("Get " + e.Name + " Exchange Start")
	var ex exchange.Exchange
	ex.Name = e.Name
	ex.Slug = e.Slug
	marketsSrc := exchange.GetSource(e.Url).(map[string]interface{})
	e.Markets = make(map[string]exchange.MarketSrc)

	if marketsSrc != nil {
		if marketsSrc["data"] != nil {
			for _, marketSrc := range marketsSrc["data"].([]interface{}) {
				m := marketSrc.(map[string]interface{})
				if q := m["quote"]; q != nil {
					if nq := q.(string); nq != e.Markets[nq].Symbol {
						e.Markets[nq] = exchange.MarketSrc{
							Symbol:     nq,
							Currencies: make(map[string]exchange.Currency),
						}
					}
					if m["pair"] != nil {
						tickerRaw := exchange.GetSource("https://api.dex-trade.com/v1/public/ticker?pair=" + m["pair"].(string)).(map[string]interface{})
						if tickerRaw["data"] != nil {
							ticker := tickerRaw["data"].(map[string]interface{})
							//fmt.Println("Get ticker",ticker)
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
			fmt.Println("Get Dex Trade Exchange Done")
		}
	} else {
		fmt.Println("Get Gate Exchange Fail")
	}
}
