package explorer

import (
	"fmt"
	jdb2 "github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"strconv"
	"time"
)

func GetExplorer(j *jdb2.JDB) *Explorer {
	e := Explorer{}
	err := j.Read("info", "explorer", &e)
	utl.ErrorLog(err)
	e.j = j
	return &e
}

func (e *Explorer) GetBlock(c, id string) map[string]interface{} {
	blockHash := ""
	block := make(map[string]interface{})
	_, err := strconv.Atoi(id)
	if err != nil {
		blockHash = id
	} else {
		blockHash = ""
		err = e.j.Read(c, "block_"+id, &blockHash)
	}
	err = e.j.Read(c, "block_"+blockHash, &block)
	utl.ErrorLog(err)
	return block
}
func (e *Explorer) GetBlocks(c string, per, page int) (blocks []map[string]interface{}) {
	blockCount := e.Status[c].Blocks
	fmt.Println("blockCount", blockCount)
	startBlock := blockCount - per*page
	minusBlockStart := int(startBlock + per)
	for ibh := minusBlockStart; ibh >= startBlock; {
		blocks = append(blocks, e.GetBlockShort(c, strconv.Itoa(ibh)))
		ibh--
	}
	return blocks
}
func (e *Explorer) GetBlockShort(c, blockhash string) map[string]interface{} {
	b := e.GetBlock(c, blockhash)
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

func (e *Explorer) GetTx(c, id string) map[string]interface{} {
	tx := make(map[string]interface{})
	err := e.j.Read(c, "tx_"+id, &tx)
	utl.ErrorLog(err)
	return tx
}
func (e *Explorer) GetAddr(c, id string) map[string]interface{} {
	addr := make(map[string]interface{})
	err := e.j.Read(c, "addr_"+id, &addr)
	utl.ErrorLog(err)
	return addr
}

func (e *Explorer) GetMemPool(c string) []string {
	mempool := []string{}
	err := e.j.Read("info", c+"_mempool", &mempool)
	utl.ErrorLog(err)
	return mempool
}

func (e *Explorer) GetMiningInfo(c string) map[string]interface{} {
	mininginfo := make(map[string]interface{})
	err := e.j.Read("info", c+"_mining", &mininginfo)
	utl.ErrorLog(err)
	return mininginfo
}

func (e *Explorer) GetInfo(c string) map[string]interface{} {
	info := make(map[string]interface{})
	err := e.j.Read("info", c+"_info", &info)
	utl.ErrorLog(err)
	return info
}

func (e *Explorer) GetNetworkInfo(c string) map[string]interface{} {
	network := make(map[string]interface{})
	err := e.j.Read("info", c+"_network", &network)
	utl.ErrorLog(err)
	return network
}

func (e *Explorer) GetPeers(c string) []interface{} {
	peers := new([]interface{})
	err := e.j.Read("info", c+"_peers", &peers)
	utl.ErrorLog(err)
	return *peers
}
