package app

import (
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type PageData struct {
	Title string
	Site     string
	HostSlug string
	Slug     string
	Section  string
	Template string
}

func (o *JORM) WWWhandleR() http.Handler {
	r := mux.NewRouter()

	//o.img(r)
	//o.out(r)
	o.jorm(r)

	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))
}

func (o *JORM) WShandleR() http.Handler {
	r := mux.NewRouter()
	r.Host("ws.okno.rs").Subrouter()
	r.StrictSlash(true)
	r.HandleFunc("/", serveHome)
	r.HandleFunc("/ws", serveWs)
	r.Headers("Access-Control-Allow-Origin", "*")

	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))
}
