package explorer

import (
	"fmt"
	nodes2 "github.com/comhttp/jorm/mod/nodes"
	cfg "github.com/comhttp/jorm/pkg/cfg"
	jdb2 "github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"path/filepath"
	"strconv"
)

type Explorer struct {
	Status map[string]*Blockchain `json:"status"`
	j      *jdb2.JDB
}

type Blockchain struct {
	Blocks    int `json:"blocks"`
	Txs       int `json:"txs"`
	Addresses int `json:"addresses"`
}

// GetExplorer updates the data from blockchain of a coin in the database
func (e *Explorer) ExploreCoins(c nodes2.NodeCoins) {
	var b []string
	for _, coin := range c.C {
		var bn nodes2.BitNoded
		fmt.Println("Coin is BitNode:", coin.Name)
		if utl.FileExists(filepath.FromSlash(cfg.Path + "nodes/" + coin.Slug)) {
			b = append(b, coin.Slug)
			for _, bitnode := range coin.Nodes {
				bitnode.Jrc = utl.NewClient(cfg.C.RPC.Username, cfg.C.RPC.Password, bitnode.IP, bitnode.Port)
				//e.status(&bitnode)
				go e.GetCoinBlockchain(&bitnode, coin.Slug)
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
func (e *Explorer) GetCoinBlockchain(b *nodes2.BitNode, c string) {
	if b.Jrc != nil {
		blockCount := b.APIGetBlockCount()
		fmt.Println("Block Count: ", blockCount)
		fmt.Println("Be.Status: ", e.Status)

		if e.Status == nil {
			e.Status = make(map[string]*Blockchain)
		}
		if e.Status[c] == nil {
			e.Status[c] = &Blockchain{
				Blocks:    0,
				Txs:       0,
				Addresses: 0,
			}
		}
		if blockCount >= e.Status[c].Blocks {
			e.blocks(b, c)
		}
		e.j.Write("info", "explorer", e)
		fmt.Println("ExplorerExplorerExplorerExplorer", cfg.C.Out+"/info")
	}
}

func (e *Explorer) blocks(b *nodes2.BitNode, c string) {
	for {
		e.Status[c].Blocks++
		blockRaw := b.APIGetBlockByHeight(e.Status[c].Blocks)
		if blockRaw != nil && blockRaw != "" {
			blockHash := blockRaw.(map[string]interface{})["hash"].(string)
			//blocksIndex[e.Status[c].Blocks] = blockHash
			e.j.Write(c, "block_"+strconv.Itoa(e.Status[c].Blocks), blockHash)
			e.j.Write(c, "block_"+blockHash, blockRaw)
			block := (blockRaw).(map[string]interface{})
			if e.Status[c].Blocks != 0 {
				for _, t := range (block["tx"]).([]interface{}) {
					e.tx(b, c, t.(string))

				}
			}
			//fmt.Println("Write " + c + " block: " + strconv.Itoa(e.Status[c].Blocks)+" - ", blockHash)
			e.j.Write("info", "explorer", e)
		} else {
			break
		}
	}
}

func (e *Explorer) tx(rpc *nodes2.BitNode, c, txid string) {
	txRaw := APIGetTx(rpc, txid)
	e.Status[c].Txs++
	e.j.Write(c, "tx_"+txid, txRaw)
	if txRaw != nil {
		tx := (txRaw).(map[string]interface{})
		if tx["vout"] != nil {
			for _, nRaw := range tx["vout"].([]interface{}) {
				if nRaw.(map[string]interface{})["scriptPubKey"] != nil {
					scriptPubKey := nRaw.(map[string]interface{})["scriptPubKey"].(map[string]interface{})
					if scriptPubKey["addresses"] != nil {
						for _, address := range (scriptPubKey["addresses"]).([]interface{}) {
							e.addr(e.j, c, address.(string))
						}
					}
				}
			}
		}
	}
	//fmt.Println("Write " + c + " transaction: ", txid)
	return
}

func (e *Explorer) addr(j *jdb2.JDB, c, address string) {
	//addressData := new(interface{})
	//if err := jdb.JDB.Read(www, "explorer", &e); err != nil {
	//	fmt.Println("Error", err)
	//}
	//addressData := address
	e.Status[c].Addresses++
	//addressesIndex[e.Status[c].Addresses] = address
	j.Write(c, "addr_"+address, address)

	//go jdb.JDB.Write(www+"/addresses", address, addressData)
	//jdb.JDB.Write(www+"/index", "addresses", addressesIndex)
	//fmt.Println("address", address)
	//fmt.Println("Write " + c + " address: ", address)

	return
}

//func (e *Explorer) status(n *nodes.BitNode) {
//
//	fmt.Println("Mempool: ", n.GetRawMemPool())
//	fmt.Println("MiningInfo: ", n.GetMiningInfo())
//	fmt.Println("NetworkInfo: ", n.GetNetworkInfo())
//	fmt.Println("Info: ", n.GetInfo())
//	fmt.Println("PeerInfo: ", n.GetPeerInfo())
//
//
//	//n.addNode(ip string)
//	//n.GetAddNodeInfo(ip string)
//}
//
//
