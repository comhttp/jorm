package explorer

import (
	"fmt"
	"github.com/comhttp/jorm/app/jorm/a"
	"github.com/comhttp/jorm/cfg"
	"github.com/comhttp/jorm/coins"
	"github.com/comhttp/jorm/jdb"
	"github.com/comhttp/jorm/nodes"
	"github.com/comhttp/jorm/pkg/utl"
	"path/filepath"
)

type Explorer struct {
	Status map[string]uint64 `json:"status"`
}

// GetExplorer updates the data from blockchain of a coin in the database
func ExploreCoins(j *jdb.JDB, c coins.NodeCoins) {
	var b []string
	for _, coin := range c.C {
		var bn nodes.BitNoded
		fmt.Println("Coin is BitNode:", coin.Name)
		if utl.FileExists(filepath.FromSlash(cfg.Path + "nodes/" + coin.Slug)) {
			b = append(b, coin.Slug)
			for _, bitnode := range coin.Nodes {
				go GetCoinBlockchain(j, bitnode)
				fmt.Println("GetCoinBlockchain:", coin.Name)
				bn.Coin = coin.Slug
			}
			//bns = append(bns, bn)
			//nodes := jdb.ReadData(filepath.FromSlash(cfg.C.Out + "/nodes"))
			//ns := make(n.Nodes, len(nodes))
			//
			//for i := range nodes {
			//	if err := json.Unmarshal(nodes[i], &ns[i]); err != nil {
			//		fmt.Println("Error", err)
			//	}
			//}
		}
	}

	//et := mod.Cache{Data: e}
	//jdb.JDB.Write(cfg.Web, "explorer", et)
}

// GetExplorer returns the full set of information about a block
func GetCoinBlockchain(j *jdb.JDB, b a.BitNode) {
	b.Jrc = utl.NewClient(cfg.C.RPC.Username, cfg.C.RPC.Password, b.IP, b.Port)
	if b.Jrc != nil {
		//if utl.FileExists(cfg.Path + cfg.C.Out + "/info/explorer") {
		e := GetExplorer(j)
		//if err := jdb.JDB.Read(cfg.C.Out+"/info", "explorer", &e); err != nil {
		//	fmt.Println("Error", err)
		//}
		blockCount := b.GetBlockCount()

		fmt.Println("Block Count: ", blockCount)

		if e.Status != nil && blockCount >= int(e.Status["blocks"]) {
			for {
				//e.block(a, www)
				fmt.Println("BlocksPre", e.Status["blocks"])
				blocksIndex := map[uint64]string{}
				//if err := jdb.JDB.Read(cfg.C.Out+"/index", "blocks", &blocksIndex); err != nil {
				//	fmt.Println("Error", err)
				//}
				e.Status["blocks"]++
				blockRaw := b.GetBlockByHeight(int(e.Status["blocks"]))
				if blockRaw != nil && blockRaw != "" {
					blockHash := blockRaw.(map[string]interface{})["hash"].(string)
					blocksIndex[e.Status["blocks"]] = blockHash
					//jdb.JDB.Write(cfg.C.Out+"/blocks", blockHash, blockRaw)
					block := (blockRaw).(map[string]interface{})
					if e.Status["blocks"] != 0 {
						for _, t := range (block["tx"]).([]interface{}) {
							e.tx(b, t.(string))
						}
					}
					fmt.Println("BlocksPosle", e.Status["blocks"])
					//jdb.JDB.Write(cfg.C.Out+"/info", "explorer", e)
					//jdb.JDB.Write(cfg.C.Out+"/index", "blocks", blocksIndex)
				} else {
					break
				}
			}
			//}
			//} else {
			//e := &Explorer{
			//	Status: map[string]uint64{"blocks": 0, "txs": 0, "addresses": 0},
			//}
			//jdb.JDB.Write(cfg.C.Out+"/info", "explorer", e)
			//jdb.JDB.Write(cfg.C.Out+"/explorer/index", "blocks", map[uint64]string{})
			//jdb.JDB.Write(cfg.C.Out+"/explorer/index", "txs", map[uint64]string{})
			//jdb.JDB.Write(cfg.C.Out+"/explorer/index", "addresses", map[uint64]string{})

			fmt.Println("ExplorerExplorerExplorerExplorer", cfg.C.Out+"/info")

		}
	}
	return
}

func (e *Explorer) tx(a a.BitNode, txid string) {
	txRaw := a.GetTx(txid)
	txsIndex := map[uint64]string{}
	//if err := jdb.JDB.Read(cfg.C.Out+"/index", "txs", &txsIndex); err != nil {
	//	fmt.Println("Error", err)
	//}
	e.Status["txs"]++
	txsIndex[e.Status["txs"]] = txid
	//fmt.Println("txid", txid)
	//go jdb.JDB.Write(cfg.C.Out+"/txs", txid, txRaw)
	if txRaw != nil {
		tx := (txRaw).(map[string]interface{})
		if tx["vout"] != nil {
			for _, nRaw := range tx["vout"].([]interface{}) {
				if nRaw.(map[string]interface{})["scriptPubKey"] != nil {
					scriptPubKey := nRaw.(map[string]interface{})["scriptPubKey"].(map[string]interface{})
					if scriptPubKey["addresses"] != nil {
						for _, address := range (scriptPubKey["addresses"]).([]interface{}) {
							e.addr(cfg.C.Out, address.(string))
						}
					}
				}
			}
		}
	}
	//jdb.JDB.Write(cfg.C.Out+"/index", "txs", txsIndex)
	return
}

func (e *Explorer) addr(www, address string) {
	addressesIndex := map[uint64]string{}
	//if err := jdb.JDB.Read(www+"/index", "addresses", &addressesIndex); err != nil {
	//	fmt.Println("Error", err)
	//}
	//addressData := new(interface{})
	//if err := jdb.JDB.Read(www, "explorer", &e); err != nil {
	//	fmt.Println("Error", err)
	//}
	//addressData := address
	e.Status["addresses"]++
	addressesIndex[e.Status["addresses"]] = address
	//go jdb.JDB.Write(www+"/addresses", address, addressData)
	//jdb.JDB.Write(www+"/index", "addresses", addressesIndex)
	//fmt.Println("address", address)
	return
}

func (e *Explorer) status() {
	//s =
}
