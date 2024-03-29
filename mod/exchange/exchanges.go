package exchange

import (
	"encoding/json"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/rs/zerolog/log"
	"net/http"
)

func Queries(j *jdb.JDB, col string) *ExchangeQueries {
	return &ExchangeQueries{
		j:   j,
		col: col,
	}
}

func ViewMarket(w http.ResponseWriter, r *http.Request) {
	//rc := mux.Vars(r)["coin"]
	//var coin coin.Coin
	var coinMarkets CoinMarkets
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
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
