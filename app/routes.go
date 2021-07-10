package app

import (
	"github.com/gorilla/mux"
)

func (o *JORM) jorm(r *mux.Router) {
	////////////////
	// jorm
	////////////////
	s := r.Host("jorm.okno.rs").Subrouter()
	s.StrictSlash(true)

	//s.HandleFunc("/", h.HomeHandler)

	//f := s.PathPrefix("/f").Subrouter()
	//f.HandleFunc("/addcoin", h.AddCoinHandler).Methods("POST")
	//f.HandleFunc("/addnode", h.AddNodeHandler).Methods("POST")

	//a := s.PathPrefix("/a").Subrouter()
	//a.HandleFunc("/coins", h2.CoinsHandler).Methods("GET")
	//a.HandleFunc("/{coin}/nodes", h2.CoinNodesHandler).Methods("GET")
	//a.HandleFunc("/{coin}/{nodeip}", h2.NodeHandler).Methods("GET")
	//
	//b := s.PathPrefix("/b").Subrouter()
	//b.HandleFunc("/{coin}/blocks/{per}/{page}", h2.ViewBlocks).Methods("GET")
	//b.HandleFunc("/{coin}/lastblock", h2.LastBlock).Methods("GET")
	//b.HandleFunc("/{coin}/block/{id}", h2.ViewBlock).Methods("GET")
	//b.HandleFunc("/{coin}/tx/{txid}", h2.ViewTx).Methods("GET")
	//
	//b.HandleFunc("/{coin}/mempool", h2.ViewRawMemPool).Methods("GET")
	//b.HandleFunc("/{coin}/mining", h2.ViewMiningInfo).Methods("GET")
	//b.HandleFunc("/{coin}/info", h2.ViewInfo).Methods("GET")
	//b.HandleFunc("/{coin}/peers", h2.ViewPeers).Methods("GET")
	//b.HandleFunc("/{coin}/market", h2.ViewMarket).Methods("GET")
	//
	//j := s.PathPrefix("/j").Subrouter()

	//j.PathPrefix("/").Handler(h2.ViewJSON())

	//j.Headers("Access-Control-Allow-Origin", "*")

	//e := s.PathPrefix("/e").Subrouter()
	//e.HandleFunc("/{coin}/blocks/{per}/{page}", h.ViewBlocks).Methods("GET")
	//e.HandleFunc("/{coin}/lastblock", h.LastBlock).Methods("GET")
	//e.HandleFunc("/{sec}/{coin}/{type}/{file}", h2.ViewJSONfolder)
	//e.HandleFunc("/{sec}/{coin}/{app}/{type}/{file}", h.ViewJSONfolder)
	//e.HandleFunc("/{coin}/hash/{blockhash}", h.ViewHash).Methods("GET")
	//e.HandleFunc("/{coin}/tx/{txid}", h.ViewTx).Methods("GET")

	//a.HandleFunc("/", o.goodBye).Methods("GET")
	//e.Headers("Access-Control-Allow-Origin", "*")

	//filename = "/okno/sites/w/jdb/home"

}
