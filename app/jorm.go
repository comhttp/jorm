package app

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/comhttp/jorm/mod/coin"
	"github.com/comhttp/jorm/mod/exchange"
	"github.com/comhttp/jorm/mod/explorer"
	"github.com/comhttp/jorm/mod/nodes"
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//const (
//	// HTTPMethodOverrideHeader is a commonly used
//	// http header to override a request method.
//	HTTPMethodOverrideHeader = "X-HTTP-Method-Override"
//	// HTTPMethodOverrideFormKey is a commonly used
//	// HTML form key to override a request method.
//	HTTPMethodOverrideFormKey = "_method"
//)

type (
	JORM struct {
		Coins     []string
		Coin      string
		NodeCoins []string
		Explorers map[string]*explorer.Explorer
		//Hosts         map[string]Host
		WWW       *http.Server
		WS        *http.Server
		TLSconfig *tls.Config
		//CertManager autocert.Manager
		//BitNodes   map[string]nodes.BitNodes
		//JDBservers map[string]string
		//JDBS       *jdb.JDBS
		comhttp    *COMHTTP
		goHTML     *template.Template
		config     cfg.Config
		jdbServers map[string]string
	}
	Index struct {
		Slug string      `json:"slug"`
		Data interface{} `json:"data"`
	}
)

func (j *JORM) setExplorers() {
	c, _ := cfg.NewCFG(j.config.Path, nil)
	bitNodesCfg, err := c.ReadAll("nodes")
	utl.ErrorLog(err)

	j.Explorers = make(map[string]*explorer.Explorer)
	explorerJDBS := make(map[string]*jdb.JDB)
	for coin, _ := range bitNodesCfg {
		//coins[coin] = j.JDBclient(coin)
		jdbCl, err := j.JDBclient(coin)
		if err != nil {
			utl.ErrorLog(err)
		} else {
			explorerJDBS[coin] = jdbCl
			j.NodeCoins = append(j.NodeCoins, coin)
			coinBitNodes := nodes.BitNodes{}
			err = c.Read("nodes", coin, &coinBitNodes)
			utl.ErrorLog(err)
			eq := explorer.Queries(explorerJDBS, "info")
			j.Explorers[coin] = eq.NewExplorer(coin)
			j.Explorers[coin].BitNodes = coinBitNodes

		}
	}
	//eq := explorer.Queries(coins, "info")
	jdbCl, err := j.JDBclient("coins")
	if err != nil {
		utl.ErrorLog(err)
	} else {
		cq := coin.Queries(jdbCl, "info")
		cq.WriteInfo("nodecoins", &coin.Coins{
			N: len(j.NodeCoins),
			C: j.NodeCoins,
		})
	}
	return
}
func (j *JORM) JDBclient(jdbId string) (*jdb.JDB, error) {
	return jdb.NewJDB(j.jdbServers[jdbId])
}

// func (j *JORM) STRAPIhandler() http.Handler {
// 	r := mux.NewRouter()
// 	r.StrictSlash(true)
// 	r.Headers()
// 	n := r.PathPrefix("/n").Subrouter()
// 	//n.HandleFunc("/{coin}/nodes", cq.CoinNodesHandler).Methods("GET")
// 	n.HandleFunc("/{coin}/{nodeip}", cq.nodeHandler).Methods("GET")

// 	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))
// }
func (j *JORM) ENSOhandlers() http.Handler {
	//coinsCollection := Queries(j.B["coins"],"coin")
	c, err := j.JDBclient("coins")
	utl.ErrorLog(err)
	cq := coin.Queries(c, "coin")

	e, err := j.JDBclient("exchanges")
	utl.ErrorLog(err)
	eq := exchange.Queries(e, "exchange")

	explorerJDBS := make(map[string]*jdb.JDB)
	rpcBitNodes := make(map[string]nodes.BitNodes)

	fmt.Println("j.Explorers j.Explorers j.Explorers j.Explorers : ", j.Explorers)

	for _, coin := range j.Explorers {
		fmt.Println("coincoincoin: ", coin)

		jdbCl, err := j.JDBclient(coin.Coin)
		if err != nil {
			utl.ErrorLog(err)
		} else {
			explorerJDBS[coin.Coin] = jdbCl
		}
		err = c.Read("nodes", coin.Coin, &coin.BitNodes)
		utl.ErrorLog(err)
		rpcBitNodes[coin.Coin] = coin.BitNodes
		fmt.Println("coincoincoin: ", coin)

	}

	exq := explorer.Queries(explorerJDBS, "info")

	//exq := exchange.Queries(j.JDBclient("exchanges"), "exchange")
	//exq := exchange.Queries(j.JDBclient("explorers"),"explorer")
	r := mux.NewRouter()
	//s := r.Host("enso.okno.rs").Subrouter()
	r.StrictSlash(true)
	//n := r.PathPrefix("/n").Subrouter()

	coin.ENSOroutes(cq, r)
	exchange.ENSOroutes(eq, r)
	explorer.ENSOroutes(exq, r)

	nodes.ENSOroutesDirect(j.config.Path, j.config.RPC, r)

	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))
}

func NewJORM(service, path, singleCoin string) (j *JORM) {
	j = new(JORM)
	j.comhttp = &COMHTTP{}
	j.config.Path = path
	c, _ := cfg.NewCFG(j.config.Path, nil)
	err := c.Read("conf", "conf", &j.config)
	utl.ErrorLog(err)

	j.jdbServers = make(map[string]string)
	err = c.Read("conf", "jdbs", &j.jdbServers)
	utl.ErrorLog(err)
	//ttt := j.JDBS.B["coins"].ReadAllPerPages("coin", 10, 1)
	j.WWW = &http.Server{
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	j.srvJORM(service, path, singleCoin)
	return j
}
