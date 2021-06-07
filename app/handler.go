package app

import (
	"github.com/comhttp/jorm/app/cfg"
	"github.com/comhttp/jorm/app/mod"
	"github.com/comhttp/jorm/app/tpl"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"text/template"
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

	//for _, h := range o.Hosts {
	//	dh := h.domain(r)
	//	o.index(dh, h)
	//	//o.section(dh, h)
	//	o.post(dh, h)
	//	//o.staticHost(dh, h.Slug)
	//	if h.Slug != "okno_rs" && h.Slug != "marcetin_com" {
	//		sh := h.sub(r)
	//		o.index(sh, h)
	//		//o.section(sh, h)
	//		o.post(sh, h)
	//		//o.staticHost(sh, h.Slug)
	//		//if h.Template != "parallelcoin" {
	//		//	o.chat(sh)
	//		//}
	//	}
	//}

	o.img(r)
	o.out(r)
	o.jorm(r)

	return handlers.CORS()(handlers.CompressHandler(InterceptHandler(r, DefaultErrorHandler)))
}

// HomeHandler handles a request for (?)
//func (o *JORM) index(rt *mux.Router, host Host) {
//	rt.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		site := mux.Vars(r)["site"]
//		title := host.Name
//		hostSlug := host.Slug
//		section := "index"
//		template := site
//		if site != "" {
//			hostSlug = site + "_" + host.Slug
//			section = site
//		} else {
//			site = hostSlug
//			template = "index"
//		}
//		data := &PageData{
//			Title:    title,
//			Host:     host,
//			HostSlug: hostSlug,
//			Site:     site,
//			Section:  section,
//			Template: template,
//			}
//		o.template(w, host, data)
//	})
//}
//
//// HomeHandler handles a request for (?)
//func (o *JORM) section(rt *mux.Router, host Host) {
//	rt.HandleFunc("/{section}", func(w http.ResponseWriter, r *http.Request) {
//		site := mux.Vars(r)["site"]
//		section := mux.Vars(r)["section"]
//		hostSlug := host.Slug
//		template := site + "/" + section
//		if site != "" {
//			hostSlug = site + "_" + host.Slug
//		}
//		data := &PageData{
//			Title:   host.Name,
//			Host:    host,
//			Site:    hostSlug,
//			Section: section,
//			Template: template,
//			//Posts: o.posts("/sites/" + hostSlug + "/jdb/" + section),
//		}
//		o.template(w, host, data)
//	})
//}
//
//// HomeHandler handles a request for (?)
//func (o *JORM) post(rt *mux.Router, host Host) {
//	rt.HandleFunc("/{section}/{slug}", func(w http.ResponseWriter, r *http.Request) {
//		id := mux.Vars(r)["slug"]
//		site := mux.Vars(r)["site"]
//		section := mux.Vars(r)["section"]
//		template := site
//		hostSlug := host.Slug
//		if site != "" {
//			hostSlug = site + "_" + host.Slug
//			template = site + "/" + section
//		}
//		//post := o.Database.ReadPost("/sites/"+ hostSlug +"/jdb", section, id)
//		data := &PageData{
//			Title:   host.Name,
//			Host:     host,
//			HostSlug: hostSlug,
//			Site:     site,
//			//Post: post,
//			Template: template,
//			Slug:    id,
//			Section: section,
//		}
//		o.template(w, host, data)
//	})
//}
//
func (o *JORM) template(w http.ResponseWriter, theme string, data interface{}) {
	funcMap := template.FuncMap{
		"truncate": utl.Truncate,
		"sha384":   utl.SHA384,
	}
	templatePath := cfg.Path + "/tpl/gohtml/index.gohtml"
	if theme != "" {
		templatePath = cfg.Path + "/tpl/gohtml/index.gohtml"
	}
	_, err := os.Stat(templatePath)
	if err != nil {
		tpl.TemplateHandler(cfg.Path).Funcs(funcMap).ExecuteTemplate(w, "err_gohtml", err)
	} else {
		if theme != "" {
			tpl.TemplateHandler(cfg.Path).Funcs(funcMap).ExecuteTemplate(w, "index_gohtml", data)
		} else {
			tpl.TemplateHandler(cfg.Path).Funcs(funcMap).ExecuteTemplate(w, "index_gohtml", data)
		}
	}
}
