package xsrc

import (
	"encoding/json"
	"fmt"
	exchange2 "github.com/comhttp/jorm/exchanges"
	jdb2 "github.com/comhttp/jorm/jdb"
	"io/ioutil"
	"net/http"

	"github.com/comhttp/jorm/pkg/utl"
)

func getBinanceExchange(j *jdb2.JDB) {
	fmt.Println("Get Binance Exchange Start")
	t := exchange2.ExchangeTicker{
		Ask:    "lowestAsk",
		Bid:    "highestBid",
		High24: "high24Hr",
		Last:   "last",
		Low24:  "low24Hr",
		Vol:    "baseVolume",
	}
	e := exchange2.ExchangeSrc{
		Name:        "Binance",
		Slug:        "binance",
		Url:         "https://api.binance.com/api/v3/exchangeInfo",
		Logo:        "",
		Description: "",
		Established: "",
		Country:     "",
		Ticker:      t,
	}
	var ex exchange2.Exchange
	ex.Name = e.Name
	ex.Slug = e.Slug

	//exchangeRaw := make(map[string]interface{})
	//resps, err := http.Get("https://api.binance.com/api/v3/exchangeInfo")
	//utl.ErrorLog(err)
	//defer resps.Body.Close()
	//mapBodyS, err := ioutil.ReadAll(resps.Body)
	//json.Unmarshal(mapBodyS, &exchangeRaw)

	marketsSrc := exchange2.GetSource(e.Url).(map[string]interface{})

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
	e.Markets = make(map[string]exchange2.MarketSrc)
	if marketsSrc != nil {
		if marketsSrc["symbols"] != nil {
			for _, marketSrcRaw := range marketsSrc["symbols"].([]interface{}) {
				marketSrc := marketSrcRaw.(map[string]interface{})
				if q := marketSrc["quoteAsset"]; q != nil {
					if nq := q.(string); nq != e.Markets[nq].Symbol {
						e.Markets[nq] = exchange2.MarketSrc{
							Symbol:     nq,
							Currencies: make(map[string]exchange2.Currency),
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
			//jdb.JDB.Write(cfg.C.Out+"/exchanges", e.Slug, e)
			ex.WriteExchange(j, e)
			fmt.Println("Get Binance Exchange Done")
		}
	}
}
