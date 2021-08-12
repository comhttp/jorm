package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	//"net/url"

	"github.com/sensepost/gowitness/chrome"
	"github.com/sensepost/gowitness/lib"
	"github.com/sensepost/gowitness/storage"
	//"strings"
)

var (
	options = lib.NewOptions()
	chrm    = chrome.NewChrome()
	db      = storage.NewDb()
)

func (o *JORM) img(r *mux.Router) {
	////////////////
	// img
	////////////////
	a := r.Host("img.okno.rs").Subrouter()
	a.StrictSlash(true)

	a.HandleFunc("/", o.viewWebImg).Methods("GET")

	//a.Headers("Access-Control-Allow-Origin", "*")
}

func (j *JORM) viewWebImg(w http.ResponseWriter, r *http.Request) {
	url := strings.TrimSpace(r.URL.Query().Get("url"))
	path := j.config.Path + "/static/img"
	_, err := os.Stat(path + "/" + url)
	if err != nil {
		log.Println(path + "/" + url)
		log.Println("nema")
		exec.Command("gowitness single --destination " + path + " -o " + url + ".png https://" + url)
	} else {
		log.Println(path)
		log.Println("ima")
	}
	http.ServeFile(w, r, path+"/"+url)
}
