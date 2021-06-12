package app

import (
	"fmt"
	"github.com/comhttp/jorm/app/jorm/coin"
	xsrc "github.com/comhttp/jorm/app/jorm/exchange/src"

	//"github.com/comhttp/jorm/app/jorm/coin/src"
	//"github.com/comhttp/jorm/app/jorm/exchange/src"
	"github.com/comhttp/jorm/app/jorm/n"
)

func Tickers(coins coin.Coins) {
	//coins := coin.Coins{}

	fmt.Println("Cron is wooikos")
	//go e.GetExplorer(coins)
	go n.GetBitNodes(coins)
	//go csrc.GetCoinSources()
	xsrc.GetExchangeSources()
	// dsrc.GetDataSources()
}
