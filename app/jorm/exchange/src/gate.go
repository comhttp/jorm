package xsrc

import (
	"encoding/json"
	"fmt"
	"github.com/comhttp/jorm/app/cfg"
	"github.com/comhttp/jorm/app/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/comhttp/jorm/app/jorm/exchange"
)

func getGateExchange() {
	fmt.Println("Get Gate Exchange Start")
	marketsRaw := make(map[string]interface{})
	slug := "gate"
	var e exchange.Exchange
	e.Name = "Gate"
	e.Slug = slug
	respcs, err := http.Get("https://data.gateapi.io/api2/1/tickers")
	utl.ErrorLog(err)
	defer respcs.Body.Close()
	mapBody, err := ioutil.ReadAll(respcs.Body)
	if mapBody != nil {
		json.Unmarshal(mapBody, &marketsRaw)
		e.Markets = make(map[string]exchange.Market)
		for symbol, marketSrc := range marketsRaw {
			m := strings.Split(symbol, "_")
			mSrc := marketSrc.(map[string]interface{})
			if nq := strings.ToUpper(m[1]); nq != e.Markets[nq].Symbol {
				e.Markets[nq] = exchange.Market{
					Symbol:     nq,
					Currencies: make(map[string]exchange.Currency),
				}
			}
			e.SetCurrencyMarket(
				strings.ToUpper(m[1]),
				strings.ToUpper(m[0]),
				mSrc["sell"],
				mSrc["buy"],
				mSrc["high"],
				mSrc["last"],
				mSrc["low"],
				mSrc["vol"])
		}
		jdb.JDB.Write(cfg.C.Out+"/exchanges", e.Slug, e)
		fmt.Println("Get Gate Exchange Done")
	} else {
		fmt.Println("Get Gate Exchange Fail")
	}
}
