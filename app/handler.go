package app

import (
	"github.com/comhttp/jorm/app/mod"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type PageData struct {
	Title string
	//Host     Host
	Site     string
	HostSlug string
	Post     mod.Post
	Slug     string
	Section  string
	Template string
	Posts    []mod.Post
	//Hosts    []Host
}

func (o *JORM) Handler() http.Handler {
	r := mux.NewRouter()

	o.img(r)
	//o.out(r)
	o.jorm(r)

	return handlers.CORS()(handlers.CompressHandler(InterceptHandler(r, DefaultErrorHandler)))
}
