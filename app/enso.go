package app

import (
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

//type ENSO struct {
//	Coins    coins.Coins
//	WWW      *http.Server
//	JDB      *jdb.JDB
//	Explorer *explorers.Explorer
//}

//func NewENSO() *ENSO {
//	e := &ENSO{
//		JDB: jdb.NewJDB(cfg.C.JDBservers),
//	}
//	e.Explorer = explorers.GetExplorers(e.JDB)
//	e.WWW = &http.Server{
//		Handler:      handler(e),
//		Addr:         ":" + cfg.C.Port["enso"],
//		WriteTimeout: 15 * time.Second,
//		ReadTimeout:  15 * time.Second,
//	}
//	return e
//}

func (j *JORM) ENSOhandlers() http.Handler {
	s := mux.NewRouter()
	//s := r.Host("enso.okno.rs").Subrouter()
	s.StrictSlash(true)

	//s.HandleFunc("/", h.HomeHandler)

	//f := s.PathPrefix("/f").Subrouter()
	//f.HandleFunc("/addcoin", h.AddCoinHandler).Methods("POST")
	//f.HandleFunc("/addnode", h.AddNodeHandler).Methods("POST")

	c := s.PathPrefix("/coins").Subrouter()
	c.HandleFunc("/", j.CoinsHandler).Methods("GET")
	c.HandleFunc("/{coin}", j.CoinHandler).Methods("GET")

	i := s.PathPrefix("/info").Subrouter()
	i.HandleFunc("/all", j.allCoinsHandler).Methods("GET")
	i.HandleFunc("/node", j.nodeCoinsHandler).Methods("GET")
	i.HandleFunc("/rest", j.restCoinsHandler).Methods("GET")
	i.HandleFunc("/algo", j.algoCoinsHandler).Methods("GET")
	i.HandleFunc("/words", j.coinsWordsHandler).Methods("GET")
	i.HandleFunc("/usable", j.usableCoinsHandler).Methods("GET")
	i.HandleFunc("/bin", j.coinsBinHandler).Methods("GET")

	n := s.PathPrefix("/n").Subrouter()
	//n.HandleFunc("/{coin}/nodes", j.CoinNodesHandler).Methods("GET")
	n.HandleFunc("/{coin}/{nodeip}", j.nodeHandler).Methods("GET")

	b := s.PathPrefix("/e").Subrouter()
	b.HandleFunc("/{coin}/status", j.ViewStatus).Methods("GET")
	b.HandleFunc("/{coin}/blocks/{per}/{page}", j.ViewBlocks).Methods("GET")
	//b.HandleFunc("/{coin}/lastblock", hnd.LastBlock).Methods("GET")
	b.HandleFunc("/{coin}/block/{id}", j.ViewBlock).Methods("GET")
	b.HandleFunc("/{coin}/tx/{txid}", j.ViewTx).Methods("GET")
	//
	b.HandleFunc("/{coin}/mempool", j.ViewRawMemPool).Methods("GET")
	b.HandleFunc("/{coin}/mining", j.ViewMiningInfo).Methods("GET")
	b.HandleFunc("/{coin}/info", j.ViewInfo).Methods("GET")
	b.HandleFunc("/{coin}/peers", j.ViewPeers).Methods("GET")
	//b.HandleFunc("/{coin}/market", j.ViewMarket).Methods("GET")

	//j := s.PathPrefix("/j").Subrouter()

	//j.PathPrefix("/").Handler(e.ViewJSON())

	//j.Headers("Access-Control-Allow-Origin", "*")

	//f.HandleFunc("/{sec}/{coin}/{type}/{file}", j.ViewJSONfolder)
	//e.HandleFunc("/{sec}/{coin}/{app}/{type}/{file}", h.ViewJSONfolder)

	//a.HandleFunc("/", o.goodBye).Methods("GET")
	//f.Headers("Access-Control-Allow-Origin", "*")
	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(s, utl.DefaultErrorHandler)))
}
