package explorers

import (
	"fmt"
	"github.com/comhttp/jorm/mod/coins"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"sort"
	"strconv"
	"time"
)

type Explorer struct {
	Status map[string]*BlockchainStatus `json:"status"`
}

type BlockchainStatus struct {
	Blocks    int `json:"blocks"`
	Txs       int `json:"txs"`
	Addresses int `json:"addresses"`
}

func GetExplorer(j *jdb.JDB) *Explorer {
	n := coins.GetNodeCoins(j)
	e := Explorer{
		Status: map[string]*BlockchainStatus{},
	}
	for _, node := range n.C {
		s := BlockchainStatus{}
		err := j.Read(node.Slug, "status", &s)
		utl.ErrorLog(err)
		e.Status[node.Slug] = &s
	}
	//e.j = j
	return &e
}

func GetBlock(j *jdb.JDB, c, id string) map[string]interface{} {
	blockHash := ""
	block := make(map[string]interface{})
	_, err := strconv.Atoi(id)
	if err != nil {
		blockHash = id
	} else {
		blockHash = ""
		err = j.Read(c, "block_"+id, &blockHash)
	}
	err = j.Read(c, "block_"+blockHash, &block)
	utl.ErrorLog(err)
	return block
}
func (e *Explorer) GetBlocks(j *jdb.JDB, c string, per, page int) (blocks []map[string]interface{}) {
	blockCount := e.Status[c].Blocks
	fmt.Println("blockCount", blockCount)
	startBlock := blockCount - per*page
	minusBlockStart := int(startBlock + per)
	for ibh := minusBlockStart; ibh >= startBlock; {
		blocks = append(blocks, GetBlockShort(j, c, strconv.Itoa(ibh)))
		ibh--
	}
	sort.SliceStable(blocks, func(i, j int) bool {
		return blocks[i]["height"].(int) < blocks[j]["height"].(int)
	})
	return blocks
}
func GetBlockShort(j *jdb.JDB, coin, blockhash string) map[string]interface{} {
	b := GetBlock(j, coin, blockhash)
	block := make(map[string]interface{})
	if b["bits"] != nil {
		block["bits"] = b["bits"].(string)
	}
	if b["confirmations"] != nil {
		block["confirmations"] = int64(b["confirmations"].(float64))
	}
	if b["difficulty"] != nil {
		block["difficulty"] = b["difficulty"].(float64)
	}
	if b["hash"] != nil {
		block["hash"] = b["hash"].(string)
	}
	if b["height"] != nil {
		block["height"] = int64(b["height"].(float64))
	}
	if b["tx"] != nil {
		var txsNumber int
		for _ = range b["tx"].([]interface{}) {
			txsNumber++
		}
		block["txs"] = txsNumber
	}
	if b["size"] != nil {
		block["size"] = int64(b["size"].(float64))
	}
	if b["time"] != nil {
		unixTimeUTC := time.Unix(int64(b["time"].(float64)), 0)
		block["time"] = unixTimeUTC.Format(time.RFC850)
		block["timeutc"] = unixTimeUTC.Format(time.RFC3339)
	}
	return block
}

func GetTx(j *jdb.JDB, c, id string) map[string]interface{} {
	tx := make(map[string]interface{})
	err := j.Read(c, "tx_"+id, &tx)
	utl.ErrorLog(err)
	return tx
}
func GetAddr(j *jdb.JDB, c, id string) map[string]interface{} {
	addr := make(map[string]interface{})
	err := j.Read(c, "addr_"+id, &addr)
	utl.ErrorLog(err)
	return addr
}

func GetMemPool(j *jdb.JDB, c string) []string {
	mempool := []string{}
	err := j.Read("info", c+"_mempool", &mempool)
	utl.ErrorLog(err)
	return mempool
}

func GetMiningInfo(j *jdb.JDB, c string) map[string]interface{} {
	mininginfo := make(map[string]interface{})
	err := j.Read("info", c+"_mining", &mininginfo)
	utl.ErrorLog(err)
	return mininginfo
}

func GetInfo(j *jdb.JDB, c string) map[string]interface{} {
	info := make(map[string]interface{})
	err := j.Read("info", c+"_info", &info)
	utl.ErrorLog(err)
	return info
}

func GetNetworkInfo(j *jdb.JDB, c string) map[string]interface{} {
	network := make(map[string]interface{})
	err := j.Read("info", c+"_network", &network)
	utl.ErrorLog(err)
	return network
}

func GetPeers(j *jdb.JDB, c string) []interface{} {
	peers := new([]interface{})
	err := j.Read("info", c+"_peers", &peers)
	utl.ErrorLog(err)
	return *peers
}
