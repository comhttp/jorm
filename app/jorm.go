package app

import (
	"crypto/tls"
	"github.com/comhttp/jorm/mod/coin"
	"github.com/comhttp/jorm/mod/exchange"
	"github.com/comhttp/jorm/mod/explorer"
	"github.com/comhttp/jorm/mod/nodes"
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"text/template"
	"time"
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
		Coins     coin.Coins
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
)

func (j *JORM) setExplorers() {
	c, _ := cfg.NewCFG(j.config.Path, nil)
	bitNodesCfg, err := c.ReadAll("nodes")
	utl.ErrorLog(err)
	j.Explorers = make(map[string]*explorer.Explorer)
	explorereQueries := make(map[string]*explorer.ExplorerQueries)
	for coin, _ := range bitNodesCfg {
		//coins[coin] = j.JDBclient(coin)
		explorereQueries[coin] = explorer.Queries(j.JDBclient(coin), "info")
		j.NodeCoins = append(j.NodeCoins, coin)
		coinBitNodes := nodes.BitNodes{}
		err = c.Read("nodes", coin, &coinBitNodes)
		utl.ErrorLog(err)
		j.Explorers[coin] = explorereQueries[coin].NewExplorer(coin)
		j.Explorers[coin].BitNodes = coinBitNodes
	}
	//eq := explorer.Queries(coins, "info")
	cq := coin.Queries(j.JDBclient("coins"), "info")
	cq.WriteInfo("nodecoins", &coin.Coins{
		N: len(j.NodeCoins),
		C: j.NodeCoins,
	})
	return
}
func (j *JORM) JDBclient(jdbId string) *jdb.JDB {
	return jdb.NewJDB(j.jdbServers[jdbId])
}

func (j *JORM) ENSOhandlers() http.Handler {
	//coinsCollection := Queries(j.B["coins"],"coin")
	cq := coin.Queries(j.JDBclient("coins"), "coin")
	eq := exchange.Queries(j.JDBclient("exchanges"), "exchange")
	//exq := exchange.Queries(j.JDBclient("explorers"),"explorer")
	r := mux.NewRouter()
	//s := r.Host("enso.okno.rs").Subrouter()
	r.StrictSlash(true)
	//n := r.PathPrefix("/n").Subrouter()
	coin.ENSOroutes(cq, r)
	exchange.ENSOroutes(eq, r)
	//explorer.ENSOroutes(exq, r)
	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))
}

func NewJORM(service, path, singleCoin string) (j *JORM) {
	j = new(JORM)
	j.comhttp = &COMHTTP{}
	if path == "" {
		j.config.Path = "/var/db/jorm/"
	}
	c, _ := cfg.NewCFG(j.config.Path, nil)
	j.config = cfg.Config{}
	err := c.Read("conf", "conf", &j.config)
	utl.ErrorLog(err)

	j.jdbServers = make(map[string]string)
	err = c.Read("conf", "jdbs", &j.jdbServers)
	utl.ErrorLog(err)

	//ttt := j.JDBS.B["coins"].ReadAllPerPages("coin", 10, 1)
	j.setExplorers()
	j.WWW = &http.Server{
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	j.srvJORM(service, path, singleCoin)
	return j
}
