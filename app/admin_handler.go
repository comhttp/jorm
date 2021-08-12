package app

import "C"
import (
	"encoding/json"
	"github.com/comhttp/jorm/mod/coins"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// CoinsHandler handles a request for coin data
func (j *JORM) ADMINCoinsHandler(w http.ResponseWriter, r *http.Request) {
	out, err := json.Marshal(coins.GetCoins(j.JDB))
	if err != nil {
		log.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (j *JORM) ADMINCoinHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	out, err := json.Marshal(coins.GetCoin(j.JDB, v["coin"]))
	if err != nil {
		log.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
