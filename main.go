package main

import (
	"flag"
	"github.com/comhttp/jorm/app"
	"log"
)

func main() {
	// Get cmd line parameters
	service := flag.String("srv", "", "Service")
	path := flag.String("path", "", "Path")
	port := flag.String("port", "", "Port")
	coin := flag.String("coin", "", "Coin")
	//loglevel := flag.String("loglevel", "info", "Logging level (debug, info, warn, error)")
	flag.Parse()
	//cfg.Path = *path
	//j.Log.SetLevel(parseLogLevel(*loglevel))

	j := app.NewJORM(*service, *path, *coin)
	//j.ServicesSRV(*service, *port, *coin)
	log.Println("Service: ", *service, "Port: ", *port, "Coin: ", *coin)
	log.Fatal(j.WWW.ListenAndServe())
}

//
//jdbs -bind 192.168.192.99:14477 -dbdir jdbinfo -loglevel info &
//jdbs -bind 192.168.192.99:14488 -dbdir jdbcoins -loglevel info &
//jdbs -bind 192.168.192.99:14489 -dbdir jdbnodes -loglevel info &
//jdbs -bind 192.168.192.99:14499 -dbdir jdbexchanges -loglevel info &
//jdbs -bind 192.168.192.99:15502 -dbdir jdbparallelcoin -loglevel info &

//systemctl restart jdbparallelcoin && systemctl restart jdbcoins && systemctl restart jdbnodes && systemctl restart jdbexchanges && systemctl restart jdbinfo
//systemctl restart jorm && systemctl restart enso && systemctl restart comhttp && systemctl restart our && systemctl restart reverseproxy

//systemctl stop jdbparallelcoin && systemctl stop jdbcoins && systemctl stop jdbnodes && systemctl stop jdbexchanges && systemctl stop jdbinfo
//systemctl stop jorm && systemctl stop enso && systemctl stop our && systemctl stop comhttp  && systemctl stop reverseproxy
