package main

import (
	"flag"
	"github.com/comhttp/jorm/app"
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/sirupsen/logrus"
	"net/http"
)

var log = logrus.New()

func wrapLogger(module string) logrus.FieldLogger {
	return log.WithField("module", module)
}

func parseLogLevel(level string) logrus.Level {
	switch level {
	case "error":
		return logrus.ErrorLevel
	case "warn", "warning":
		return logrus.WarnLevel
	case "info", "notice":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "trace":
		return logrus.TraceLevel
	default:
		return logrus.InfoLevel
	}
}

func main() {
	// Get cmd line parameters
	service := flag.String("srv", "", "Service")
	path := flag.String("path", "/var/db/jorm/", "Path")
	port := flag.String("port", "14444", "Port")
	coin := flag.String("coin", "parallelcoin", "Coin")
	loglevel := flag.String("loglevel", "info", "Logging level (debug, info, warn, error)")
	flag.Parse()
	cfg.Path = *path
	log.SetLevel(parseLogLevel(*loglevel))
	err := cfg.CFG.Read("conf", "conf", &cfg.C)
	utl.ErrorLog(err)

	j := app.NewJORM()

	switch *service {
	case "proxy":
		log.Println("reverse proxy")
		h := &app.BaseHandle{}
		http.Handle("/", h)
		j.WWW.Handler = h
		j.WWW.Addr = ":" + cfg.C.Port["proxy"]
	case "jorm":
		log.Println("jorm")
		j.JormSRV()
		j.WWW.Handler = j.JORMhandlers()
		j.WWW.Addr = ":" + cfg.C.Port["jorm"]
	case "enso":
		log.Println("enso")
		j.WWW.Handler = j.ENSOhandlers()
		j.WWW.Addr = ":" + cfg.C.Port["enso"]
	case "our":
		log.Println("our")
	//ourSRV()
	case "comhttp":
		log.Println("comhttp")
		j.WWW.Handler = j.COMHTTPhandlers()
		j.WWW.Addr = ":" + cfg.C.Port["comhttp"]
	case "admin":
		log.Println("admin")
		j.WWW.Handler = j.ADMINhandlers()
		j.WWW.Addr = ":" + cfg.C.Port["admin"]
	case "explorer":
		log.Println("explorer " + *coin)
		j.ExplorerSRV(*port, *coin)
	}
	//log.Println("JORM node is on: ", ":"+*port)
	log.Fatal(j.WWW.ListenAndServe())
	//}
	//}
}

//
//jdbs -bind 192.168.192.99:14477 -dbdir jdbinfo -loglevel info &
//jdbs -bind 192.168.192.99:14488 -dbdir jdbcoins -loglevel info &
//jdbs -bind 192.168.192.99:14489 -dbdir jdbnodes -loglevel info &
//jdbs -bind 192.168.192.99:14499 -dbdir jdbexchanges -loglevel info &
//jdbs -bind 192.168.192.99:15502 -dbdir jdbparallelcoin -loglevel info &

//systemctl restart jdbparallelcoin && systemctl restart jdbcoins && systemctl restart jdbnodes && systemctl restart jdbexchanges && systemctl restart jdbinfo && systemctl restart jorm && systemctl restart enso && systemctl restart comhttp && systemctl restart our && systemctl restart reverseproxy && systemctl restart comhttp

//systemctl stop jdbparallelcoin && systemctl stop jdbcoins && systemctl stop jdbnodes && systemctl stop jdbexchanges && systemctl stop jdbinfo && systemctl stop jorm && systemctl stop enso && systemctl stop our && systemctl stop reverseproxy && systemctl stop comhttp
