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
}
