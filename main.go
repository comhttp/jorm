package main

import (
	"fmt"
	"github.com/comhttp/jorm/app"
	"github.com/comhttp/jorm/app/jorm/exchange"
	"log"
	"net/http"
	"time"
)

func main() {
	jorm := app.NewJORM()
	exchange.ReadAllExchanges()
	go app.Tickers(jorm.Coins)
	ticker := time.NewTicker(999 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				app.Tickers(jorm.Coins)
				fmt.Println("OKNO wooikos")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// setup a simple handler which sends a HTHS header for six months (!)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=15768000 ; includeSubDomains")
		fmt.Fprintf(w, "Hello, HTTPS world!")
	})

	//log.Printf("Serving http/https for domains: %+v", domains)
	go func() {
		// serve HTTP, which will redirect automatically to HTTPS
		h := jorm.CertManager.HTTPHandler(nil)
		log.Fatal(http.ListenAndServe(":4487", h))
	}()

	// serve HTTPS!
	log.Fatal(jorm.Server.ListenAndServeTLS("", ""))
}

//log.Fatal(srv.ListenAndServeTLS("./cfg/server.pem", "./cfg/server.key"))
//fmt.Println("Listening on port: ", cfg.C.Port)
//log.Fatal(jorm.Server.ListenAndServe())
// port := 9898
// fmt.Println("Listening on port:", port)
// log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(r)))

//http.HandleFunc("/", myHandler)
//server := &http.Server{
//	ReadTimeout:    10 * time.Second,
//	WriteTimeout:   10 * time.Second,
//	MaxHeaderBytes: 1 << 20,
//	TLSConfig:      tlsConfig,
//}

//listener, err := tls.Listen("tcp", ":8443", jorm.TLSconfig)
//if err != nil {
//	fmt.Println("err:", err)
//}
//log.Fatal(jorm.Server.Serve(listener))

//	go http.ListenAndServe(":http", jorm.CertManager.HTTPHandler(nil))
//
//
//	log.Fatal(jorm.Server.ListenAndServeTLS("", ""))
//}
