package xsrc

import (
	"encoding/json"
	"github.com/comhttp/jorm/mod/exchange"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

func getBitTrexExchange() {
	log.Print("Get Bit Trex Exchange Start")
	var exchangeRaw []map[string]interface{}
	slug := "bittrex"
	var e exchange.Exchange
	e.Name = "Bit Trex"
	e.Slug = slug
	resps, err := http.Get("https://api.bittrex.com/v3/markets")
	utl.ErrorLog(err)
	defer resps.Body.Close()
	mapBodyS, err := ioutil.ReadAll(resps.Body)
	json.Unmarshal(mapBodyS, &exchangeRaw)

	var exchangeSummariesRaw []map[string]interface{}
	respsSummaries, err := http.Get("https://api.bittrex.com/v3/markets/summaries")
	utl.ErrorLog(err)
	defer respsSummaries.Body.Close()
	mapBodySummaries, err := ioutil.ReadAll(respsSummaries.Body)
	json.Unmarshal(mapBodySummaries, &exchangeSummariesRaw)
	summaries := make(map[string]map[string]interface{})
	for _, exchangeSummaries := range exchangeSummariesRaw {
		if exchangeSummaries != nil {
			summaries[exchangeSummaries["symbol"].(string)] = exchangeSummaries
		}
	}
	var exchangeTickersRaw []map[string]interface{}
	respsTickers, err := http.Get("https://api.bittrex.com/v3/markets/tickers")
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

	//e.Markets = make(map[string]exchange.Market)
	//if exchangeRaw != nil {
	//	for _, marketSrc := range exchangeRaw {
	//		//m := marketSrc.(map[string]interface{})
	//		if q := marketSrc["quoteCurrencySymbol"]; q != nil {
	//			if nq := q.(string); nq != e.Markets[nq].Symbol {
	//				e.Markets[nq] = exchange.Market{
	//					Symbol:     nq,
	//					Currencies: make(map[string]exchange.Currency),
	//				}
	//			}
	//			if marketSrc["symbol"] != nil {
	//				m := marketSrc["symbol"].(string)
	//				e.SetCurrencyMarket(
	//					q.(string),
	//					marketSrc["baseCurrencySymbol"].(string),
	//					tickers[m]["askRate"],
	//					tickers[m]["bidRate"],
	//					summaries[m]["high"],
	//					tickers[m]["lastTradeRate"],
	//					summaries[m]["low"],
	//					summaries[m]["volume"])
	//			}
	//		}
	//	}
	//	jdb.JDB.Write(cfg.C.Out+"/exchanges", e.Slug, e)
	//	log.Print("Get BitTrex Exchange Done")
	//}
}
