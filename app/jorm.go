package app

import (
	"github.com/comhttp/jorm/app/cfg"
	"github.com/comhttp/jorm/app/jdb"
	"github.com/comhttp/jorm/app/jorm/coin"

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
	o := &JORM{
		Coins: coin.LoadCoinsBase(),
	}

	srv := &http.Server{
		Handler:      o.Handler(),
		Addr:         ":" + cfg.C.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	o.Server = srv
	return o
}
