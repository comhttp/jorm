package app

import (
	"crypto/tls"
	"github.com/comhttp/jorm/app/cfg"
	"github.com/comhttp/jorm/app/jdb"
	"github.com/comhttp/jorm/app/jorm/coin"
	"golang.org/x/crypto/acme/autocert"

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

func NewJORM() *JORM {
	err := jdb.JDB.Read("conf", "conf", &cfg.C)
	utl.ErrorLog(err)

	//go u.CloudFlare()
	j := &JORM{
		Coins: coin.LoadCoinsBase(),
		CertManager: autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("ws.okno.rs", "wss.okno.rs", "ns.okno.rs"),
			Cache: autocert.DirCache(cfg.Path),
		},
	}
	j.WWW = &http.Server{
		Handler:      j.WWWhandleR(),
		Addr:         ":" + cfg.C.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	j.WS = &http.Server{
		Handler:      j.WShandleR(),
		Addr: ":4489",
		TLSConfig: &tls.Config{
			GetCertificate: j.CertManager.GetCertificate,
		},
	}

	return j
}
