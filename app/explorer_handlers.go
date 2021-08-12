package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (j *JORM) ViewStatus(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	out, err := json.Marshal(j.Explorers[v["coin"]].status)
	if err != nil {
		log.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (j *JORM) ViewBlocks(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	per, _ := strconv.Atoi(v["per"])
	page, _ := strconv.Atoi(v["page"])
	ex := GetExplorer(j.JDB, v["coin"])
	lastblock := ex.Blocks - 1
	log.Println("lastblocklastblocklastblock", lastblock)

	lb := map[string]interface{}{
		"currentPage": page,
		"pageCount":   lastblock / per,
		"blocks":      j.Explorers[v["coin"]].GetBlocks(j.JDB, v["coin"], per, page),
		"lastBlock":   lastblock,
	}

	out, err := json.Marshal(lb)
	if err != nil {
		log.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (j *JORM) LastBlock(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	out, err := json.Marshal(j.Explorers[v["coin"]].status.Blocks)
	if err != nil {
		log.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (j *JORM) ViewBlock(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	out, err := json.Marshal(GetBlock(j.JDB, v["coin"], v["id"]))
	if err != nil {
		log.Println("Error encoding JSON")
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
//		log.Println("Error encoding JSON")
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

func (j *JORM) ViewTx(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	out, err := json.Marshal(GetTx(j.JDB, v["coin"], v["txid"]))
	if err != nil {
		log.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (j *JORM) ViewAddr(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	var block interface{}
	block = GetBlock(j.JDB, v["coin"], v["id"])
	out, err := json.Marshal(block)
	if err != nil {
		log.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (j *JORM) ViewRawMemPool(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	rawMemPool := GetMemPool(j.JDB, v["coin"])
	out, err := json.Marshal(rawMemPool)
	if err != nil {
		log.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (j *JORM) ViewMiningInfo(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	miningInfo := GetMiningInfo(j.JDB, v["coin"])

	out, err := json.Marshal(miningInfo)
	if err != nil {
		log.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (j *JORM) ViewInfo(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	info := GetInfo(j.JDB, v["coin"])
	out, err := json.Marshal(info)
	if err != nil {
		log.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (j *JORM) ViewPeers(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	peers := GetPeers(j.JDB, v["coin"])
	out, err := json.Marshal(peers)
	if err != nil {
		log.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
