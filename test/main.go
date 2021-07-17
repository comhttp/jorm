package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	hostTarget = map[string]string{
		"okno.rs":                    "http://localhost:4433",
		"parallelcoin.info":          "http://localhost:4433",
		"explorer.parallelcoin.info": "http://localhost:4433",
		"enso.okno.rs":               "http://localhost:14433",
	}
)

type baseHandle struct{}

func (h *baseHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host

	if target, ok := hostTarget[host]; ok {
		remoteUrl, err := url.Parse(target)
		if err != nil {
			log.Println("target parse fail:", err)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
		proxy.ServeHTTP(w, r)
		return
	}
	w.Write([]byte("403: Host forbidden " + host))
}

func main() {
	h := &baseHandle{}
	http.Handle("/", h)
	server := &http.Server{
		Addr:    ":80",
		Handler: h,
	}
	log.Fatal(server.ListenAndServe())
}
