package app

import (
	"crypto/tls"
	"github.com/comhttp/jorm/mod/coins"
	"github.com/comhttp/jorm/mod/explorers"
	"github.com/comhttp/jorm/mod/nodes"
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/jcelliott/lumber"

	//csrc "github.com/comhttp/jorm/app/jorm/c/src"
	"github.com/comhttp/jorm/pkg/utl"
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
		NodeCoins nodes.NodeCoins
		Explorer  *explorers.Explorer
		//Hosts         map[string]Host
		WWW       *http.Server
		WS        *http.Server
		TLSconfig *tls.Config
		//CertManager autocert.Manager
		BitNodes   map[string]nodes.BitNodes
		JDBservers map[string]string
		JDB        *jdb.JDB
		Log        Logger
	}
)

func NewJORM() *JORM {
	err := cfg.CFG.Read("conf", "conf", &cfg.C)
	utl.ErrorLog(err)
	bitNodesCfg, err := cfg.CFG.ReadAll("nodes")
	utl.ErrorLog(err)
	bitNodes := make(map[string]nodes.BitNodes)

	for _, coin := range bitNodesCfg {
		coinBitNodes := nodes.BitNodes{}
		err := cfg.CFG.Read("nodes", coin, &coinBitNodes)
		utl.ErrorLog(err)
		bitNodes[coin] = coinBitNodes
	}
	//fmt.Println("Get ", cfg.C)
	j := &JORM{
		//CertManager: autocert.Manager{
		//	Prompt:     autocert.AcceptTOS,
		//	HostPolicy: autocert.HostWhitelist("ws.okno.rs", "wss.okno.rs", "ns.okno.rs"),
		//	Cache:      autocert.DirCache(cfg.Path),
		//},
		BitNodes:   bitNodes,
		JDBservers: cfg.C.JDBservers,
		JDB:        jdb.NewJDB(cfg.C.JDBservers),
		Log:        lumber.NewConsoleLogger(lumber.INFO),
	}
	//j.Coins = coin.LoadCoinsBase(j.JDB)
	j.WWW = &http.Server{
		Handler:      j.WWWhandleR(),
		Addr:         ":" + cfg.C.Port["jorm"],
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	j.JDB.Write("info", "jdbs", j.JDBservers)
	j.JDB.Write("info", "bitnodes", j.BitNodes)

	//j.WS = &http.Server{
	//	Handler: j.WShandleR(),
	//	Addr:    ":4489",
	//	TLSConfig: &tls.Config{
	//		GetCertificate: j.CertManager.GetCertificate,
	//	},
	//}

	return j
}
