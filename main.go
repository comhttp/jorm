package main

import (
	"flag"
	"fmt"
	"github.com/comhttp/jorm/app"
	"github.com/comhttp/jorm/mod/coins"
	"github.com/comhttp/jorm/mod/enso"
	"github.com/comhttp/jorm/mod/explorers"
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
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
	//command := flag.String("cmd", "", "Command")
	service := flag.String("srv", "", "Service")
	path := flag.String("path", "./", "Path")
	port := flag.String("port", "14444", "Port")
	coin := flag.String("coin", "parallelcoin", "Coin")
	loglevel := flag.String("loglevel", "info", "Logging level (debug, info, warn, error)")
	flag.Parse()
	cfg.Path = *path
	log.SetLevel(parseLogLevel(*loglevel))
	err := cfg.CFG.Read("conf", "conf", &cfg.C)
	utl.ErrorLog(err)
	//fmt.Println("jorm", c)

	//err = daemon.Init(*service, map[string]interface{}{}, "./"+*service+".pid")
	//if err != nil {
	//	return
	//}
	//switch *command {
	//case "start":
	//	err = daemon.Start()
	//case "stop":
	//	err = daemon.Stop()
	//case "restart":
	//	err = daemon.Stop()
	//	err = daemon.Start()
	//case "status":
	//	status := "stopped"
	//	if daemon.IsRun() {
	//		status = "started"
	//	}
	//	fmt.Printf("Application is %s\n", status)
	//	return
	//case "":
	//default:
	//www := &http.Server{
	//	//Handler:      j.WWWhandleR(),
	//	Addr:         ":" + *port,
	//	WriteTimeout: 15 * time.Second,
	//	ReadTimeout:  15 * time.Second,
	//}
	switch *service {
	case "jorm":
		fmt.Println("jorm")
		jormSRV()
	case "enso":
		fmt.Println("enso")
		ensoSRV()
	case "our":
		fmt.Println("our")
	case "explorer":
		fmt.Println("explorer " + *coin)
		explorerSRV(*port, *coin)
	}
	fmt.Println("JORM node is on: ", ":"+*port)
	//log.Fatal(www.ListenAndServe())
	//}
}

//
//import (
//	"fmt"
//	"github.com/comhttp/jorm/app"
//	"github.com/comhttp/jorm/mod/coins"
//	csrc "github.com/comhttp/jorm/mod/coins/src"
//	cfg "github.com/comhttp/jorm/pkg/cfg"
//	"time"
//
//	"github.com/comhttp/jorm/mod/exchanges"
//	//csrc "github.com/comhttp/jorm/coins/src"
//	"log"
//	//"log"
//	//"time"
//)
//
func jormSRV() {
	j := app.NewJORM()
	//exchanges.GetAllExchanges(j.JDB)
	//csrc.GetCoinSources(j.JDB)
	//coins.ProcessCoins(j.JDB)
	//cloudflare.CloudFlare(j.JDB)
	j.NodeCoins = coins.GetNodeCoins(j.JDB)
	//nodes.GetBitNodes(j.JDB, j.NodeCoins)
	//j.Explorer = explorers.GetExplorer(j.JDB)
	//j.Explorer.ExploreCoins(j.NodeCoins)
	//fmt.Println("nodessss: ", j.NodeCoins)
	//go j.Tickers()
	ticker := time.NewTicker(999 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				j.Tickers()
				fmt.Println("OKNO wooikos")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	//log.Fatal(srv.ListenAndServeTLS("./cfg/server.pem", "./cfg/server.key"))
	fmt.Println("JORM is listening on port: ", cfg.C.Port["jorm"])
	log.Fatal(j.WWW.ListenAndServe())
	//log.Fatal(j.WS.ListenAndServeTLS("", ""))

}

func ensoSRV() {
	j := enso.NewENSO()

	fmt.Println("Listening on port: ", cfg.C.Port["enso"])
	log.Fatal(j.WWW.ListenAndServe())
}

//jdbs -bind 192.168.192.99:14477 -dbdir jdbinfo -loglevel info &
//jdbs -bind 192.168.192.99:14488 -dbdir jdbcoins -loglevel info &
//jdbs -bind 192.168.192.99:14489 -dbdir jdbnodes -loglevel info &
//jdbs -bind 192.168.192.99:14499 -dbdir jdbexchanges -loglevel info &
//jdbs -bind 192.168.192.99:15502 -dbdir jdbparallelcoin -loglevel info &

func explorerSRV(port, coin string) {
	e := explorers.NewJORMexplorer(coin)
	fmt.Println("coincoincoincoincoin", coin)
	http.HandleFunc("/", status)

	ticker := time.NewTicker(12 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				explorers.ExploreCoin(e, coin)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	fmt.Println("JORM explorer is listening: ", port)
	// Start HTTP server
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func status(w http.ResponseWriter, r *http.Request) {
	// Handles top-level page.
	fmt.Fprintf(w, "You are on the status home page")
}
