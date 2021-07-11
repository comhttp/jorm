package xsrc

import (
	"fmt"
	exchange2 "github.com/comhttp/jorm/exchanges"
	jdb2 "github.com/comhttp/jorm/jdb"
)

func getDexTradeExchange(j *jdb2.JDB) {
	t := exchange2.ExchangeTicker{
		Ask:    "lowestAsk",
		Bid:    "highestBid",
		High24: "high24Hr",
		Last:   "last",
		Low24:  "low24Hr",
		Vol:    "baseVolume",
	}
	e := exchange2.ExchangeSrc{
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
	var ex exchange2.Exchange
	ex.Name = e.Name
	ex.Slug = e.Slug
	marketsSrc := exchange2.GetSource(e.Url).(map[string]interface{})
	e.Markets = make(map[string]exchange2.MarketSrc)

	if marketsSrc != nil {
		if marketsSrc["data"] != nil {
			for _, marketSrc := range marketsSrc["data"].([]interface{}) {
				m := marketSrc.(map[string]interface{})
				if q := m["quote"]; q != nil {
					if nq := q.(string); nq != e.Markets[nq].Symbol {
						e.Markets[nq] = exchange2.MarketSrc{
							Symbol:     nq,
							Currencies: make(map[string]exchange2.Currency),
						}
					}
					if m["pair"] != nil {
						tickerRaw := exchange2.GetSource("https://api.dex-trade.com/v1/public/ticker?pair=" + m["pair"].(string)).(map[string]interface{})
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
