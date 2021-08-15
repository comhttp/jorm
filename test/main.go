package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"time"
)

//import (
//	"github.com/rs/zerolog/log"
//	"net/http"
//	"net/http/httputil"
//	"net/url"
//)
//
//var (
//	hostTarget = map[string]string{
//		"okno.rs":                    "http://localhost:4433",
//		"parallelcoin.info":          "http://localhost:4433",
//		"explorer.parallelcoin.info": "http://localhost:4433",
//		"enso.okno.rs":               "http://localhost:14433",
//	}
//)
//
//type baseHandle struct{}
//
//func (h *baseHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	host := r.Host
//
//	if target, ok := hostTarget[host]; ok {
//		remoteUrl, err := url.Parse(target)
//		if err != nil {
//			log.Print("target parse fail:", err)
//			return
//		}
//
//		proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
//		proxy.ServeHTTP(w, r)
//		return
//	}
//	w.Write([]byte("403: Host forbidden " + host))
//}
//
//func main() {
//	h := &baseHandle{}
//	http.Handle("/", h)
//	server := &http.Server{
//		Addr:    ":80",
//		Handler: h,
//	}
//	log.Fatal(server.ListenAndServe())
//}

type JSONParams struct {
	Action string   `json:"action"`
	Subs   []string `json:"subs"`
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// this is where you paste your api key
	apiKey := "30054c457ea44e183bb614813c325675d53f4eef2151406b4745f8baaeeaa381"
	c, _, err := websocket.DefaultDialer.Dial("wss://streamer.cryptocompare.com/v2?api_key="+apiKey, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	jsonObj := JSONParams{Action: "SubAdd", Subs: []string{"2~*~*~*"}}
	s, _ := json.Marshal(jsonObj)
	fmt.Println(string(s))
	err = c.WriteMessage(websocket.TextMessage, []byte(string(s)))
	if err != nil {
		log.Fatal("message:", err)
	}

	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
