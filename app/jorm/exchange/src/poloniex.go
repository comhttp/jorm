package xsrc

import (
	"encoding/json"
	"fmt"
	"github.com/comhttp/jorm/app/jdb"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/comhttp/jorm/app/jorm/exchange"
)

func getPoloniexExchange() {
	fmt.Println("GetPoloniexExchangeStart")
	marketsRaw := make(map[string]interface{})

	slug := "poloniex"
	var e exchange.Exchange
	e.Name = "Poloniex"
	e.Slug = slug
	respcs, err := http.Get("https://poloniex.com/public?command=returnTicker")
	if err != nil {
	}
	defer respcs.Body.Close()
	mapBody, err := ioutil.ReadAll(respcs.Body)
	json.Unmarshal(mapBody, &marketsRaw)
	e.Markets = make(map[string]exchange.Market)
	for key, marketSrcRaw := range marketsRaw {
		marketSrc := marketSrcRaw.(map[string]interface{})
		m := strings.Split(key, "_")
		if nq := m[0]; nq != e.Markets[nq].Symbol {
			e.Markets[nq] = exchange.Market{
				Symbol:     nq,
				Currencies: make(map[string]exchange.Currency),
			}
		}
		e.SetCurrencyMarket(
			m[0],
			m[1],
			marketSrc["lowestAsk"],
			marketSrc["highestBid"],
			marketSrc["high24Hr"],
			marketSrc["last"],
			marketSrc["low24Hr"],
			marketSrc["baseVolume"])
	}
	jdb.JDB.Write("jorm/exchanges", e.Slug, e)
	fmt.Println("GetPoloniexExchangeDone")
}
