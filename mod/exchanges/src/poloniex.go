package xsrc

import (
	"github.com/comhttp/jorm/mod/exchanges"
	"github.com/comhttp/jorm/pkg/jdb"
)

func getPoloniexExchange(j *jdb.JDB) {
	//e.Markets = make(map[string]exchange.Market)
	//for key, marketSrcRaw := range marketsRaw {
	//	marketSrc := marketSrcRaw.(map[string]interface{})
	//	m := strings.Split(key, "_")
	//	if nq := m[0]; nq != e.Markets[nq].Symbol {
	//		e.Markets[nq] = exchange.Market{
	//			Symbol:     nq,
	//			Currencies: make(map[string]exchange.Currency),
	//		}
	//	}
	//}
	t := exchanges.ExchangeTicker{
		Ask:    "lowestAsk",
		Bid:    "highestBid",
		High24: "high24Hr",
		Last:   "last",
		Low24:  "low24Hr",
		Vol:    "baseVolume",
	}
	e := exchanges.ExchangeSrc{
		Name:        "Poloniex",
		Slug:        "poloniex",
		Url:         "https://poloniex.com/public?command=returnTicker",
		Logo:        "",
		Description: "",
		Established: "",
		Country:     "",
		Ticker:      t,
	}
	e.GetExchange(j)
}
