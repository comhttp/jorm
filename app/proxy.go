package app

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	hostTarget = map[string]string{
		"jorm.okno.rs": "http://localhost:14411",
		"our.okno.rs":  "http://localhost:14422",
		"enso.okno.rs": "http://localhost:14433",
	}
)

type BaseHandle struct{}

func (h *BaseHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	log.Println("target hosthosthost fail:", host)

	if target, ok := hostTarget[host]; ok {
		reverseproxy(w, r, target)
	} else {
		reverseproxy(w, r, "http://localhost:14444")
	}
	w.Write([]byte("403: Host forbidden " + host))
}

func reverseproxy(w http.ResponseWriter, r *http.Request, target string) {
	remoteUrl, err := url.Parse(target)
	if err != nil {
		log.Println("target parse fail:", err)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
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
