package exchange

import (
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
)

func (eq *ExchangeQueries) GetExchange(slug string) (Exchange, error) {
	e, err := eq.getExchange(slug)
	return *e, err
}

func (eq *ExchangeQueries) getExchange(key string) (e *Exchange, err error) {
	err = eq.j.Read("exchange", key, &e)
	//utl.ErrorLog(err)
	return e, err
}

func (cq *ExchangeQueries) GetExchanges() Exchanges {
	exchanges, err := cq.j.ReadAll("exchange")
	utl.ErrorLog(err)
	allExchanges := Exchanges{N: 0}
	for i, e := range exchanges {
		allExchanges.E = append(allExchanges.E, e)
		allExchanges.N = i
	}
	return allExchanges
}

func (eq *ExchangeQueries) GetExchangeMarkets(exchange string) (map[string]Market, error) {
	e, err := eq.getMarkets(exchange)
	return e, err
}

func (eq *ExchangeQueries) getMarkets(exchange string) (m map[string]Market, err error) {
	err = eq.j.Read("market", exchange, &m)
	//utl.ErrorLog(err)
	return m, err
}

func (eq *ExchangeQueries) GetExchangeMarket(exchange, market string) (Market, error) {
	m, err := eq.getMarkets(exchange)
	return m[market], err
}

// ReadAllExchanges reads in all of the data about all exchanges in the database
func GetAllExchanges(j *jdb.JDB) {
	////e := getExchanges(j)
	//baseExchanges := []BaseExchange{}
	//exchanges := []Exchange{}
	////ex := make(map[string]map[string]map[string]float64)
	//for i := range e {
	//	baseExchanges = append(baseExchanges, BaseExchange{
	//		Name:   e[i].Name,
	//		Slug:   e[i].Slug,
	//		Volume: e[i].Volume,
	//	})
	//	m := make(map[string]map[string]float64)
	//	for _, market := range e[i].Markets {
	//		c := make(map[string]float64)
	//		for _, currency := range market.Currencies {
	//			c[currency.Symbol] = currency.Volume
	//		}
	//		m[market.Symbol] = c
	//	}
	//	exchanges = append(exchanges, Exchange{
	//		Name:   e[i].Name,
	//		Slug:   e[i].Slug,
	//		Volume: e[i].Volume,
	//	})
	//}
	//j.Write("exchanges", "exchanges", map[string]interface{}{
	//	"n": len(e),
	//	"e": exchanges,
	//})
	//j.Write("exchanges", "base", map[string]interface{}{
	//	"n": len(e),
	//	"e": baseExchanges,
	//})
}

//func getExchanges(j *jdb.JDB) []Exchange {
//	exchanges, err := j.ReadAll("exchange")
//	utl.ErrorLog(err)
//	exs := make([]Exchange, len(exchanges))
//	for _, e := range exchanges {
//		ex := Exchange{}
//		err := j.Read("exchanges", "ex_"+e, &ex)
//		utl.ErrorLog(err)
//		exs = append(exs, ex)
//	}
//	return exs
//}

// GetExchanges reads in all of the data about all exchanges in the database
func GetExchanges(j *jdb.JDB) {
	exchanges := []Exchange{}
	err := j.Read("exchanges", "exchanges", &exchanges)
	utl.ErrorLog(err)
}

// GetBaseExchanges reads in all of the data about all exchanges in the database
func GetBaseExchanges(j *jdb.JDB) {
	//baseExchanges := []BaseExchange{}
	//err := j.Read("exchanges", "base", &baseExchanges)
	//utl.ErrorLog(err)
}
