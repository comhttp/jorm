package app

import (
	"fmt"
	xsrc "github.com/comhttp/jorm/app/jorm/exchange/src"
	"github.com/comhttp/jorm/app/jorm/explorer"
	"github.com/comhttp/jorm/app/jorm/n"
	//"github.com/comhttp/jorm/app/jorm/coin/src"
	//"github.com/comhttp/jorm/app/jorm/exchange/src"
	//"github.com/comhttp/jorm/app/jorm/n"
)

func (j *JORM) Tickers() {
	//coins := coin.Coins{}

	fmt.Println("Cron is wooikos")
	//go csrc.GetCoinSources(j.JDB)
	go explorer.GetExplorer(j.JDB, j.Coins)
	go n.GetBitNodes(j.Coins)

	go xsrc.GetExchangeSources(j.JDB)
	// dsrc.GetDataSources()
}
