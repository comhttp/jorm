package app

import (
	"github.com/gorilla/mux"
	"github.com/p9c/jorm/app/cfg"
	"github.com/p9c/jorm/app/tpl"
	"net/http"
	"strings"
	//"net/url"
	//"strings"
)

func (o *JORM) out(r *mux.Router) {
	////////////////
	// out
	////////////////
	a := r.Host("out.okno.rs").Subrouter()
	a.StrictSlash(true)

	a.HandleFunc("/", o.goodBye).Methods("GET")

	//a.Headers("Access-Control-Allow-Origin", "*")
}

func (o *JORM) goodBye(w http.ResponseWriter, r *http.Request) {
	url := strings.TrimSpace(r.URL.Query().Get("url"))
	img := strings.TrimSpace(r.URL.Query().Get("img"))
	title := strings.TrimSpace(r.URL.Query().Get("title"))
	out := map[string]string{
		"url":   url,
		"img":   img,
		"title": title,
	}
	tpl.TemplateHandler(cfg.Path+"/templates/okno").ExecuteTemplate(w, "out_gohtml", out)
}
