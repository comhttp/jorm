package explorer

import (
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"sort"
	"strconv"
	"time"
)

//
//type Explorers struct {
//	Status map[string]*BlockchainStatus `json:"status"`
//}

func GetExplorers(j *jdb.JDBS, nodeCoins []string) {
	//status: &BlockchainStatus{},
	for _, coin := range nodeCoins {
		s := &BlockchainStatus{}
		err := j.B[coin].Read(coin, "status", &s)
		utl.ErrorLog(err)
		//j.Explorers[coin].status = s
	}
	return
}

func (eq *ExplorersQueries) GetExplorer(coin string) *BlockchainStatus {
	s := &BlockchainStatus{}
	err := eq.j.Read(coin, "status", &s)
	utl.ErrorLog(err)
	return s
}

func (eq *ExplorersQueries) GetBlock(c, id string) map[string]interface{} {
	blockHash := ""
	block := make(map[string]interface{})
	_, err := strconv.Atoi(id)
	if err != nil {
		blockHash = id
	} else {
		blockHash = ""
		err = eq.j.Read(c, "block_"+id, &blockHash)
	}
	err = eq.j.Read(c, "block_"+blockHash, &block)
	utl.ErrorLog(err)
	return block
}

func (eq *ExplorersQueries) GetBlocks(coin string, per, page int) (blocks []map[string]interface{}) {
	s := &BlockchainStatus{}
	err := eq.j.Read(coin, "status", &s)
	utl.ErrorLog(err)
	blockCount := s.Blocks
	//app.log.Print("blockCount", blockCount)
	startBlock := blockCount - per*page
	minusBlockStart := int(startBlock + per)
	for ibh := minusBlockStart; ibh >= startBlock; {
		blocks = append(blocks, eq.GetBlockShort(coin, strconv.Itoa(ibh)))
		ibh--
	}
	sort.SliceStable(blocks, func(i, j int) bool {
		return int(blocks[i]["height"].(int64)) < int(blocks[j]["height"].(int64))
	})
	return blocks
}
func (eq *ExplorersQueries) GetBlockShort(coin, blockhash string) map[string]interface{} {
	b := eq.GetBlock(coin, blockhash)
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

func (eq *ExplorersQueries) GetTx(c, id string) map[string]interface{} {
	tx := make(map[string]interface{})
	err := eq.j.Read(c, "tx_"+id, &tx)
	utl.ErrorLog(err)
	return tx
}
func (eq *ExplorersQueries) GetAddr(c, id string) map[string]interface{} {
	addr := make(map[string]interface{})
	err := eq.j.Read(c, "addr_"+id, &addr)
	utl.ErrorLog(err)
	return addr
}

func (eq *ExplorersQueries) GetMemPool(c string) []string {
	mempool := []string{}
	err := eq.j.Read("info", c+"_mempool", &mempool)
	utl.ErrorLog(err)
	return mempool
}

func (eq *ExplorersQueries) GetMiningInfo(c string) map[string]interface{} {
	mininginfo := make(map[string]interface{})
	err := eq.j.Read("info", c+"_mining", &mininginfo)
	utl.ErrorLog(err)
	return mininginfo
}

func (eq *ExplorersQueries) GetInfo(c string) map[string]interface{} {
	info := make(map[string]interface{})
	err := eq.j.Read("info", c+"_info", &info)
	utl.ErrorLog(err)
	return info
}

func (eq *ExplorersQueries) GetNetworkInfo(c string) map[string]interface{} {
	network := make(map[string]interface{})
	err := eq.j.Read("info", c+"_network", &network)
	utl.ErrorLog(err)
	return network
}

func (eq *ExplorersQueries) GetPeers(c string) []interface{} {
	peers := new([]interface{})
	err := eq.j.Read("info", c+"_peers", &peers)
	utl.ErrorLog(err)
	return *peers
}
