package nodes

import (
	"fmt"
)

func (b *BitNode) APIGetRawMemPool() interface{} {
	bparams := []int{}
	get, err := b.Jrc.MakeRequest("getrawmempool", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Raw Mem Pool Error", err)
	}
	return get
}

func (b *BitNode) APIGetMiningInfo() interface{} {
	bparams := []int{}
	get, err := b.Jrc.MakeRequest("getmininginfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Mining Info Error", err)
	}
	return get
}

func (b *BitNode) APIGetNetworkInfo() interface{} {
	bparams := []int{}
	get, err := b.Jrc.MakeRequest("getnetworkinfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Network Info Error", err)
	}
	return get
}

func (b *BitNode) APIGetInfo() interface{} {
	bparams := []int{}
	get, err := b.Jrc.MakeRequest("getinfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Info Error", err)
	}
	return get
}

func (b *BitNode) APIGetPeerInfo() interface{} {
	bparams := []int{}
	get, err := b.Jrc.MakeRequest("getpeerinfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Peer Info Error", err)
	}
	return get
}

func (b *BitNode) addNode(ip string) interface{} {
	bparams := []string{ip, "add"}
	get, err := b.Jrc.MakeRequest("addnode", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Peer Info Error", err)
	}
	return get
}

func (b *BitNode) APIGetAddNodeInfo(ip string) interface{} {
	bparams := []int{}
	get, err := b.Jrc.MakeRequest("getaddednodeinfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Peer Info Error", err)
	}
	return get
}
