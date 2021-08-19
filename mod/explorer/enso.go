package explorer

import (
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/gorilla/mux"
)

func Queries(j *jdb.JDBS, col string) *ExplorerQueries {
	return &ExplorerQueries{
		j:   j,
		col: col,
	}
}

func ENSOroutes(j *jdb.JDBS, r *mux.Router) *mux.Router {
	info := Queries(j, "info")
	r.StrictSlash(true)
	//n := s.PathPrefix("/n").Subrouter()
	//n.HandleFunc("/{coin}/nodes", explorersCollection.CoinNodesHandler).Methods("GET")
	//n.HandleFunc("/{coin}/{nodeip}", explorersCollection.nodeHandler).Methods("GET")

	b := r.PathPrefix("/explorer").Subrouter()
	b.HandleFunc("/{coin}/status", info.ViewStatus).Methods("GET")
	b.HandleFunc("/{coin}/blocks/{per}/{page}", info.ViewBlocks).Methods("GET")
	b.HandleFunc("/{coin}/lastblock", info.LastBlock).Methods("GET")
	b.HandleFunc("/{coin}/block/{id}", info.ViewBlock).Methods("GET")
	b.HandleFunc("/{coin}/tx/{txid}", info.ViewTx).Methods("GET")
	b.HandleFunc("/{coin}/mempool", info.ViewRawMemPool).Methods("GET")
	b.HandleFunc("/{coin}/mining", info.ViewMiningInfo).Methods("GET")
	b.HandleFunc("/{coin}/info", info.ViewInfo).Methods("GET")
	b.HandleFunc("/{coin}/peers", info.ViewPeers).Methods("GET")
	//b.HandleFunc("/{coin}/market", explorersCollection.ViewMarket).Methods("GET")
	return r
}
