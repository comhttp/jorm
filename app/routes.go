package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (j *JORM) jorm(r *mux.Router) {
	////////////////
	// jorm
	////////////////
	// s := r.Host("jorm.okno.rs").Subrouter()
	r.StrictSlash(true)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("JORM!"))
	}) //s.HandleFunc("/", h.HomeHandler)
	//f := s.PathPrefix("/f").Subrouter()
	//f.HandleFunc("/addcoin", h.AddCoinHandler).Methods("POST")
}
