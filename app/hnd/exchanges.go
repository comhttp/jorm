package hnd

import (
	"encoding/json"
	"fmt"
	exchange2 "github.com/comhttp/jorm/exchanges"
	"net/http"
)

func ViewMarket(w http.ResponseWriter, r *http.Request) {
	//rc := mux.Vars(r)["coin"]
	//var coin coin.Coin
	var coinMarkets exchange2.CoinMarkets
	//jdb.JDB.Read(cfg.Path+"/www/coins", rc, &coin)
	//exchanges := exchange.ReadAllExchanges()
	//for _, e := range exchanges {
	//	for _, market := range e.Markets {
	//		for _, cur := range market.Currencies {
	//			if cur.Symbol == coin.Ticker {
	//				coinMarket := exchange.CoinMarket{
	//					Exchange:     e.Name,
	//					ExchangeSlug: e.Slug,
	//					Market:       market.Symbol,
	//					Ticker:       cur,
	//				}
	//				coinMarkets[cur.Symbol] = coinMarket
	//			}
	//		}
	//	}
	//}
	x := map[string]interface{}{
		"d": coinMarkets,
	}
	out, err := json.Marshal(x)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
