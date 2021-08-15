package explorer

import (
	"github.com/comhttp/jorm/mod/nodes"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/rs/zerolog/log"
	"strconv"
)

// GetExplorer updates the data from blockchain of a coin in the database
func (e Explorers) ExploreCoin(j *jdb.JDB, username, password, coin string) {
	var b []string
	log.Print("Coin is BitNode:", coin)
	if e[coin].BitNodes != nil {
		b = append(b, coin)
		log.Print("bitnodebitnodebitnodebitnodebitnodebitnodebitnode", e[coin].BitNodes)
		for _, bitnode := range e[coin].BitNodes {
			log.Print("bitnodebitnodebitnodebitnodebitnodebitnodebitnode", bitnode)
			bitnode.Jrc = utl.NewClient(username, password, bitnode.IP, bitnode.Port)
			log.Print("Get Coin Blockchain:", coin)
			go e[coin].GetCoinBlockchain(j, &bitnode, coin)
		}
	}
}

// GetExplorer returns the full set of information about a block
func (e *Explorer) GetCoinBlockchain(j *jdb.JDB, b *nodes.BitNode, c string) {
	if b.Jrc != nil {
		blockCount := b.APIGetBlockCount()
		log.Print("Block Count from the chain: ", blockCount)
		log.Print("Status :  :::"+c+" - - ", e.Status.Blocks)
		if blockCount >= e.Status.Blocks {
			e.blocks(j, b, blockCount, c)
		}
	}
}

func (e *Explorer) blocks(j *jdb.JDB, b *nodes.BitNode, bc int, c string) {
	for {
		blockRaw := b.APIGetBlockByHeight(e.Status.Blocks)
		if blockRaw != nil && blockRaw != "" {
			blockHash := blockRaw.(map[string]interface{})["hash"].(string)
			j.Write(c, "block_"+strconv.Itoa(e.Status.Blocks), blockHash)
			j.Write(c, "block_"+blockHash, blockRaw)
			block := (blockRaw).(map[string]interface{})
			if e.Status.Blocks != 0 {
				for _, t := range (block["tx"]).([]interface{}) {
					e.tx(j, b, c, t.(string))
				}
			}
			bl := blockRaw.(map[string]interface{})
			e.Status.Blocks = int(bl["height"].(float64))
			log.Print("Write "+c+" block: "+strconv.Itoa(e.Status.Blocks)+" - ", blockHash)
			j.Write(c, "status", e.Status)
		} else {
			break
		}
		if bc != 0 {
			e.Status.Blocks++
		}
		log.Print("StatusBlocks   "+c, e.Status.Blocks)
	}
}

func (e *Explorer) tx(j *jdb.JDB, b *nodes.BitNode, c, txid string) {
	txRaw := b.APIGetTx(txid)
	e.Status.Txs++
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
	log.Print("Write "+c+" transaction: ", txid)
	return
}

func (e *Explorer) addr(j *jdb.JDB, c, address string) {
	e.Status.Addresses++
	j.Write(c, "addr_"+address, address)
	log.Print("Write "+c+" address: ", address)
	return
}

//func (e *explorer.Explorer) status(n *nodes.BitNode) {
//
//	log.Print("Mempool: ", n.GetRawMemPool())
//	log.Print("MiningInfo: ", n.GetMiningInfo())
//	log.Print("NetworkInfo: ", n.GetNetworkInfo())
//	log.Print("Info: ", n.GetInfo())
//	log.Print("PeerInfo: ", n.GetPeerInfo())
//
//
//	//n.addNode(ip string)
//	//n.GetAddNodeInfo(ip string)
//}
//
//
