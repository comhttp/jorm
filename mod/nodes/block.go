package nodes

import (
	"fmt"
	"strconv"
)

func (b *BitNode) APIGetBlockCount() (blockcount int) {
	bparams := []int{}
	gbc, err := b.Jrc.MakeRequest("getblockcount", bparams)
	if err != nil {
		fmt.Println("Error n call: ", err)
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

func (b *BitNode) APIGetBlock(blockhash string) (block interface{}) {
	bparams := []string{blockhash}
	block, err := b.Jrc.MakeRequest("getblock", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Block Error", err)
	}
	return
}

func (b *BitNode) APIGetBlockByHeight(blockheight int) (block interface{}) {
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

func (b *BitNode) APIGetTx(txid string) (t interface{}) {
	verbose := int(1)
	var grtx []interface{}
	grtx = append(grtx, txid)
	grtx = append(grtx, verbose)
	t, err := b.Jrc.MakeRequest("getrawtransaction", grtx)
	if err != nil {
		fmt.Println("Jorm Node Get Tx Error", err)
	}
	return
}
