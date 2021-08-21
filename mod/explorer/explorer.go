package explorer

import (
	"github.com/comhttp/jorm/mod/nodes"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/rs/zerolog/log"
	"strconv"
)

// GetExplorer updates the data from blockchain of a coin in the database
func (eq *ExplorerQueries) ExploreCoin(bn nodes.BitNodes, username, password, coin string) {
	var b []string
	log.Print("Coin is BitNode:", coin)
	if bn != nil {
		b = append(b, coin)
		for _, bitnode := range bn {
			log.Print("Bitnode: ", bitnode)
			bitnode.Jrc = utl.NewClient(username, password, bitnode.IP, bitnode.Port)
			eq.j.Write("info", "info", bitnode.APIGetInfo())
			eq.j.Write("info", "peers", bitnode.APIGetPeerInfo())
			eq.j.Write("info", "mempool", bitnode.APIGetRawMemPool())
			eq.j.Write("info", "mining", bitnode.APIGetMiningInfo())
			eq.j.Write("info", "network", bitnode.APIGetNetworkInfo())
			log.Print("Get Coin Blockchain:", coin)
			eq.blockchain(&bitnode, coin)
		}
	}
}

// GetExplorer returns the full set of information about a block
func (eq *ExplorerQueries) blockchain(bn *nodes.BitNode, coin string) {
	if bn.Jrc != nil {
		blockCount := bn.APIGetBlockCount()
		log.Print("Block Count from the chain: ", blockCount)
		log.Print("Status :  :::"+coin+" - - ", eq.status.Blocks)
		if blockCount >= eq.status.Blocks {
			eq.blocks(bn, blockCount, coin)
		}
	}
}

func (eq *ExplorerQueries) blocks(b *nodes.BitNode, bc int, coin string) {
	for {
		blockRaw := b.APIGetBlockByHeight(eq.status.Blocks)
		if blockRaw != nil && blockRaw != "" {
			blockHash := blockRaw.(map[string]interface{})["hash"].(string)
			eq.j.Write("block", strconv.Itoa(eq.status.Blocks), blockHash)
			eq.j.Write("block", blockHash, blockRaw)
			block := (blockRaw).(map[string]interface{})
			if eq.status.Blocks != 0 {
				for _, t := range (block["tx"]).([]interface{}) {
					eq.tx(b, coin, t.(string))
				}
			}
			bl := blockRaw.(map[string]interface{})
			eq.status.Blocks = int(bl["height"].(float64))
			log.Print("Write "+coin+" block: "+strconv.Itoa(eq.status.Blocks)+" - ", blockHash)
			eq.j.Write("info", "status", eq.status)
		} else {
			break
		}
		if bc != 0 {
			eq.status.Blocks++
		}
		log.Print("StatusBlocks   "+coin, eq.status.Blocks)
	}
}

func (eq *ExplorerQueries) tx(b *nodes.BitNode, coin, txid string) {
	txRaw := b.APIGetTx(txid)
	eq.status.Txs++
	eq.j.Write("tx", txid, txRaw)
	if txRaw != nil {
		tx := (txRaw).(map[string]interface{})
		if tx["vout"] != nil {
			for _, nRaw := range tx["vout"].([]interface{}) {
				if nRaw.(map[string]interface{})["scriptPubKey"] != nil {
					scriptPubKey := nRaw.(map[string]interface{})["scriptPubKey"].(map[string]interface{})
					if scriptPubKey["addresses"] != nil {
						for _, address := range (scriptPubKey["addresses"]).([]interface{}) {
							eq.addr(coin, address.(string))
						}
					}
				}
			}
		}
	}
	log.Print("Write "+coin+" transaction: ", txid)
	return
}

func (eq *ExplorerQueries) addr(coin, address string) {
	eq.status.Addresses++
	eq.j.Write(coin, "addr_"+address, address)
	log.Print("Write "+coin+" address: ", address)
	return
}
