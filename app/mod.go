package app

import "C"
import (
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func (j *JORM) JormSRV() {
	//exchanges.GetAllExchanges(j.JDB)
	//go csrc.GetCoinSources(j.JDB)
	//coin.ProcessCoins(j.JDB)
	//coins.BitNodeCoins(j.NodeCoins, j.JDB)
	//cloudflare.CloudFlare(j.config, j.JDB)
	//j.NodeCoins = coins.GetNodeCoins(j.JDB)
	//nodes.GetBitNodes(j.JDB, j.NodeCoins)
	//j.Explorer = explorers.GetExplorer(j.JDB)
	//j.Explorer.ExploreCoins(j.NodeCoins)
	//log.Print("nodessss: ", j.NodeCoins)
	//go j.Tickers()

	ticker := time.NewTicker(999 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				j.Tickers()
				log.Print("OKNO wooikos")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	//log.Fatal(srv.ListenAndServeTLS("./cfg/server.pem", "./cfg/server.key"))
	//log.Print("JORM is listening on port: ", cfg.C.Port["jorm"])
	//log.Fatal(j.WWW.ListenAndServe())
	//log.Fatal(j.WS.ListenAndServeTLS("", ""))

}

func (j *JORM) ExplorerSRV(coin string) {
	log.Print("coincoincoincoincoin", coin)
	http.HandleFunc("/", status)

	ticker := time.NewTicker(12 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				//j.JDBS.ExploreCoin(j.config.RPC.Username, j.config.RPC.Password, coin)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	//log.Print("JORM explorer is listening: ", port)
	// Start HTTP server
	//log.Fatal(http.ListenAndServe(":"+port, nil))
}
