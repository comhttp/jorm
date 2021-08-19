package xsrc

import (
	"encoding/json"
	"github.com/comhttp/jorm/mod/exchange"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

func getBitZExchange() {
	log.Print("Get BitZ Exchange Start")
	marketsRaw := make(map[string]interface{})
	slug := "bitz"
	var e exchange.Exchange
	e.Name = "BitZ"
	e.Slug = slug
	respcs, err := http.Get("https://apiv2.bitz.com/Market/tickerall")
	utl.ErrorLog(err)

	defer respcs.Body.Close()
	mapBody, err := ioutil.ReadAll(respcs.Body)
	json.Unmarshal(mapBody, &marketsRaw)
	//e.Markets = make(map[string]exchange.Market)
	//if marketsRaw["data"] != nil {
	//	for symbol, marketSrc := range marketsRaw["data"].(map[string]interface{}) {
	//		mSrc := marketSrc.(map[string]interface{})
	//		if mSrc[symbol] != nil {
	//			m := strings.Split(symbol, "_")
	//			if nq := strings.ToUpper(m[1]); nq != e.Markets[nq].Symbol {
	//				e.Markets[nq] = exchange.Market{
	//					Symbol:     strings.ToUpper(nq),
	//					Currencies: make(map[string]exchange.Currency),
	//				}
	//			}
	//			e.SetCurrencyMarket(
	//				strings.ToUpper(m[1]),
	//				strings.ToUpper(m[0]),
	//				mSrc["askPrice"],
	//				mSrc["bidPrice"],
	//				mSrc["high"],
	//				mSrc["now"],
	//				mSrc["low"],
	//				mSrc["volume"])
	//		}
	//	}
	//	jdb.JDB.Write(cfg.C.Out+"/exchanges", e.Slug, e)
	//	log.Print("Get BitZ Exchange Done")
	//}
}
