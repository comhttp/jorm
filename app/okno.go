package app

import (
	"github.com/p9c/jorm/app/cfg"
	"github.com/p9c/jorm/app/jdb"
	//csrc "github.com/p9c/jorm/app/jorm/c/src"
	"github.com/p9c/jorm/pkg/utl"
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
	//jdb.JDB.Write("conf", "conf", cfg.CONFIG)
	err := jdb.JDB.Read("conf", "conf", &cfg.CONFIG)
	utl.ErrorLog(err)

	//go csrc.GetCoinSources()

	//fmt.Println(":ajdeeeeee", cfg.CONFIG)
	//go u.CloudFlare()
	o := &JORM{}
	//o.Hosts = o.GetHosts()

	srv := &http.Server{
		Handler:      o.Handler(),
		Addr:         ":" + cfg.CONFIG.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	o.Server = srv
	return o
}
