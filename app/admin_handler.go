package app

import "C"
import (
	"net/http"
)

// CoinsHandler handles a request for coin data
func (j *JORM) ADMINCoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.GetCoins(j.JDB))
	//if err != nil {
	//	log.Print("Error encoding JSON")
	//	return
	//}
	//w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (j *JORM) ADMINCoinHandler(w http.ResponseWriter, r *http.Request) {
	//v := mux.Vars(r)
	//out, err := json.Marshal(coin.GetCoin(j.JDB, v["coin"]))
	//if err != nil {
	//	log.Print("Error encoding JSON")
	//	return
	//}
	//w.Write([]byte(out))
}
