package coin

import "C"
import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/comhttp/jorm/pkg/utl/img"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type home struct {
	D []Coin
	C Coins
}

// HomeHandler handles a request for (?)
//func HomeHandler(w http.ResponseWriter, r *http.Request) {
//	coins := coin.ReadAllCoins()
//	var bitnoded []coin.Coin
//	for _, coin := range coins.C {
//		if utl.FileExists(cfg.Path + "/jorm/" + coin.Slug + "/info/bitnodes") {
//			bitnoded = append(bitnoded, coin)
//		}
//		//log.Print("coin", coin)
//	}
//	data := home{
//		D: bitnoded,
//		C: coins,
//	}
//	tpl.TemplateHandler(cfg.Path+"/templates/comhttp").ExecuteTemplate(w, "base_gohtml", data)
//}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) CoinHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	coin, err := cq.GetCoin(v["coin"])
	out, err := json.Marshal(coin)
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) CoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(cq.GetCoins())
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) restCoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(cq.GetRestCoins())
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) algoCoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(cq.GetAlgoCoins())
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) coinsWordsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(cq.GetCoinsWords())
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) usableCoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(cq.GetUsableCoins())
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) allCoinsHandler(w http.ResponseWriter, r *http.Request) {
	//out, err := json.Marshal(coin.LoadCoinsBase())
	out, err := json.Marshal(cq.GetAllCoins())
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) nodeCoinsHandler(w http.ResponseWriter, r *http.Request) {
	out, err := json.Marshal(cq.GetNodeCoins())
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinsHandler handles a request for coin data
func (cq *CoinsQueries) coinsBinHandler(w http.ResponseWriter, r *http.Request) {
	out, err := json.Marshal(cq.GetCoinsBin())
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

// CoinNodesHandler handles a request for (?)
func (cq *CoinsQueries) CoinNodesHandler(w http.ResponseWriter, r *http.Request) {

}

// NodeHandler handles a request for (?)
func (cq *CoinsQueries) nodeHandler(w http.ResponseWriter, r *http.Request) {
	//v := mux.Vars(r)
	//out, err := json.Marshal(nodes.GetNode(j.JDB, v["coin"], v["nodeip"]))
	//if err != nil {
	//	log.Print("Error encoding JSON")
	//	return
	//}
	//w.Write([]byte(out))
}

//// NodeHandler handles a request for (?)
//func (cq *CoinsQueries) ViewJSON() http.Handler {
//	m := minify.New()
//	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), mjson.Minify)
//
//	return http.StripPrefix("/j", m.Middleware(http.FileServer(http.Dir(j.config.Path))))
//}

// NodeHandler handles a request for (?)
func (cq *CoinsQueries) ViewJSONfolder(w http.ResponseWriter, r *http.Request) {
	//v := mux.Vars(r)
	//m := minify.New()
	//height, err := strconv.ParseUint(v["file"], 10, 64)
	//if err != nil {
	//m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), mjson.Minify)
	//path := v["sec"] + "/" + v["coin"] + "/" + v["type"]
	//http.StripPrefix("/e/"+path, m.Middleware(http.FileServer(http.Dir(j.config.Path+"/www/"+path)))).ServeHTTP(w, r)
	//} else {
	//	index := map[uint64]string{}
	//if err := jdb.JDB.Read("/www/data/"+v["coin"]+"/index", v["type"], &index); err != nil {
	//	log.Print("Error", err)
	//}

	//log.Print("index[height]", index[height])
	//out := map[string]interface{}{}
	//if err := jdb.JDB.Read(cfg.C.Out+"/"+v["coin"]+"/"+v["type"], index[height], &out); err != nil {
	//	log.Print("Error", err)
	//}
	//m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), mjson.Minify)
	//json.NewEncoder(w).Encode(out)
	json.NewEncoder(w).Encode("out")
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//}
}

// LogoHandler handles a request for logo data
func (cq *CoinsQueries) LogoHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	size, err := strconv.ParseFloat(v["size"], 32)
	log.Print("Error encoding JSON: ", err)
	w.Write(cq.getLogo(v["coin"], size))
}

// LogoHandler handles a request for logo data
func (cq *CoinsQueries) getLogo(coin string, size float64) []byte {
	logoRawString, err := cq.GetLogo(coin)
	if err != nil {
		log.Print("Error encoding JSON")
	}
	logoRaw, err := hex.DecodeString(logoRawString)
	logo, _ := img.ImageResize(logoRaw, img.Options{Width: size, Height: size})
	return logo
}

// jsonHandler handles a request for json data
func (cq *CoinsQueries) jsonAlgoCoinsHandler(w http.ResponseWriter, r *http.Request) {
	algoCoins := cq.GetAlgoCoins()
	var algoCoinsLogo []CoinShortLogo
	for _, ac := range algoCoins.C {
		algoCoinsLogo = append(algoCoinsLogo, CoinShortLogo{
			Rank:   ac.Rank,
			Name:   ac.Name,
			Symbol: ac.Symbol,
			Slug:   ac.Slug,
			Algo:   ac.Algo,
			Logo:   base64.StdEncoding.EncodeToString(cq.getLogo(ac.Slug, 64)),
		})
	}
	// out, err := json.Marshal(algoCoinsLogo)
	// if err != nil {
	// 	log.Print("Error encoding JSON: ", err)
	// }
	//v := mux.Vars(r)
	//m := minify.New()
	//height, err := strconv.ParseUint(v["file"], 10, 64)
	//if err != nil {
	//m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), mjson.Minify)
	//path := v["sec"] + "/" + v["coin"] + "/" + v["type"]
	//http.StripPrefix("/e/"+path, m.Middleware(http.FileServer(http.Dir(j.config.Path+"/www/"+path)))).ServeHTTP(w, r)
	//} else {
	//	index := map[uint64]string{}
	//if err := jdb.JDB.Read("/www/data/"+v["coin"]+"/index", v["type"], &index); err != nil {
	//	log.Print("Error", err)
	//}

	//log.Print("index[height]", index[height])
	//out := map[string]interface{}{}
	//if err := jdb.JDB.Read(cfg.C.Out+"/"+v["coin"]+"/"+v["type"], index[height], &out); err != nil {
	//	log.Print("Error", err)
	//}
	//m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), mjson.Minify)
	//json.NewEncoder(w).Encode(out)
	json.NewEncoder(w).Encode(algoCoinsLogo)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//}
}
