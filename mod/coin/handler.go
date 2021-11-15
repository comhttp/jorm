package coin

import "C"
import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/comhttp/jorm/pkg/utl/img"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type home struct {
	D []Coin
	C Coins
}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) CoinHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	coin, err := cq.GetCoin(v["coin"])
	out, err := json.Marshal(coin)
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) CoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(cq.GetCoins())
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) restCoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(cq.GetRestCoins())
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) algoCoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(cq.GetAlgoCoins())
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) coinsWordsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(cq.GetCoinsWords())
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) usableCoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(cq.GetUsableCoins())
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) allCoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(cq.GetAllCoins())
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) nodeCoinsHandler(w http.ResponseWriter, r *http.Request) {
	out, err := json.Marshal(cq.GetNodeCoins())
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) coinsBinHandler(w http.ResponseWriter, r *http.Request) {
	out, err := json.Marshal(cq.GetCoinsBin())
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinNodesHandler handles a request for (?)
func (cq *CoinsQueries) CoinNodesHandler(w http.ResponseWriter, r *http.Request) {

}

// NodeHandler handles a request for (?)
func (cq *CoinsQueries) nodeHandler(w http.ResponseWriter, r *http.Request) {
	//v := mux.Vars(r)
	//out, err := json.Marshal(nodes.GetNode(j.JDB, v["coin"], v["nodeip"]))
	//if err != nil {
	//	log.Print("Error encoding JSON")
	//	return
	//}
	//w.Write([]byte(out))
}

//// NodeHandler handles a request for (?)
//func (cq *CoinsQueries) ViewJSON() http.Handler {
//	m := minify.New()
//	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), mjson.Minify)
//
//	return http.StripPrefix("/j", m.Middleware(http.FileServer(http.Dir(j.config.Path))))
//}

// LogoHandler handles a request for logo data
func (cq *CoinsQueries) LogoHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	size, err := strconv.ParseFloat(v["size"], 32)
	log.Print("Error encoding JSON: ", err)
	w.Write(cq.getLogo(v["coin"], size))
}

// LogoHandler handles a request for logo data
func (cq *CoinsQueries) getLogo(coin string, size float64) []byte {
	logoRawString, err := cq.GetLogo(coin)
	if err != nil {
		log.Print("Error encoding JSON")
	}
	logoRaw, err := hex.DecodeString(logoRawString)
	logo, _ := img.ImageResize(logoRaw, img.Options{Width: size, Height: size})
	return logo
}

// jsonHandler handles a request for json data
func (cq *CoinsQueries) jsonAlgoCoinsHandler(w http.ResponseWriter, r *http.Request) {
	algoCoinsLogo := cq.GetAlgoCoinsLogo()
	out, err := json.Marshal(algoCoinsLogo)
	if err != nil {
		log.Print("Error encoding JSON", err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(out))
}
