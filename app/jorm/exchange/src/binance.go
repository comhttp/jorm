package xsrc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/p9c/jorm/app/jdb"
	"github.com/p9c/jorm/app/jorm/exchange"
	"github.com/p9c/jorm/pkg/utl"
)

func getBinanceExchange() {
	fmt.Println("Get Binance Exchange Start")
	exchangeRaw := make(map[string]interface{})
	slug := "binance"
	var e exchange.Exchange
	e.Name = "Binance"
	e.Slug = slug
	resps, err := http.Get("https://api.binance.com/api/v3/exchangeInfo")
	utl.ErrorLog(err)
	defer resps.Body.Close()
	mapBodyS, err := ioutil.ReadAll(resps.Body)
	json.Unmarshal(mapBodyS, &exchangeRaw)

	var exchangeTickersRaw []map[string]interface{}
	respsTickers, err := http.Get("https://api.binance.com/api/v3/ticker/24hr")
	utl.ErrorLog(err)
	defer respsTickers.Body.Close()
	mapBodyTickers, err := ioutil.ReadAll(respsTickers.Body)
	json.Unmarshal(mapBodyTickers, &exchangeTickersRaw)
	tickers := make(map[string]map[string]interface{})
	for _, exchangeTicker := range exchangeTickersRaw {
		if exchangeTicker != nil {
			tickers[exchangeTicker["symbol"].(string)] = exchangeTicker
		}
	}
	e.Markets = make(map[string]exchange.Market)
	if exchangeRaw != nil {
		if exchangeRaw["symbols"] != nil {
			for _, marketSrcRaw := range exchangeRaw["symbols"].([]interface{}) {
				marketSrc := marketSrcRaw.(map[string]interface{})
				if q := marketSrc["quoteAsset"]; q != nil {
					if nq := q.(string); nq != e.Markets[nq].Symbol {
						e.Markets[nq] = exchange.Market{
							Symbol:     nq,
							Currencies: make(map[string]exchange.Currency),
						}
					}
					if marketSrc["symbol"] != nil {
						m := marketSrc["symbol"].(string)
						e.SetCurrencyMarket(
							q.(string),
							marketSrc["baseAsset"].(string),
							tickers[m]["askPrice"],
							tickers[m]["bidPrice"],
							tickers[m]["highPrice"],
							tickers[m]["lastPrice"],
							tickers[m]["lowPrice"],
							tickers[m]["volume"])
					}
				}
			}
			jdb.JDB.Write("jorm/exchanges", e.Slug, e)
			fmt.Println("Get Binance Exchange Done")
		}
	}
}
