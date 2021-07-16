package xsrc

import (
	"encoding/json"
	"fmt"
	"github.com/comhttp/jorm/mod/exchanges"
	"github.com/comhttp/jorm/pkg/utl"
	"io/ioutil"
	"net/http"
)

func getBitZExchange() {
	fmt.Println("Get BitZ Exchange Start")
	marketsRaw := make(map[string]interface{})
	slug := "bitz"
	var e exchanges.Exchange
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
	//	fmt.Println("Get BitZ Exchange Done")
	//}
}
