package cryptocompare

import (
	"github.com/comhttp/jorm/pkg/utl"
)

type rawAllTradingPairs struct {
	Data Exchanges `json:"Data"`
}

type fsym struct {
	Tsyms map[string]tsym `json:"tsyms"`
}
type tsym struct {
	HistoMinuteStartTs int    `json:"histo_minute_start_ts"`
	HistoMinuteStart   string `json:"histo_minute_start"`
	HistoMinuteEndTs   int    `json:"histo_minute_end_ts"`
	HistoMinuteEnd     string `json:"histo_minute_end"`
}

type Exchanges struct {
	Exchanges map[string]tradingPairs `json:"exchanges"`
}
type tradingPairs struct {
	Pairs map[string]fsym `json:"pairs"`
}

func (c *cryptocompare) GetAllTradingPairs() Exchanges {
	allAllTradingPairs := &rawAllTradingPairs{}
	utl.GetSourceHeadersAPIkey(c.apiEndpoint+"data/v4/all/exchanges", c.apiKey, allAllTradingPairs)
	return allAllTradingPairs.Data
}
