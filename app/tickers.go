package app

import (
	"fmt"
	"github.com/p9c/jorm/app/jorm/coin"
	csrc "github.com/p9c/jorm/app/jorm/coin/src"

	//"github.com/p9c/jorm/app/jorm/coin/src"
	//"github.com/p9c/jorm/app/jorm/exchange/src"
	"github.com/p9c/jorm/app/jorm/n"
)

func Tickers(coins coin.Coins) {
	//coins := coin.Coins{}

	fmt.Println("Cron is wooikos")
	//go e.GetExplorer(coins)
	go n.GetBitNodes(coins)
	go csrc.GetCoinSources()
	//go xsrc.GetExchangeSources()
	// dsrc.GetDataSources()
}
