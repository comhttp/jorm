package main

import (
	"fmt"
	"github.com/comhttp/jorm/app"
	"github.com/comhttp/jorm/mod/coins"
	csrc "github.com/comhttp/jorm/mod/coins/src"
	"github.com/comhttp/jorm/mod/explorer"
	cfg "github.com/comhttp/jorm/pkg/cfg"
	"time"

	"github.com/comhttp/jorm/mod/exchanges"
	//csrc "github.com/comhttp/jorm/coins/src"
	"log"
	//"log"
	//"time"
)

func main() {
	j := app.NewJORM()
	exchanges.GetAllExchanges(j.JDB)

	csrc.GetCoinSources(j.JDB)

	coins.ProcessCoins(j.JDB)

	//cloudflare.CloudFlare(j.JDB)

	j.NodeCoins = coins.GetNodeCoins(j.JDB)
	//nodes.GetBitNodes(j.JDB, j.NodeCoins)
	j.Explorer = explorer.GetExplorer(j.JDB)
	//j.Explorer.ExploreCoins(j.NodeCoins)
	//fmt.Println("nodessss: ", j.NodeCoins)

	//go j.Tickers()
	ticker := time.NewTicker(23 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				j.Tickers()
				fmt.Println("OKNO wooikos")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	//log.Fatal(srv.ListenAndServeTLS("./cfg/server.pem", "./cfg/server.key"))
	fmt.Println("JORM is listening on port: ", cfg.C.Port["jorm"])
	log.Fatal(j.WWW.ListenAndServe())
	//log.Fatal(j.WS.ListenAndServeTLS("", ""))

	// port := 9898
	// fmt.Println("Listening on port:", port)
	// log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(r)))

}
