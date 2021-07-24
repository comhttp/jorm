package app

import (
	"fmt"
	csrc "github.com/comhttp/jorm/mod/coins/src"
	xsrc "github.com/comhttp/jorm/mod/exchanges/src"
	"github.com/comhttp/jorm/mod/nodes"
	//"github.com/comhttp/jorm/app/jorm/coin"
	//xsrc "github.com/comhttp/jorm/app/jorm/exchange/src"
	//"github.com/comhttp/jorm/app/jorm/n"
	//"github.com/comhttp/jorm/app/jorm/coin/src"
	//"github.com/comhttp/jorm/app/jorm/exchange/src"
	//"github.com/comhttp/jorm/app/jorm/n"
)

func (j *JORM) Tickers() {
	//coins := coin.Coins{}
	go nodes.GetBitNodes(j.JDB, j.NodeCoins)

	go j.Explorer.ExploreCoins(j.NodeCoins)

	fmt.Println("Cron is wooikos")
	//go explorer.GetExplorer(coins)
	//n.GetBitNodes(coins)
	//go csrc.GetCoinSources()
	go xsrc.GetExchangeSources(j.JDB)

	csrc.GetCoinSources(j.JDB)

	// dsrc.GetDataSources()
}
