package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	// setup a simple handler which sends a HTHS header for six months (!)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=15768000 ; includeSubDomains")
		fmt.Fprintf(w, "Hello, HTTPS world!")
	})

	// look for the domains to be served from command line args
	flag.Parse()
	domains := flag.Args()
	if len(domains) == 0 {
		log.Fatalf("fatal; specify domains as arguments")
	}

	// create the autocert.Manager with domains and path to the cache
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domains...),
	}

	// optionally use a cache dir
	dir := cacheDir()
	if dir != "" {
		certManager.Cache = autocert.DirCache(dir)
	}

	// create the server itself
	server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	log.Printf("Serving http/https for domains: %+v", domains)
	go func() {
		// serve HTTP, which will redirect automatically to HTTPS
		h := certManager.HTTPHandler(nil)
		log.Fatal(http.ListenAndServe(":http", h))
	}()

	// serve HTTPS!
	log.Fatal(server.ListenAndServeTLS("", ""))
}

// cacheDir makes a consistent cache directory inside /tmp. Returns "" on error.
func cacheDir() (dir string) {
	if u, _ := user.Current(); u != nil {
		dir = filepath.Join(os.TempDir(), "cache-golang-autocert-"+u.Username)
		if err := os.MkdirAll(dir, 0700); err == nil {
			return dir
		}
	}
	return ""
}

//// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
//// Use of this source code is governed by a BSD-style
//// license that can be found in the LICENSE file.
//
//package main
//
//import (
//	"flag"
//
//	"github.com/comhttp/jorm/app"
//	"github.com/gorilla/handlers"
//	"github.com/gorilla/mux"
//	"html/template"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"os"
//	"strconv"
//	"time"
//
//	"github.com/gorilla/websocket"
//)
//
//const (
//	// Time allowed to write the file to the client.
//	writeWait = 10 * time.Second
//
//	// Time allowed to read the next pong message from the client.
//	pongWait = 60 * time.Second
//
//	// Send pings to client with this period. Must be less than pongWait.
//	pingPeriod = (pongWait * 9) / 10
//
//	// Poll file for changes with this period.
//	filePeriod = 10 * time.Second
//)
//
//var (
//	addr      = flag.String("addr", ":4488", "http service address")
//	homeTempl = template.Must(template.New("").Parse(homeHTML))
//	filename  string
//	upgrader  = websocket.Upgrader{
//		ReadBufferSize:  1024,
//		WriteBufferSize: 1024,
//	}
//)
//
//func readFileIfModified(lastMod time.Time) ([]byte, time.Time, error) {
//	fi, err := os.Stat(filename)
//	if err != nil {
//		return nil, lastMod, err
//	}
//	if !fi.ModTime().After(lastMod) {
//		return nil, lastMod, nil
//	}
//	p, err := ioutil.ReadFile(filename)
//	if err != nil {
//		return nil, fi.ModTime(), err
//	}
//	return p, fi.ModTime(), nil
//}
//
//func reader(ws *websocket.Conn) {
//	defer ws.Close()
//	ws.SetReadLimit(512)
//	ws.SetReadDeadline(time.Now().Add(pongWait))
//	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
//	for {
//		_, _, err := ws.ReadMessage()
//		if err != nil {
//			break
//		}
//	}
//}
//
//func writer(ws *websocket.Conn, lastMod time.Time) {
//	lastError := ""
//	pingTicker := time.NewTicker(pingPeriod)
//	fileTicker := time.NewTicker(filePeriod)
//	defer func() {
//		pingTicker.Stop()
//		fileTicker.Stop()
//		ws.Close()
//	}()
//	for {
//		select {
//		case <-fileTicker.C:
//			var p []byte
//			var err error
//
//			p, lastMod, err = readFileIfModified(lastMod)
//
//			if err != nil {
//				if s := err.Error(); s != lastError {
//					lastError = s
//					p = []byte(lastError)
//				}
//			} else {
//				lastError = ""
//			}
//
//			if p != nil {
//				ws.SetWriteDeadline(time.Now().Add(writeWait))
//				if err := ws.WriteMessage(websocket.TextMessage, p); err != nil {
//					return
//				}
//			}
//		case <-pingTicker.C:
//			ws.SetWriteDeadline(time.Now().Add(writeWait))
//			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
//				return
//			}
//		}
//	}
//}
//
//func serveWs(w http.ResponseWriter, r *http.Request) {
//	ws, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		if _, ok := err.(websocket.HandshakeError); !ok {
//			log.Println(err)
//		}
//		return
//	}
//
//	var lastMod time.Time
//	if n, err := strconv.ParseInt(r.FormValue("lastMod"), 16, 64); err == nil {
//		lastMod = time.Unix(0, n)
//	}
//
//	go writer(ws, lastMod)
//	reader(ws)
//}
//
//func serveHome(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/" {
//		http.Error(w, "Not found", http.StatusNotFound)
//		return
//	}
//	if r.Method != "GET" {
//		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//		return
//	}
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//	p, lastMod, err := readFileIfModified(time.Time{})
//	if err != nil {
//		p = []byte(err.Error())
//		lastMod = time.Unix(0, 0)
//	}
//	var v = struct {
//		Host    string
//		Data    string
//		LastMod string
//	}{
//		r.Host,
//		string(p),
//		strconv.FormatInt(lastMod.UnixNano(), 16),
//	}
//	homeTempl.Execute(w, &v)
//}
//func aHandler() http.Handler {
//	r := mux.NewRouter()
//	s := r.Host("ws.okno.rs").Subrouter()
//
//	//w := s.PathPrefix("/w").Subrouter()
//	s.StrictSlash(true)
//	s.HandleFunc("/", serveHome)
//	s.HandleFunc("/ws", serveWs)
//
//	return handlers.CORS()(handlers.CompressHandler(app.InterceptHandler(r, app.DefaultErrorHandler)))
//}
//
//func main() {
//	filename = "/jorm/out/exchanges/poloniex"
//
//	srv := &http.Server{
//		Handler:      aHandler(),
//		Addr:         ":4488",
//		WriteTimeout: 15 * time.Second,
//		ReadTimeout:  15 * time.Second,
//	}
//
//	log.Fatal(srv.ListenAndServe())
//}
//
//const homeHTML = `<!DOCTYPE html>
//<html lang="en">
//    <head>
//        <title>WebSocket Example</title>
//    </head>
//    <body>
//        <pre id="fileData">{{.Data}}</pre>
//        <script type="text/javascript">
//            (function() {
//                var data = document.getElementById("fileData");
//                var conn = new WebSocket("ws://{{.Host}}/ws?lastMod={{.LastMod}}");
//                conn.onclose = function(evt) {
//                    data.textContent = 'Connection closed';
//                }
//                conn.onmessage = function(evt) {
//                    console.log('file updated');
//                    data.textContent = evt.data;
//                }
//            })();
//        </script>
//    </body>
//</html>
//`
