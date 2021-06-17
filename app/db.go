package app

import (
	"github.com/comhttp/jorm/app/jorm/exchange"
)

func (j *JORM) WriteExchange(e *exchange.Exchange, ex exchange.ExchangeSrc) {
	for _, exs := range ex.Markets {
		mSrc := exchange.Market{}
		for _, cur := range exs.Currencies {
			mSrc.Symbol = exs.Symbol
			mSrc.Currencies = append(mSrc.Currencies, cur)
		}
		e.Markets = append(e.Markets, mSrc)
	}
	j.JDB.Write("exchanges", e.Slug, e)
}
