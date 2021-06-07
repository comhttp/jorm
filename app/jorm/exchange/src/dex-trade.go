package xsrc

import (
	"encoding/json"
	"fmt"
	"github.com/p9c/jorm/app/jdb"
	"github.com/p9c/jorm/app/jorm/exchange"
	"io/ioutil"
	"net/http"
)

func getDexTradeExchange() {
	fmt.Println("Get Dex Trade Exchange Start")
	exchangeRaw := make(map[string]interface{})
	slug := "dex-trade"
	var e exchange.Exchange
	e.Name = "Dex Trade"
	e.Slug = slug
	resps, err := http.Get("https://api.dex-trade.com/v1/public/symbols")
	if err != nil {
	}
	defer resps.Body.Close()
	mapBodyS, err := ioutil.ReadAll(resps.Body)
	json.Unmarshal(mapBodyS, &exchangeRaw)
	e.Markets = make(map[string]exchange.Market)
	if exchangeRaw != nil {
		if exchangeRaw["data"] != nil {
			for _, marketSrc := range exchangeRaw["data"].([]interface{}) {
				m := marketSrc.(map[string]interface{})
				if q := m["quote"]; q != nil {
					if nq := q.(string); nq != e.Markets[nq].Symbol {
						e.Markets[nq] = exchange.Market{
							Symbol:     nq,
							Currencies: make(map[string]exchange.Currency),
						}
					}
					if m["pair"] != nil {
						tickerRaw := make(map[string]interface{})
						respt, err := http.Get("https://api.dex-trade.com/v1/public/ticker?pair=" + m["pair"].(string))
						if err != nil {
						}
						defer respt.Body.Close()
						mapBodyT, err := ioutil.ReadAll(respt.Body)
						json.Unmarshal(mapBodyT, &tickerRaw)
						if tickerRaw["data"] != nil {
							ticker := tickerRaw["data"].(map[string]interface{})
							//fmt.Println("ticker", ticker)
							go e.SetCurrencyMarket(
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
			jdb.JDB.Write("jorm/exchanges", e.Slug, e)
			fmt.Println("Get Dex Trade Exchange Done")
		}
	}
}
