package app

import "C"
import (
	"crypto/tls"
	"github.com/comhttp/jorm/mod/coins"
	"github.com/comhttp/jorm/mod/nodes"
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/jcelliott/lumber"
	"log"
	"net/http"
	"time"
)

const (
	// HTTPMethodOverrideHeader is a commonly used
	// http header to override a request method.
	HTTPMethodOverrideHeader = "X-HTTP-Method-Override"
	// HTTPMethodOverrideFormKey is a commonly used
	// HTML form key to override a request method.
	HTTPMethodOverrideFormKey = "_method"
)

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}
	JORM struct {
		Coins     coins.Coins
		Coin      string
		NodeCoins []string
		Explorers map[string]*Explorer
		//Hosts         map[string]Host
		WWW       *http.Server
		WS        *http.Server
		TLSconfig *tls.Config
		//CertManager autocert.Manager
		//BitNodes   map[string]nodes.BitNodes
		JDBservers map[string]string
		JDB        *jdb.JDB
		comhttp    *COMHTTP
		Log        *lumber.ConsoleLogger
		config     cfg.Config
	}
)

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

	j.JDBservers = j.config.JDBservers
	j.JDB = jdb.NewJDB(j.JDBservers)
	j.Log = lumber.NewConsoleLogger(lumber.INFO)
	//}
	j.Explorers = make(map[string]*Explorer)
	for coin, _ := range bitNodesCfg {
		j.NodeCoins = append(j.NodeCoins, coin)
		coinBitNodes := nodes.BitNodes{}
		err = c.Read("nodes", coin, &coinBitNodes)
		utl.ErrorLog(err)
		j.Explorers[coin] = NewExplorer(j.JDB, coin)
		j.Explorers[coin].BitNodes = coinBitNodes
	}

	//log.Println("Get ", cfg.C)

	//j.Coins = coin.LoadCoinsBase(j.JDB)
	j.WWW = &http.Server{
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	switch service {
	case "proxy":
		log.Println("reverse proxy")
		h := &baseHandle{}
		http.Handle("/", h)
		j.WWW.Handler = h
		j.WWW.Addr = ":" + j.config.Port["proxy"]
		return
	case "jorm":
		log.Println("jorm")
		j.JormSRV()
		j.WWW.Handler = j.JORMhandlers()
		j.WWW.Addr = ":" + j.config.Port["jorm"]
		return
	case "enso":
		log.Println("enso")
		j.WWW.Handler = j.ENSOhandlers()
		j.WWW.Addr = ":" + j.config.Port["enso"]
		return
	case "our":
		log.Println("our")
		j.WWW.Handler = j.OURhandlers()
		j.WWW.Addr = ":" + j.config.Port["our"]
		return
	//ourSRV()
	case "comhttp":
		log.Println("comhttp")
		j.WWW.Handler = j.COMHTTPhandlers()
		j.WWW.Addr = ":" + j.config.Port["comhttp"]
		return
	case "admin":
		log.Println("admin")
		j.WWW.Handler = j.ADMINhandlers()
		j.WWW.Addr = ":" + j.config.Port["admin"]
		return
	case "explorer":
		//log.Println("explorer " + *coin)
		if path == "" {
			j.ExplorerSRV(singleCoin)
			j.WWW.Addr = ":" + j.config.Port[singleCoin]
		} else {
			log.Println("Missing coin for explorer!")
		}
		return
	}

	j.JDB.Write("info", "jdbs", j.JDBservers)
	j.JDB.Write("info", "explorers", j.Explorers)

	//j.WS = &http.Server{
	//	Handler: j.WShandleR(),
	//	Addr:    ":4489",
	//	TLSConfig: &tls.Config{
	//		GetCertificate: j.CertManager.GetCertificate,
	//	},
	//}

	return j
}
