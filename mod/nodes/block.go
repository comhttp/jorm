package nodes

import (
	"fmt"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func (b *BitNode) APIGetBlockCount() (blockcount int) {
	bparams := []int{}
	fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb : :", b)

	gbc, err := b.Jrc.MakeRequest("getblockcount", bparams)
	if err != nil {
		log.Print("Error n call: ", err)
	}
	switch gbc.(type) {
	case float64:
		return int(gbc.(float64))
	case string:
		blockcount, _ := strconv.Atoi(gbc.(string))
		return blockcount
	default:
		//b, _ := strconv.Atoi(gbc.(string))
		return blockcount
	}
	return
}

func (b *BitNode) APIGetBlock(blockhash string) (block map[string]interface{}) {
	bparams := []string{blockhash}
	blockRaw, err := b.Jrc.MakeRequest("getblock", bparams)
	if err != nil {
		log.Print("Jorm Node Get Block Error", err)
	}
	block = blockRaw.(map[string]interface{})
	return
}

func (b *BitNode) APIGetBlockByHeight(blockheight int) (block interface{}) {
	bparams := []int{blockheight}
	blockHash, err := b.Jrc.MakeRequest("getblockhash", bparams)
	if err != nil {
		log.Print("Jorm Node Get Block By Height Error", err)
	}
	if blockHash != nil {
		block = b.APIGetBlock((blockHash).(string))
	}
	return block
}

func (b *BitNode) APIGetTx(txid string) (t interface{}) {
	verbose := int(1)
	var grtx []interface{}
	grtx = append(grtx, txid)
	grtx = append(grtx, verbose)
	t, err := b.Jrc.MakeRequest("getrawtransaction", grtx)
	if err != nil {
		log.Print("Jorm Node Get Tx Error", err)
	}
	return
}

func (b *BitNode) APIGetBlocks(per, page int) (blocks []map[string]interface{}) {
	blockCount := b.APIGetBlockCount()

	startBlock := blockCount - per*page
	minusBlockStart := int(startBlock + per)

	for ibh := minusBlockStart; ibh >= startBlock; {
		blocks = append(blocks, b.GetBlockShortByHeight(ibh))
		ibh--
	}
	return blocks
}

func (b *BitNode) GetBlockShort(blockhash string) map[string]interface{} {
	bparams := []string{blockhash}
	rawb, err := b.Jrc.MakeRequest("getblock", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Block Error", err)
	}
	blockRaw := rawb.(map[string]interface{})

	block := make(map[string]interface{})
	if blockRaw["bits"] != nil {
		block["bits"] = blockRaw["bits"].(string)
	}
	if blockRaw["confirmations"] != nil {
		block["confirmations"] = int64(blockRaw["confirmations"].(float64))
	}
	if blockRaw["difficulty"] != nil {
		block["difficulty"] = blockRaw["difficulty"].(float64)
	}
	if blockRaw["hash"] != nil {
		block["hash"] = blockRaw["hash"].(string)
	}
	if blockRaw["height"] != nil {
		block["height"] = int64(blockRaw["height"].(float64))
	}
	if blockRaw["nTx"] != nil {
		block["ntx"] = int(blockRaw["nTx"].(float64))
	}
	if blockRaw["size"] != nil {
		block["size"] = int64(blockRaw["size"].(float64))
	}
	if blockRaw["time"] != nil {
		unixTimeUTC := time.Unix(int64(blockRaw["time"].(float64)), 0)
		block["time"] = unixTimeUTC.Format(time.RFC850)
		block["timeutc"] = unixTimeUTC.Format(time.RFC3339)
	}
	return block
}
func (b *BitNode) GetBlockShortByHeight(blockheight int) (block map[string]interface{}) {
	bparams := []int{blockheight}
	blockHash, err := b.Jrc.MakeRequest("getblockhash", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Block By Height Error", err)
	}
	if blockHash != nil {
		block = b.GetBlockShort((blockHash).(string))
	}
	return block
}

func (b *BitNode) GetBlockByHeight(blockheight int) (block map[string]interface{}) {

	bparams := []int{blockheight}
	blockHash, err := b.Jrc.MakeRequest("getblockhash", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Block By Height Error", err)
	}
	if blockHash != nil {
		block = b.APIGetBlock((blockHash).(string))
	}
	return block
}
