package xsrc

import (
	"encoding/json"
	"fmt"
	"github.com/comhttp/jorm/mod/exchanges"
	"github.com/comhttp/jorm/pkg/jdb"
	"io/ioutil"
	"net/http"

	"github.com/comhttp/jorm/pkg/utl"
)

func getBinanceExchange(j *jdb.JDB) {
	log.Println("Get Binance Exchange Start")
	t := exchanges.ExchangeTicker{
		Ask:    "lowestAsk",
		Bid:    "highestBid",
		High24: "high24Hr",
		Last:   "last",
		Low24:  "low24Hr",
		Vol:    "baseVolume",
	}
	e := exchanges.ExchangeSrc{
		Name:        "Binance",
		Slug:        "binance",
		Url:         "https://api.binance.com/api/v3/exchangeInfo",
		Logo:        "",
		Description: "",
		Established: "",
		Country:     "",
		Ticker:      t,
	}
	var ex exchanges.Exchange
	ex.Name = e.Name
	ex.Slug = e.Slug

	//exchangeRaw := make(map[string]interface{})
	//resps, err := http.Get("https://api.binance.com/api/v3/exchangeInfo")
	//utl.ErrorLog(err)
	//defer resps.Body.Close()
	//mapBodyS, err := ioutil.ReadAll(resps.Body)
	//json.Unmarshal(mapBodyS, &exchangeRaw)

	marketsSrc := exchanges.GetSource(e.Url).(map[string]interface{})

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
	e.Markets = make(map[string]exchanges.MarketSrc)
	if marketsSrc != nil {
		if marketsSrc["symbols"] != nil {
			for _, marketSrcRaw := range marketsSrc["symbols"].([]interface{}) {
				marketSrc := marketSrcRaw.(map[string]interface{})
				if q := marketSrc["quoteAsset"]; q != nil {
					if nq := q.(string); nq != e.Markets[nq].Symbol {
						e.Markets[nq] = exchanges.MarketSrc{
							Symbol:     nq,
							Currencies: make(map[string]exchanges.Currency),
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
			log.Println("Get Binance Exchange Done")
		}
	}
}
