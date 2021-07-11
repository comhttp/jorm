package xsrc

import (
	exchange2 "github.com/comhttp/jorm/exchanges"
	jdb2 "github.com/comhttp/jorm/jdb"
)

func getPoloniexExchange(j *jdb2.JDB) {
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
	t := exchange2.ExchangeTicker{
		Ask:    "lowestAsk",
		Bid:    "highestBid",
		High24: "high24Hr",
		Last:   "last",
		Low24:  "low24Hr",
		Vol:    "baseVolume",
	}
	e := exchange2.ExchangeSrc{
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
