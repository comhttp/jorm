package app

import (
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"text/template"
)

var (
	funcMap = template.FuncMap{
		"truncate": utl.Truncate,
	}
)

type Data struct {
	Our, Base, TLD, Slug, Coin, Bg, App, Section, URL, Page, ID, Title, Path, ProtoURL, Canonical string
}

func newData(our string) *Data {
	return &Data{
		Our: our,
	}
}

func (d *Data) base(base string) {
	d.Base = base
	return
}
func (d *Data) tld(tld string) {
	d.TLD = tld
	return
}
func (d *Data) slug(slug string) {
	d.Slug = slug
	return
}
func (d *Data) app(app string) {
	d.App = app
	return
}
func (d *Data) section(section string) {
	d.Section = section
	return
}

func (j *JORM) COMHTTPhandlers() http.Handler {
	r := mux.NewRouter()
	tld := r.Host("com-http.{tld}").Subrouter()
	tld.HandleFunc("/", j.appHandler())
	tld.HandleFunc("/{app}", j.appHandler())
	tld.HandleFunc("/{app}/{page}", j.appHandler())
	tld.Headers("Access-Control-Allow-Origin", "*")
	tld.StrictSlash(true)
	sub := r.Host("{slug}.com-http.{tld}").Subrouter()
	sub.HandleFunc("/", j.appHandler())
	sub.HandleFunc("/{app}", j.appHandler())
	sub.HandleFunc("/{app}/{page}", j.appHandler())
	sub.StrictSlash(true)
	sub.Headers("Access-Control-Allow-Origin", "*")
	sub.Headers("Content-Type", "application/json")
	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))
}

func (j *JORM) appHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		d := newData("com-http")
		d.TLD = mux.Vars(r)["tld"]
		d.Slug = mux.Vars(r)["slug"]
		d.App = mux.Vars(r)["app"]
		d.Page = mux.Vars(r)["page"]
		d.Base = "amp"
		d.Bg = "parallelcoin"
		d.Path = "rts/tld/" + d.TLD
		d.ProtoURL = "https://" + d.Our + "."
		d.Title = "Beyond blockchain - " + d.TLD
		if d.Page != "" {
			d.ID = d.Page
			d.Page = d.App
			d.Title = d.Slug + "-" + d.App + "-Beyond blockchain-" + d.TLD + "-" + d.Page
		} else {
			d.Page = "index"
		}

		if d.App != "" {
			d.Path = "rts/coin/" + d.TLD
			d.Title = d.Slug + " - Beyond blockchain - " + d.TLD
			d.ProtoURL = "https://" + d.Slug + "." + d.Our + "."
			d.Canonical = d.ProtoURL + d.TLD

		} else {
			d.App = d.TLD
		}

		if d.Slug != "" {
			d.Bg = d.Slug
			d.Path = "rts/coin/" + d.TLD
			d.Title = d.Slug + " - Beyond blockchain - " + d.TLD
			d.ProtoURL = "https://" + d.Slug + "." + d.Our + "."
			d.Canonical = d.ProtoURL + d.TLD
		} else {
			d.Section = "coin"
			d.Path = "rts/tld/" + d.TLD
		}

		if d.TLD == "net" {
			d.Base = "vue"
			d.Path = "vue"
			d.Page = "index"
			d.Title = "Beyond blockchain - " + d.TLD
			if d.Slug != "" {

			} else {
				d.Section = "coin"
			}
		}
		funcMap := template.FuncMap{
			"truncate": utl.Truncate,
		}
		log.Println("Top level domain keyword 1:  ", d.TLD)
		log.Println("App 1: ", d.App)
		log.Println("Slug 1: ", d.Slug)
		log.Println("Page 1: ", d.Page)
		log.Println("Path 1: ", d.Path)
		log.Println("Base 1: ", d.Base)
		template.Must(j.parseFiles(d.Base, d.Path+"/"+d.Page+".gohtml")).Funcs(funcMap).ExecuteTemplate(w, d.Base, d)
		log.Println("d.Base", d.Base)
		log.Println("dadadad", d.Path+"/"+d.Page+".gohtml")
	}
}
