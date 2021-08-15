package app

import "C"
import (
	"crypto/tls"
	"fmt"
	"github.com/comhttp/jorm/mod/cloudflare"
	"github.com/comhttp/jorm/mod/coin"
	"github.com/comhttp/jorm/mod/explorer"
	"github.com/comhttp/jorm/mod/nodes"
	"github.com/comhttp/jorm/mod/src/cryptocompare"
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
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
		JDBservers map[string]string
		JDBS       *jdb.JDBS
		comhttp    *COMHTTP
		goHTML     *template.Template
		config     cfg.Config
	}
)

func (j *JORM) ENSOhandlers() http.Handler {
	//coinsCollection := Queries(j.B["coins"],"coin")
	r := mux.NewRouter()
	//s := r.Host("enso.okno.rs").Subrouter()
	r.StrictSlash(true)

	//n := r.PathPrefix("/n").Subrouter()
	coin.ENSOroutes(j.JDBS, r)
	explorer.ENSOroutes(j.JDBS, r)
	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))
}

func NewJORM(service, path, singleCoin string) (j *JORM) {
	j = new(JORM)
	if path == "" {
		j.config.Path = "/var/db/jorm/"
	}
	c, _ := cfg.NewCFG(j.config.Path, nil)
	j.config = cfg.Config{}
	err := c.Read("conf", "conf", &j.config)
	utl.ErrorLog(err)
	//j.config = confSRC
	bitNodesCfg, err := c.ReadAll("nodes")
	utl.ErrorLog(err)
	//bitNodes := make(map[string]nodes.BitNodes)

	j.comhttp = &COMHTTP{}
	//j := &JORM{
	//CertManager: autocert.Manager{
	//	Prompt:     autocert.AcceptTOS,
	//	HostPolicy: autocert.HostWhitelist("ws.okno.rs", "wss.okno.rs", "ns.okno.rs"),
	//	Cache:      autocert.DirCache(cfg.Path),
	//},
	log.Print("pre: ", path)

	//j.goHTML = j.parseTemplates("amp", j.goHTML)
	log.Print("posle: ", path)

	j.JDBservers = j.config.JDBservers
	j.JDBS = jdb.NewJDBS(j.JDBservers)
	//}

	ttt := j.JDBS.B["coins"].ReadAllPerPages("coin", 10, 1)

	for _, t := range ttt {
		if t["slug"] != "" {
		} else {
			//fmt.Println("ssssssssssssssttttttttt2222222222ttttttt",t)
		}

	}
	cc := cryptocompare.NewCryptoCompareAPI(j.config.ApiKeys["cryptocompare"])

	ccc := cc.GetAllCoins()
	fmt.Println("ssssssssssssssttttttttt2222222222tttttttccccccccccccccccccc", ccc)
	j.Explorers = make(map[string]*explorer.Explorer)
	for coin, _ := range bitNodesCfg {
		j.NodeCoins = append(j.NodeCoins, coin)
		coinBitNodes := nodes.BitNodes{}
		err = c.Read("nodes", coin, &coinBitNodes)
		utl.ErrorLog(err)
		j.Explorers[coin] = explorer.NewExplorer(j.JDBS.B[coin], coin)
		j.Explorers[coin].BitNodes = coinBitNodes
	}

	//log.Print("Get ", cfg.C)

	//j.Coins = coin.LoadCoinsBase(j.JDB)
	j.WWW = &http.Server{
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	switch service {
	case "proxy":
		log.Print("reverse proxy")
		h := &baseHandle{}
		http.Handle("/", h)
		j.WWW.Handler = h
		j.WWW.Addr = ":" + j.config.Port["proxy"]
		log.Fatal().Err(j.WWW.ListenAndServe())
	case "jorm":
		log.Print("jorm")
		j.JormSRV()
		j.WWW.Handler = j.JORMhandlers()
		j.WWW.Addr = ":" + j.config.Port["jorm"]
		log.Fatal().Err(j.WWW.ListenAndServe())
	case "enso":
		log.Print("enso")
		j.WWW.Handler = j.ENSOhandlers()
		j.WWW.Addr = ":" + j.config.Port["enso"]
		log.Fatal().Err(j.WWW.ListenAndServe())
	case "our":
		log.Print("our")
		j.WWW.Handler = j.OURhandlers()
		j.WWW.Addr = ":" + j.config.Port["our"]
		log.Fatal().Err(j.WWW.ListenAndServe())
	//ourSRV()
	case "comhttp":
		log.Print("comhttp")
		j.WWW.Handler = j.COMHTTPhandlers()
		j.WWW.Addr = ":" + j.config.Port["comhttp"]
		log.Fatal().Err(j.WWW.ListenAndServe())
	case "admin":
		log.Print("admin")
		j.WWW.Handler = j.ADMINhandlers()
		j.WWW.Addr = ":" + j.config.Port["admin"]
		log.Fatal().Err(j.WWW.ListenAndServe())
	case "explorer":
		//log.Print("explorer " + *coin)
		if path == "" {
			j.ExplorerSRV(singleCoin)
			j.WWW.Addr = ":" + j.config.Port[singleCoin]
		} else {
			log.Print("Missing coin for explorer!")
		}
		log.Fatal().Err(j.WWW.ListenAndServe())
	case "cloudflare":
		log.Print("cloudflare")
		cloudflare.CloudFlare(j.config, &coin.CoinsShort{})
		return
	}

	j.JDBS.B["info"].Write("info", "jdbs", j.JDBservers)
	j.JDBS.B["info"].Write("info", "explorers", j.Explorers)

	//j.WS = &http.Server{
	//	Handler: j.WShandleR(),
	//	Addr:    ":4489",
	//	TLSConfig: &tls.Config{
	//		GetCertificate: j.CertManager.GetCertificate,
	//	},
	//}

	return j
}
