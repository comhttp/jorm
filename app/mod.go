package app

import "C"
import (
	"fmt"
	"github.com/comhttp/jorm/mod/explorer"
	"github.com/comhttp/jorm/mod/nodes"
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
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
	//cq := coin.Queries(j.JDBS.B["coins"], "coin")

	//eq := exchange.Queries(j.JDBS.B["exchanges"], "exchange")
	//xsrc.GetPoloniexExchange(eq)
	//xsrc.GetDexTradeExchange(eq)

	//for _, t := range ttt {
	//	if t["slug"] != "" {
	//	} else {
	//		fmt.Println("ssssssssssssssttttttttt2222222222ttttttt",t)
	//	}
	//
	//}

	//cc := cryptocompare.NewCryptoCompareAPI(j.config.ApiKeys["cryptocompare"])
	//cc.GetAllCoins(cq)
	//cc.GetAllExchanges(eq)
	//minerstat.GetAllCoins(cq)

	//cm := coinmarketcap.NewCoinMarketCapAPI(j.config.ApiKeys["coinmarketcap"])
	//cm.GetAllCoins(cq)

	//coingecko.GetAllCoins(cq)
	//log.Print("ssssssssssssssttttttttt2222222222tttttttccccccccccccccccccc", ccc)
	//spew.Dump(ccc.N)
	//for _,ss:=range sss.C{
	//	log.Print("SSSSSS: ",ss)
	//}
	//ccc := cq.GetCoin("parallelcoin")
	//log.Print("Slug: ", ccc.Slug)
	//log.Print("Name: ", ccc.Name)
	//log.Print("Ticker: ", ccc.Symbol)
	//log.Print("Description: ", ccc.Description)
	//log.Print("Algo: ", ccc.Algo)
	//log.Print("Proof: ", ccc.Proof)
	//log.Print("Start: ", ccc.GenesisDate.String())

	//cq.ProcessCoins()

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
	log.Print("Coin: ", coin)
	http.HandleFunc("/", status)
	//info := explorer.Queries(j.JDBS, "info")
	jdbCl, err := j.JDBclient(coin)
	utl.ErrorLog(err)
	jdbs := map[string]*jdb.JDB{
		coin: jdbCl,
	}
	c, _ := cfg.NewCFG(j.config.Path, nil)
	coinBitNodes := nodes.BitNodes{}
	err = c.Read("nodes", coin, &coinBitNodes)
	utl.ErrorLog(err)
	eq := explorer.Queries(jdbs, "info")

	j.Explorers = make(map[string]*explorer.Explorer)
	j.Explorers[coin] = eq.NewExplorer(coin)
	j.Explorers[coin].BitNodes = coinBitNodes

	//info.status = info.GetStatus()
	j.Explorers[coin].Status, err = eq.GetStatus(coin)
	utl.ErrorLog(err)
	fmt.Println("ssss", j.Explorers[coin].Status)
	ticker := time.NewTicker(12 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:

				eq.ExploreCoin(j.Explorers[coin].BitNodes, j.config.RPC.Username, j.config.RPC.Password, coin)
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
