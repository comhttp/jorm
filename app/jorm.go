package app

import (
	"crypto/tls"
	"fmt"
	"github.com/comhttp/jorm/app/cfg"
	"github.com/comhttp/jorm/app/jdb"
	"github.com/comhttp/jorm/app/jorm/coin"

	//csrc "github.com/comhttp/jorm/app/jorm/c/src"
	"github.com/comhttp/jorm/pkg/utl"
	"golang.org/x/crypto/acme/autocert"
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
		Coins:     coin.LoadCoinsBase(),
		TLSconfig: &tls.Config{},
	}

	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		//HostPolicy: autocert.HostWhitelist("example.com"),
		Cache: autocert.DirCache("./certs"), //Folder for storing certificates
	}

	//
	//j.TLSconfig.Certificates = make([]tls.Certificate, 3)
	//// go http server treats the 0'th key as a default fallback key
	//j.TLSconfig.Certificates[0], err = tls.LoadX509KeyPair("test0.pem", "key.pem")
	//if err != nil {
	//	fmt.Println("err:", err )
	//}
	//j.TLSconfig.Certificates[1], err = tls.LoadX509KeyPair("test1.pem", "key.pem")
	//if err != nil {
	//	fmt.Println("err:", err )
	//}
	//j.TLSconfig.Certificates[2], err = tls.LoadX509KeyPair("test2.pem", "key.pem")
	//if err != nil {
	//	fmt.Println("err:", err )
	//}
	//j.TLSconfig.BuildNameToCertificate()
	srv := &http.Server{
		Handler: j.Handler(),
		//Addr:         ":" + cfg.C.Port,
		Addr:           ":https",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		//TLSConfig:      j.TLSconfig,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	j.Server = srv
	return j
}
