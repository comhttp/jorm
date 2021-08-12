package app

import (
	"github.com/comhttp/jorm/mod/coins"
	csrc "github.com/comhttp/jorm/mod/coins/src"
	"github.com/comhttp/jorm/mod/explorers"
	"github.com/comhttp/jorm/mod/nodes"
	"log"
	"net/http"
	"time"
)

func (j *JORM) JormSRV() {
	//exchanges.GetAllExchanges(j.JDB)
	csrc.GetCoinSources(j.JDB)
	coins.ProcessCoins(j.JDB)
	//coins.BitNodeCoins(j.NodeCoins, j.JDB)
	//cloudflare.CloudFlare(j.JDB)
	//j.NodeCoins = coins.GetNodeCoins(j.JDB)
	nodes.GetBitNodes(j.JDB, j.NodeCoins)
	//j.Explorer = explorers.GetExplorer(j.JDB)
	//j.Explorer.ExploreCoins(j.NodeCoins)
	//log.Println("nodessss: ", j.NodeCoins)
	//go j.Tickers()
	ticker := time.NewTicker(999 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				j.Tickers()
				log.Println("OKNO wooikos")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	//log.Fatal(srv.ListenAndServeTLS("./cfg/server.pem", "./cfg/server.key"))
	//log.Println("JORM is listening on port: ", cfg.C.Port["jorm"])
	//log.Fatal(j.WWW.ListenAndServe())
	//log.Fatal(j.WS.ListenAndServeTLS("", ""))

}

func (j *JORM) ExplorerSRV(port, coin string) {
	e := explorers.NewJORMexplorer(coin)
	log.Println("coincoincoincoincoin", coin)
	http.HandleFunc("/", status)

	ticker := time.NewTicker(12 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				explorers.ExploreCoin(e, coin)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	log.Println("JORM explorer is listening: ", port)
	// Start HTTP server
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
