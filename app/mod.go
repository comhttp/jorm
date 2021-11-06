package app

import "C"
import (
	"fmt"
	"net/http"
	"time"

	"github.com/comhttp/jorm/mod/coin"
	"github.com/comhttp/jorm/mod/explorer"
	"github.com/comhttp/jorm/mod/nodes"
	"github.com/comhttp/jorm/mod/src/cryptocompare"
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/strapi"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/rs/zerolog/log"
)

func (j *JORM) JormSRV() {

	s := strapi.New(j.config.Strapi)

	// s.DelAll("coins")
	// s.DelAll("logos")

	//c, err := s.GetCoin("parallelcoin")
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Println("Coin: ",c)
	// c, err := j.JDBclient("coins")
	// utl.ErrorLog(err)
	// cq := coin.Queries(c, "coin")
	//exchanges.GetAllExchanges(j.JDB)
	//go csrc.GetCoinSources()
	//coin.ProcessCoins(j.JDB)
	//coins.BitNodeCoins(j.NodeCoins, j.JDB)
	//cloudflare.CloudFlare(j.config, j.JDB)
	//j.NodeCoins = coins.GetNodeCoins(j.JDB)
	//nodes.GetBitNodes(j.JDB, j.NodeCoins)
	//j.Explorer = explorers.GetExplorer(j.JDB)
	//j.Explorer.ExploreCoins(j.NodeCoins)
	//log.Print("nodessss: ", j.NodeCoins)
	//go j.Tickers()
	// cq := coin.Queries(j.JDBS.B["coins"], "coin")

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

	cc := cryptocompare.NewCryptoCompareAPI(j.config.ApiKeys["cryptocompare"])

	cc.GetAllCoins(s)
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

	ticker := time.NewTicker(99999 * time.Second)
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

func (j *JORM) OurSRV() {

	s := strapi.New(j.config.Strapi)

	fmt.Println("cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc: ")
	fmt.Println("Start OUR")
	fmt.Println("cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc: ")
	fmt.Println("cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc: ")
	// coins := coin.GetCoins(s)
	fmt.Println("Start OUR process")

	coins := s.GetAll("coins")

	c, err := j.JDBclient("coins")
	utl.ErrorLog(err)

	// go func() {
	// 	fmt.Println("Start logos import")

	// 	logos := s.GetAll("logos")
	// 	lq := coin.Queries(c, "logo")

	// 	for _, logo := range logos {
	// 		// l := logo.([]map[string]interface{})[0].(map[string]interface{})
	// 		if logo != nil {
	// 			// l := logo.(map[string]interface{})
	// 			// l := ll[0].(map[string]interface{})
	// 			lq.WriteLogo(logo["slug"].(string), logo["data"])
	// 			time.Sleep(999 * time.Millisecond)
	// 		}
	// 	}
	// 	fmt.Println("End logos import")
	// }()

	// fmt.Println("logoslogoslogoslogoslogoslogoslogoslogoslogoslogos:", logos)

	// for i, cc := range coins {
	// 	fmt.Println("coinscoinscoinscoinscoinscoinscoinscoinscoinscoinscoins:", cc)
	// 	fmt.Println("coiiiii:", i)
	// }

	cq := coin.Queries(c, "coin")

	cq.ProcessCoins(coins)

	// // cq := &coin.CoinsQueries{}
	// jdbCl, err := j.JDBclient("coins")
	// if err != nil {
	// 	utl.ErrorLog(err)
	// }
	// fmt.Println("jdbCljdbCljdbCljdbCljdbCljdbCl: ", jdbCl)
	// cq := coin.Queries(jdbCl, "coins")
	// if cq != nil {
	// 	cq.ProcessCoins(coins)
	// }

	// fmt.Println("cqcqcqcqcqcqcq: ", cq.GetAllCoins())

	// ticker := time.NewTicker(999 * time.Second)
	// quit := make(chan struct{})
	// go func() {
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			j.Tickers()
	// 			log.Print("OKNO wooikos")
	// 		case <-quit:
	// 			ticker.Stop()
	// 			return
	// 		}
	// 	}
	// }()
	fmt.Println("End OUR process")
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

	c, _ := cfg.NewCFG("/var/db/jorm", nil)
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
