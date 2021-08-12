package app

import (
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func (j *JORM) ADMINhandlers() http.Handler {
	s := mux.NewRouter()
	//s := r.Host("enso.okno.rs").Subrouter()
	s.StrictSlash(true)

	//s.HandleFunc("/", h.HomeHandler)

	//f := s.PathPrefix("/f").Subrouter()
	//f.HandleFunc("/addcoin", h.AddCoinHandler).Methods("POST")
	//f.HandleFunc("/addnode", h.AddNodeHandler).Methods("POST")

	c := s.PathPrefix("/coins").Subrouter()
	c.HandleFunc("/", j.ADMINCoinsHandler).Methods("GET")
	c.HandleFunc("/{coin}", j.ADMINCoinHandler).Methods("GET")

	//a.HandleFunc("/", o.goodBye).Methods("GET")
	//f.Headers("Access-Control-Allow-Origin", "*")
	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(s, utl.DefaultErrorHandler)))
}
