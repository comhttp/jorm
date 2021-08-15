package explorer

import (
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/gorilla/mux"
)

func Queries(j *jdb.JDB, col string) *ExplorersQueries {
	return &ExplorersQueries{
		j:   j,
		col: col,
	}
}

func ENSOroutes(j *jdb.JDBS, r *mux.Router) *mux.Router {
	explorersCollection := Queries(j.B["parallelcoin"], "coin")
	r.StrictSlash(true)
	//n := s.PathPrefix("/n").Subrouter()
	//n.HandleFunc("/{coin}/nodes", explorersCollection.CoinNodesHandler).Methods("GET")
	//n.HandleFunc("/{coin}/{nodeip}", explorersCollection.nodeHandler).Methods("GET")

	b := r.PathPrefix("/e").Subrouter()
	b.HandleFunc("/{coin}/status", explorersCollection.ViewStatus).Methods("GET")
	b.HandleFunc("/{coin}/blocks/{per}/{page}", explorersCollection.ViewBlocks).Methods("GET")
	b.HandleFunc("/{coin}/lastblock", explorersCollection.LastBlock).Methods("GET")
	b.HandleFunc("/{coin}/block/{id}", explorersCollection.ViewBlock).Methods("GET")
	b.HandleFunc("/{coin}/tx/{txid}", explorersCollection.ViewTx).Methods("GET")
	b.HandleFunc("/{coin}/mempool", explorersCollection.ViewRawMemPool).Methods("GET")
	b.HandleFunc("/{coin}/mining", explorersCollection.ViewMiningInfo).Methods("GET")
	b.HandleFunc("/{coin}/info", explorersCollection.ViewInfo).Methods("GET")
	b.HandleFunc("/{coin}/peers", explorersCollection.ViewPeers).Methods("GET")
	//b.HandleFunc("/{coin}/market", explorersCollection.ViewMarket).Methods("GET")
	return r
}
