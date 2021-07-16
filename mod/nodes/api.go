package nodes

import (
	"fmt"
)

func (b *BitNode) APIGetRawMemPool() interface{} {
	//jrc := utl.NewClient(cfg.C.RPC.Username, cfg.C.RPC.Password, rpc.IP, rpc.Port)
	//if jrc == nil {
	//	fmt.Println("Error n status write")
	//}
	bparams := []int{}
	get, err := b.Jrc.MakeRequest("getrawmempool", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Raw Mem Pool Error", err)
	}
	return get
}

func (b *BitNode) APIGetMiningInfo() interface{} {
	//jrc := utl.NewClient(cfg.C.RPC.Username, cfg.C.RPC.Password, rpc.IP, rpc.Port)
	//if jrc == nil {
	//	fmt.Println("Error n status write")
	//}
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

// BitNodeSrc is a node's address
//type BitNode struct {
//	IP   string `json:"ip"`
//	Port int64  `json:"p"`
//	Jrc  *utl.Endpoint
//}
//
//func RPCSRC(c string) (b *BitNode) {
//	bitNodes := BitNodes{}
//	if err := cfg.CFG.Read("nodes", c, &bitNodes); err != nil {
//		fmt.Println("Errdor", err)
//	}
//	for _, bn := range bitNodes {
//		b = &bn
//		b.Jrc = utl.NewClient(cfg.C.RPC.Username, cfg.C.RPC.Password, b.IP, b.Port)
//		if b.Jrc != nil {
//			fmt.Println("bitb:", b.IP)
//			break
//		}
//	}
//	fmt.Println("b:", b)
//	return
//}
