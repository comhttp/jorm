package app

import (
	"crypto/tls"
	"github.com/comhttp/jorm/mod/coins"
	"github.com/comhttp/jorm/mod/nodes"
	cfg "github.com/comhttp/jorm/pkg/cfg"
	jdb2 "github.com/comhttp/jorm/pkg/jdb"

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

type JORM struct {
	Coins     coins.Coins
	NodeCoins nodes.NodeCoins
	//Hosts         map[string]Host
	WWW       *http.Server
	WS        *http.Server
	TLSconfig *tls.Config
	//CertManager autocert.Manager
	JDB *jdb2.JDB
}

func NewJORM() *JORM {
	err := cfg.CFG.Read("conf", "conf", &cfg.C)
	utl.ErrorLog(err)

	//fmt.Println("Get ", cfg.C)
	j := &JORM{
		//CertManager: autocert.Manager{
		//	Prompt:     autocert.AcceptTOS,
		//	HostPolicy: autocert.HostWhitelist("ws.okno.rs", "wss.okno.rs", "ns.okno.rs"),
		//	Cache:      autocert.DirCache(cfg.Path),
		//},
		JDB: jdb2.NewJDB(cfg.C.JDBservers),
	}
	//j.Coins = coin.LoadCoinsBase(j.JDB)
	j.WWW = &http.Server{
		Handler:      j.WWWhandleR(),
		Addr:         ":" + cfg.C.Port["jorm"],
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	//j.WS = &http.Server{
	//	Handler: j.WShandleR(),
	//	Addr:    ":4489",
	//	TLSConfig: &tls.Config{
	//		GetCertificate: j.CertManager.GetCertificate,
	//	},
	//}

	return j
}
