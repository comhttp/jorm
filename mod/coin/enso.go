package coin

import (
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/gorilla/mux"
)

func Queries(j *jdb.JDB, col string) *CoinsQueries {
	return &CoinsQueries{
		j,
		col,
	}
}

func ENSOroutes(cq *CoinsQueries, r *mux.Router) *mux.Router {
	//cq := j.CollectionQueries("coin").(CoinsQueries)
	//s := r.Host("enso.okno.rs").Subrouter()
	r.StrictSlash(true)

	//s.HandleFunc("/", h.HomeHandler)

	//f := s.PathPrefix("/f").Subrouter()
	//f.HandleFunc("/addcoin", h.AddCoinHandler).Methods("POST")
	//f.HandleFunc("/addnode", h.AddNodeHandler).Methods("POST")

	c := r.PathPrefix("/coins").Subrouter()
	c.HandleFunc("/", cq.CoinsHandler).Methods("GET")
	c.HandleFunc("/{coin}", cq.CoinHandler).Methods("GET")

	i := r.PathPrefix("/info").Subrouter()
	i.HandleFunc("/all", cq.allCoinsHandler).Methods("GET")
	i.HandleFunc("/node", cq.nodeCoinsHandler).Methods("GET")
	i.HandleFunc("/rest", cq.restCoinsHandler).Methods("GET")
	i.HandleFunc("/algo", cq.algoCoinsHandler).Methods("GET")
	i.HandleFunc("/words", cq.coinsWordsHandler).Methods("GET")
	i.HandleFunc("/usable", cq.usableCoinsHandler).Methods("GET")
	i.HandleFunc("/bin", cq.coinsBinHandler).Methods("GET")

	n := r.PathPrefix("/nodes").Subrouter()
	//n.HandleFunc("/{coin}/nodes", cq.CoinNodesHandler).Methods("GET")
	n.HandleFunc("/{coin}/{nodeip}", cq.nodeHandler).Methods("GET")

	return r
}
