package explorer

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func (eq *ExplorerQueries) ViewStatus(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	out, err := json.Marshal(eq.GetExplorer(v["coin"]))
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (eq *ExplorerQueries) ViewBlocks(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	per, _ := strconv.Atoi(v["per"])
	page, _ := strconv.Atoi(v["page"])
	//ex := eq.GetExplorer(v["coin"])
	lastblock := eq.GetLastBlock(v["coin"])
	log.Print("lastblocklastblocklastblock", lastblock)

	lb := map[string]interface{}{
		"currentPage": page,
		"pageCount":   lastblock / per,
		"blocks":      eq.GetBlocks(v["coin"], per, page),
		"lastBlock":   lastblock,
	}

	out, err := json.Marshal(lb)
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (eq *ExplorerQueries) LastBlock(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	out, err := json.Marshal(eq.GetLastBlock(v["coin"]))
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (eq *ExplorerQueries) ViewBlock(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	out, err := json.Marshal(eq.GetBlock(v["coin"], v["id"]))
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

//func ViewBlockHeight(w http.ResponseWriter, r *http.Request) {
//	v := mux.Vars(r)
//	bh := v["blockheight"]
//	// node := Node{}
//	bhi, _ := strconv.Atoi(bh)
//	block := a.RPCSRC(v["coin"]).GetBlockByHeight(bhi)
//	out, err := json.Marshal(block)
//	if err != nil {
//		log.Print("Error encoding JSON")
//		return
//	}
//	w.Write([]byte(out))
//}
//
//func ViewHash(w http.ResponseWriter, r *http.Request) {
//	v := mux.Vars(r)
//	bh := v["blockhash"]
//	block := (a.RPCSRC(v["coin"]).GetBlock(bh)).(map[string]interface{})
//	h := strconv.FormatInt(block["height"].(int64), 10)
//	http.Redirect(w, r, "/b/"+v["coin"]+"/block/"+h, 301)
//}

func (eq *ExplorerQueries) ViewTx(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	out, err := json.Marshal(eq.GetTx(v["coin"], v["txid"]))
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (eq *ExplorerQueries) ViewAddr(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	var block interface{}
	block = eq.GetBlock(v["coin"], v["id"])
	out, err := json.Marshal(block)
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (eq *ExplorerQueries) ViewRawMemPool(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	rawMemPool := eq.GetMemPool(v["coin"])
	out, err := json.Marshal(rawMemPool)
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (eq *ExplorerQueries) ViewMiningInfo(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	miningInfo := eq.GetMiningInfo(v["coin"])

	out, err := json.Marshal(miningInfo)
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (eq *ExplorerQueries) ViewInfo(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	info := eq.GetInfo(v["coin"])
	out, err := json.Marshal(info)
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (eq *ExplorerQueries) ViewPeers(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	peers := eq.GetPeers(v["coin"])
	out, err := json.Marshal(peers)
	if err != nil {
		log.Print("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
