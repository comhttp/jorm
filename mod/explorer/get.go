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

func (eq *ExplorerQueries) GetExplorer(coin string) *BlockchainStatus {
	s := &BlockchainStatus{}
	err := eq.j.B[coin].Read(coin, "status", &s)
	utl.ErrorLog(err)
	return s
}

func (eq *ExplorerQueries) GetStatus(coin string) (*BlockchainStatus, error) {
	eq.status = &BlockchainStatus{}
	err := eq.j.B[coin].Read("info", "status", &eq.status)
	utl.ErrorLog(err)
	return eq.status, err
}

func (eq *ExplorerQueries) GetLastBlock(coin string) int {
	status, err := eq.GetStatus(coin)
	utl.ErrorLog(err)
	return status.Blocks
}

func (eq *ExplorerQueries) GetBlock(coin, id string) map[string]interface{} {
	blockHash := ""
	block := make(map[string]interface{})
	_, err := strconv.Atoi(id)
	if err != nil {
		blockHash = id
	} else {
		blockHash = ""
		err = eq.j.B[coin].Read("block", id, &blockHash)
	}
	err = eq.j.B[coin].Read("block", blockHash, &block)
	utl.ErrorLog(err)
	return block
}

func (eq *ExplorerQueries) GetBlocks(coin string, per, page int) (blocks []map[string]interface{}) {
	s := &BlockchainStatus{}
	err := eq.j.B[coin].Read("info", "status", &s)
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
		return int(blocks[i]["height"].(int64)) > int(blocks[j]["height"].(int64))
	})
	return blocks
}
func (eq *ExplorerQueries) GetBlockShort(coin, blockhash string) map[string]interface{} {
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

func (eq *ExplorerQueries) GetTx(coin, id string) map[string]interface{} {
	tx := make(map[string]interface{})
	err := eq.j.B[coin].Read("tx", id, &tx)
	utl.ErrorLog(err)
	return tx
}
func (eq *ExplorerQueries) GetAddr(coin, id string) map[string]interface{} {
	addr := make(map[string]interface{})
	err := eq.j.B[coin].Read("addr", id, &addr)
	utl.ErrorLog(err)
	return addr
}

func (eq *ExplorerQueries) GetMemPool(coin string) []string {
	mempool := []string{}
	err := eq.j.B[coin].Read("info", "mempool", &mempool)
	utl.ErrorLog(err)
	return mempool
}

func (eq *ExplorerQueries) GetMiningInfo(coin string) map[string]interface{} {
	mininginfo := make(map[string]interface{})
	err := eq.j.B[coin].Read("info", "mining", &mininginfo)
	utl.ErrorLog(err)
	return mininginfo
}

func (eq *ExplorerQueries) GetInfo(coin string) map[string]interface{} {
	info := make(map[string]interface{})
	err := eq.j.B[coin].Read("info", "info", &info)
	utl.ErrorLog(err)
	return info
}

func (eq *ExplorerQueries) GetNetworkInfo(coin string) map[string]interface{} {
	network := make(map[string]interface{})
	err := eq.j.B[coin].Read("info", "network", &network)
	utl.ErrorLog(err)
	return network
}

func (eq *ExplorerQueries) GetPeers(coin string) []interface{} {
	peers := new([]interface{})
	err := eq.j.B[coin].Read("info", "peers", &peers)
	utl.ErrorLog(err)
	return *peers
}
