package nodes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (b *bitNodesRPC) directViewBlocks(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	// per, _ := strconv.Atoi(v["per"])
	page, _ := strconv.Atoi(v["page"])

	// rpc := RPCSRC(bitNodes(b.path, v["coin"]), b.username, b.password)

	lb := map[string]interface{}{
		"d": map[string]interface{}{
			"currentPage": page,
			// "pageCount":   rpc.APIGetBlockCount() / per,
			// "blocks":      rpc.APIGetBlocks(per, page),
		},
	}

	out, err := json.Marshal(lb)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (b *bitNodesRPC) directLastBlock(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)

	rpc := RPCSRC(bitNodes(b.path, v["coin"]), b.username, b.password)
	lastblock := rpc.APIGetBlockCount()
	bl := map[string]interface{}{
		"d": lastblock,
	}
	out, err := json.Marshal(bl)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (b *bitNodesRPC) directViewBlockHeight(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	bh := v["blockheight"]
	// node := Node{}
	rpc := RPCSRC(bitNodes(b.path, v["coin"]), b.username, b.password)
	bhi, _ := strconv.Atoi(bh)
	block := rpc.GetBlockByHeight(bhi)

	bl := map[string]interface{}{
		"d": block,
	}
	fmt.Println("IP RPC source:", block)
	out, err := json.Marshal(bl)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func (b *bitNodesRPC) directViewBlock(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	bh := v["blockhash"]
	rpc := RPCSRC(bitNodes(b.path, v["coin"]), b.username, b.password)
	block := rpc.APIGetBlock(bh)
	h := strconv.FormatInt(block["height"].(int64), 10)
	http.Redirect(w, r, "/b/"+v["coin"]+"/block/"+h, 301)
}

func (b *bitNodesRPC) directViewTx(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	txid := v["txid"]
	// node := Node{}
	rpc := RPCSRC(bitNodes(b.path, v["coin"]), b.username, b.password)
	tX := rpc.APIGetTx(txid)

	tx := map[string]interface{}{
		"d": tX,
	}
	out, err := json.Marshal(tx)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (b *bitNodesRPC) directViewRawMemPool(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	rpc := RPCSRC(bitNodes(b.path, v["coin"]), b.username, b.password)
	rawMemPool := rpc.APIGetRawMemPool()
	rmp := map[string]interface{}{
		"d": rawMemPool,
	}
	out, err := json.Marshal(rmp)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (b *bitNodesRPC) directViewMiningInfo(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	rpc := RPCSRC(bitNodes(b.path, v["coin"]), b.username, b.password)
	miningInfo := rpc.APIGetMiningInfo()

	mi := map[string]interface{}{
		"d": miningInfo,
	}
	out, err := json.Marshal(mi)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (b *bitNodesRPC) directViewInfo(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	rpc := RPCSRC(bitNodes(b.path, v["coin"]), b.username, b.password)
	info := rpc.APIGetInfo()

	in := map[string]interface{}{
		"d": info,
	}
	out, err := json.Marshal(in)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func (b *bitNodesRPC) directViewPeers(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	rpc := RPCSRC(bitNodes(b.path, v["coin"]), b.username, b.password)
	info := rpc.APIGetPeerInfo()
	pi := map[string]interface{}{
		"d": info,
	}
	out, err := json.Marshal(pi)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
