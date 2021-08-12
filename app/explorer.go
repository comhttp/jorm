package app

import (
	"github.com/comhttp/jorm/mod/nodes"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"log"
	"strconv"
)

// GetExplorer updates the data from blockchain of a coin in the database
func (j *JORM) ExploreCoin(username, password, coin string) {
	var b []string
	log.Println("Coin is BitNode:", coin)
	if j.Explorers[coin].BitNodes != nil {
		b = append(b, coin)
		log.Println("bitnodebitnodebitnodebitnodebitnodebitnodebitnode", j.Explorers[coin].BitNodes)

		for _, bitnode := range j.Explorers[coin].BitNodes {

			log.Println("bitnodebitnodebitnodebitnodebitnodebitnodebitnode", bitnode)

			bitnode.Jrc = utl.NewClient(username, password, bitnode.IP, bitnode.Port)
			log.Println("Get Coin Blockchain:", coin)
			go j.Explorers[coin].GetCoinBlockchain(j.JDB, &bitnode, coin)
		}
	}
}

// GetExplorer returns the full set of information about a block
func (e *Explorer) GetCoinBlockchain(j *jdb.JDB, b *nodes.BitNode, c string) {
	if b.Jrc != nil {
		blockCount := b.APIGetBlockCount()
		log.Println("Block Count from the chain: ", blockCount)
		log.Println("Status :  :::"+c+" - - ", e.status.Blocks)
		if blockCount >= e.status.Blocks {
			e.blocks(j, b, blockCount, c)
		}
	}
}

func (e *Explorer) blocks(j *jdb.JDB, b *nodes.BitNode, bc int, c string) {
	for {
		blockRaw := b.APIGetBlockByHeight(e.status.Blocks)
		if blockRaw != nil && blockRaw != "" {
			blockHash := blockRaw.(map[string]interface{})["hash"].(string)
			j.Write(c, "block_"+strconv.Itoa(e.status.Blocks), blockHash)
			j.Write(c, "block_"+blockHash, blockRaw)
			block := (blockRaw).(map[string]interface{})
			if e.status.Blocks != 0 {
				for _, t := range (block["tx"]).([]interface{}) {
					e.tx(j, b, c, t.(string))
				}
			}
			bl := blockRaw.(map[string]interface{})
			e.status.Blocks = int(bl["height"].(float64))
			log.Println("Write "+c+" block: "+strconv.Itoa(e.status.Blocks)+" - ", blockHash)
			j.Write(c, "status", e.status)
		} else {
			break
		}
		if bc != 0 {
			e.status.Blocks++
		}
		log.Println("StatusBlocks   "+c, e.status.Blocks)
	}
}

func (e *Explorer) tx(j *jdb.JDB, b *nodes.BitNode, c, txid string) {
	txRaw := b.APIGetTx(txid)
	e.status.Txs++
	j.Write(c, "tx_"+txid, txRaw)
	if txRaw != nil {
		tx := (txRaw).(map[string]interface{})
		if tx["vout"] != nil {
			for _, nRaw := range tx["vout"].([]interface{}) {
				if nRaw.(map[string]interface{})["scriptPubKey"] != nil {
					scriptPubKey := nRaw.(map[string]interface{})["scriptPubKey"].(map[string]interface{})
					if scriptPubKey["addresses"] != nil {
						for _, address := range (scriptPubKey["addresses"]).([]interface{}) {
							e.addr(j, c, address.(string))
						}
					}
				}
			}
		}
	}
	log.Println("Write "+c+" transaction: ", txid)
	return
}

func (e *Explorer) addr(j *jdb.JDB, c, address string) {
	e.status.Addresses++
	j.Write(c, "addr_"+address, address)
	log.Println("Write "+c+" address: ", address)
	return
}

//func (e *explorer.Explorer) status(n *nodes.BitNode) {
//
//	log.Println("Mempool: ", n.GetRawMemPool())
//	log.Println("MiningInfo: ", n.GetMiningInfo())
//	log.Println("NetworkInfo: ", n.GetNetworkInfo())
//	log.Println("Info: ", n.GetInfo())
//	log.Println("PeerInfo: ", n.GetPeerInfo())
//
//
//	//n.addNode(ip string)
//	//n.GetAddNodeInfo(ip string)
//}
//
//
