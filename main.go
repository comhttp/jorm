package main

import (
	"fmt"
	"github.com/comhttp/jorm/app"
	"github.com/comhttp/jorm/app/cfg"
	"github.com/comhttp/jorm/app/jorm/coin"
	//"github.com/comhttp/jorm/app/jorm/exchange"
	"log"
	"time"
)

func main() {
	jorm := app.NewJORM()
	coins := coin.LoadCoinsBase()
	//_ = exchange.ReadAllExchanges()
	ticker := time.NewTicker(999 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				app.Tickers(coins)
				fmt.Println("OKNO wooikos")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	//log.Fatal(srv.ListenAndServeTLS("./cfg/server.pem", "./cfg/server.key"))
	fmt.Println("Listening on port: ", cfg.CONFIG.Port)
	log.Fatal(jorm.Server.ListenAndServe())
	// port := 9898
	// fmt.Println("Listening on port:", port)
	// log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(r)))
}
