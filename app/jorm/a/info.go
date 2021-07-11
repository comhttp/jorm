package a

import (
	"fmt"
	cfg2 "github.com/comhttp/jorm/cfg"

	"github.com/comhttp/jorm/pkg/utl"
)

func (rpc *BitNode) GetRawMemPool() interface{} {
	jrc := utl.NewClient(cfg2.C.RPC.Username, cfg2.C.RPC.Password, rpc.IP, rpc.Port)
	if jrc == nil {
		fmt.Println("Error n status write")
	}
	bparams := []int{}
	get, err := jrc.MakeRequest("getrawmempool", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Raw Mem Pool Error", err)
	}
	return get
}

func (rpc *BitNode) GetMiningInfo() interface{} {
	jrc := utl.NewClient(cfg2.C.RPC.Username, cfg2.C.RPC.Password, rpc.IP, rpc.Port)
	if jrc == nil {
		fmt.Println("Error n status write")
	}
	bparams := []int{}
	get, err := jrc.MakeRequest("getmininginfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Mining Info Error", err)
	}
	return get
}

func (rpc *BitNode) GetNetworkInfo() interface{} {
	bparams := []int{}
	get, err := rpc.Jrc.MakeRequest("getnetworkinfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Network Info Error", err)
	}
	return get
}

func (rpc *BitNode) GetInfo() interface{} {
	bparams := []int{}
	get, err := rpc.Jrc.MakeRequest("getinfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Info Error", err)
	}
	return get
}

func (rpc *BitNode) GetPeerInfo() interface{} {
	bparams := []int{}
	get, err := rpc.Jrc.MakeRequest("getpeerinfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Peer Info Error", err)
	}
	return get
}

func (rpc *BitNode) addNode(ip string) interface{} {
	bparams := []string{ip, "add"}
	get, err := rpc.Jrc.MakeRequest("addnode", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Peer Info Error", err)
	}
	return get
}

func (rpc *BitNode) GetAddNodeInfo(ip string) interface{} {
	bparams := []int{}
	get, err := rpc.Jrc.MakeRequest("getaddednodeinfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Peer Info Error", err)
	}
	return get
}
