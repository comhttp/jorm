package app

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/rs/zerolog/log"
)

var (
	hostTarget = map[string]string{
		"okno.rs":                    "http://127.0.0.1:4433",
		"parallelcoin.info":          "http://127.0.0.1:4433",
		"explorer.parallelcoin.info": "http://127.0.0.1:4433",
		"jorm.okno.rs":               "http://127.0.0.1:14411",
		"our.okno.rs":                "http://127.0.0.1:14422",
		"enso.okno.rs":               "http://127.0.0.1:14433",
		"p9c.okno.rs":                "http://127.0.0.1:1337",
		"api.parallelcoin.io":        "http://127.0.0.1:11121",
		"admin.parallelcoin.io":      "http://127.0.0.1:11121",
		"api.parallelcoin.info":      "http://127.0.0.1:11121",
		"admin.parallelcoin.info":    "http://127.0.0.1:11121",
	}
)

type baseHandle struct{}

func (h *baseHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	log.Print("target hosthosthost fail:", host)

	if target, ok := hostTarget[host]; ok {
		reverseproxy(w, r, target)
	} else {
		reverseproxy(w, r, "http://localhost:14444")
	}
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")
}

func reverseproxy(w http.ResponseWriter, r *http.Request, target string) {
	remoteUrl, err := url.Parse(target)
	if err != nil {
		log.Print("target parse fail:", err)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
	w.Header().Set("AMP-Access-Control-Allow-Source-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin")
	
	w.Header().Set("Access-Control-Allow-Credentials","true")
	
	w.Header().Set("AMP-Same-Origin","true")

	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods","POST, GET, OPTIONS")

    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	
	proxy.ServeHTTP(w, r)
	return
}

//
//func (j *JORM) ReverseProxySRV() {
//	h := &BaseHandle{}
//	http.Handle("/", h)
//	server := &http.Server{
//		Addr:    ":80",
//		Handler: h,
//	}
//	log.Fatal(server.ListenAndServe())
//}
func status(w http.ResponseWriter, r *http.Request) {
	// Handles top-level page.
	fmt.Fprintf(w, "You are on the status home page")
}
