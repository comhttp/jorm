package exchange

import (
	"encoding/json"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

func ENSOroutes(eq *ExchangeQueries, r *mux.Router) *mux.Router {
	//cq := j.CollectionQueries("coin").(CoinQueries)
	//s := r.Host("enso.okno.rs").Subrouter()
	r.StrictSlash(true)

	//s.HandleFunc("/", h.HomeHandler)

	//f := s.PathPrefix("/f").Subrouter()
	//f.HandleFunc("/addcoin", h.AddCoinHandler).Methods("POST")
	//f.HandleFunc("/addnode", h.AddNodeHandler).Methods("POST")

	e := r.PathPrefix("/exchanges").Subrouter()
	e.HandleFunc("/", eq.ExchangesHandler).Methods("GET")
	e.HandleFunc("/{exchange}", eq.ExchangeHandler).Methods("GET")
	e.HandleFunc("/{exchange}/markets", eq.MarketsHandler).Methods("GET")
	e.HandleFunc("/{exchange}/markets/{market}", eq.MarketHandler).Methods("GET")

	return r
}

// CoinsHandler handles a request for coin data
func (eq *ExchangeQueries) ExchangeHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	exchange, err := eq.GetExchange(v["exchange"])
	out, err := json.Marshal(exchange)
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (eq *ExchangeQueries) ExchangesHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(eq.GetExchanges())
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (eq *ExchangeQueries) MarketHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	market, err := eq.GetExchangeMarket(v["exchange"], strings.ToLower(v["market"]))
	utl.ErrorLog(err)
	out, err := json.Marshal(market)
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (eq *ExchangeQueries) MarketsHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	//out, err := json.Marshal(coin.LoadCoinsBase())
	markets, err := eq.GetExchangeMarkets(v["exchange"])
	utl.ErrorLog(err)
	out, err := json.Marshal(markets)
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
