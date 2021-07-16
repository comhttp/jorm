package main

import (
	"fmt"
	"github.com/comhttp/jorm/app"
	"github.com/comhttp/jorm/mod/cloudflare"
	"github.com/comhttp/jorm/mod/coins"
	cfg2 "github.com/comhttp/jorm/pkg/cfg"

	//csrc "github.com/comhttp/jorm/coins/src"
	"log"
	//"github.com/comhttp/jorm/app/jorm/exchange"
	//"log"
	//"time"
)

func main() {
	j := app.NewJORM()
	//exchange.ReadAllExchanges()

	//csrc.GetCoinSources(j.JDB)

	coins.ProcessCoins(j.JDB)

	cloudflare.CloudFlare(j.JDB)

	//coins.BitNodeCoins(j.JDB)
	//j.NodeCoins = coins.GetNodeCoins(j.JDB)
	//nodes.GetBitNodes(j.JDB, j.NodeCoins)
	//e := explorer.GetExplorer(j.JDB)
	//e.ExploreCoins(j.NodeCoins)
	//fmt.Println("nodessss: ", j.NodeCoins)

	//go j.Tickers()
	//ticker := time.NewTicker(999 * time.Second)
	//quit := make(chan struct{})
	//go func() {
	//	for {
	//		select {
	//		case <-ticker.C:
	//			j.Tickers()
	//			fmt.Println("OKNO wooikos")
	//		case <-quit:
	//			ticker.Stop()
	//			return
	//		}
	//	}
	//}()
	//log.Fatal(srv.ListenAndServeTLS("./cfg/server.pem", "./cfg/server.key"))
	fmt.Println("JORM is listening on port: ", cfg2.C.Port["jorm"])
	log.Fatal(j.WWW.ListenAndServe())
	//log.Fatal(j.WS.ListenAndServeTLS("", ""))

	// port := 9898
	// fmt.Println("Listening on port:", port)
	// log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(r)))

}
