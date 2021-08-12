package main

import (
	"flag"
	"fmt"
	"github.com/comhttp/jorm/app"
	"github.com/comhttp/jorm/mod/coins"
	csrc "github.com/comhttp/jorm/mod/coins/src"
	"github.com/comhttp/jorm/mod/comhttp"
	"github.com/comhttp/jorm/mod/enso"
	"github.com/comhttp/jorm/mod/explorers"
	"github.com/comhttp/jorm/mod/nodes"
	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
	"net/url"
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
	path := flag.String("path", "/var/db/jorm/", "Path")
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
	case "proxy":
		fmt.Println("reverse proxy")
		reverseProxySRV()
	case "jorm":
		fmt.Println("jorm")
		jormSRV()
	case "enso":
		fmt.Println("enso")
		ensoSRV()
	case "our":
		fmt.Println("our")
	//ourSRV()
	case "comhttp":
		fmt.Println("comhttp")
		comhttpSRV()
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
	csrc.GetCoinSources(j.JDB)
	coins.ProcessCoins(j.JDB)
	//coins.BitNodeCoins(j.NodeCoins, j.JDB)
	//cloudflare.CloudFlare(j.JDB)
	//j.NodeCoins = coins.GetNodeCoins(j.JDB)
	nodes.GetBitNodes(j.JDB, j.NodeCoins)
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
	fmt.Println("cfg.Pathcfg.Pathcfg.Pathsssssssssss1111s", cfg.Path)

	j := enso.NewENSO()
	fmt.Println("cfg.Pathcfg.Pathcfg.Pathssssssssssss2222", cfg.Path)

	fmt.Println("Listening on port: ", cfg.C.Port["enso"])
	log.Fatal(j.WWW.ListenAndServe())
}

func ourSRV() {
	srv := &http.Server{
		Handler:      comhttp.Handlers(),
		Addr:         ":" + cfg.C.Port["our"],
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Listening on port: ", cfg.C.Port["our"])
	log.Fatal(srv.ListenAndServe())
}
func comhttpSRV() {
	srv := &http.Server{
		Handler:      comhttp.Handlers(),
		Addr:         ":" + cfg.C.Port["comhttp"],
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Listening on port: ", cfg.C.Port["comhttp"])
	log.Fatal(srv.ListenAndServe())
}

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

var (
	hostTarget = map[string]string{
		"okno.rs":                    "http://localhost:4433",
		"parallelcoin.info":          "http://localhost:4433",
		"explorer.parallelcoin.info": "http://localhost:4433",
		"jorm.okno.rs":               "http://localhost:14411",
		"our.okno.rs":                "http://localhost:14422",
		"enso.okno.rs":               "http://localhost:14433",
	}
)

type baseHandle struct{}

func (h *baseHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host

	if target, ok := hostTarget[host]; ok {
		reverseproxy(w, r, target)
	} else {
		reverseproxy(w, r, host)
	}
	w.Write([]byte("403: Host forbidden " + host))
}

func reverseproxy(w http.ResponseWriter, r *http.Request, target string) {
	remoteUrl, err := url.Parse(target)
	if err != nil {
		log.Println("target parse fail:", err)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
	proxy.ServeHTTP(w, r)
	return
}

func reverseProxySRV() {
	h := &baseHandle{}
	http.Handle("/", h)
	server := &http.Server{
		Addr:    ":80",
		Handler: h,
	}
	log.Fatal(server.ListenAndServe())
}
func status(w http.ResponseWriter, r *http.Request) {
	// Handles top-level page.
	fmt.Fprintf(w, "You are on the status home page")
}

//jdbs -bind 192.168.192.99:14477 -dbdir jdbinfo -loglevel info &
//jdbs -bind 192.168.192.99:14488 -dbdir jdbcoins -loglevel info &
//jdbs -bind 192.168.192.99:14489 -dbdir jdbnodes -loglevel info &
//jdbs -bind 192.168.192.99:14499 -dbdir jdbexchanges -loglevel info &
//jdbs -bind 192.168.192.99:15502 -dbdir jdbparallelcoin -loglevel info &

//systemctl restart jdbparallelcoin && systemctl restart jdbcoins && systemctl restart jdbnodes && systemctl restart jdbexchanges && systemctl restart jdbinfo && systemctl restart jorm && systemctl restart enso && systemctl restart comhttp && systemctl restart our && systemctl restart reverseproxy

//systemctl stop jdbparallelcoin && systemctl stop jdbcoins && systemctl stop jdbnodes && systemctl stop jdbexchanges && systemctl stop jdbinfo && systemctl stop jorm && systemctl stop enso && systemctl stop our && systemctl stop reverseproxy
