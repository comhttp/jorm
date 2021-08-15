package coin

import (
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/gorilla/mux"
)

func Queries(j *jdb.JDB, col string) *CoinQueries {
	return &CoinQueries{
		j:   j,
		col: col,
	}
}

func ENSOroutes(j *jdb.JDBS, r *mux.Router) *mux.Router {
	coinsCollection := Queries(j.B["coins"], "coin")
	//s := r.Host("enso.okno.rs").Subrouter()
	r.StrictSlash(true)

	//s.HandleFunc("/", h.HomeHandler)

	//f := s.PathPrefix("/f").Subrouter()
	//f.HandleFunc("/addcoin", h.AddCoinHandler).Methods("POST")
	//f.HandleFunc("/addnode", h.AddNodeHandler).Methods("POST")

	c := r.PathPrefix("/coins").Subrouter()
	c.HandleFunc("/", coinsCollection.CoinsHandler).Methods("GET")
	c.HandleFunc("/{coin}", coinsCollection.CoinHandler).Methods("GET")

	i := r.PathPrefix("/info").Subrouter()
	i.HandleFunc("/all", coinsCollection.allCoinsHandler).Methods("GET")
	i.HandleFunc("/node", coinsCollection.nodeCoinsHandler).Methods("GET")
	i.HandleFunc("/rest", coinsCollection.restCoinsHandler).Methods("GET")
	i.HandleFunc("/algo", coinsCollection.algoCoinsHandler).Methods("GET")
	i.HandleFunc("/words", coinsCollection.coinsWordsHandler).Methods("GET")
	i.HandleFunc("/usable", coinsCollection.usableCoinsHandler).Methods("GET")
	i.HandleFunc("/bin", coinsCollection.coinsBinHandler).Methods("GET")

	n := r.PathPrefix("/n").Subrouter()
	//n.HandleFunc("/{coin}/nodes", coinsCollection.CoinNodesHandler).Methods("GET")
	n.HandleFunc("/{coin}/{nodeip}", coinsCollection.nodeHandler).Methods("GET")

	return r
}
