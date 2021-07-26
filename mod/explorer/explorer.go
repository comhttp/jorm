package explorer

import (
	"github.com/comhttp/jorm/pkg/jdb"
)

type Explorer struct {
	Status map[string]*Blockchain `json:"status"`
	j      *jdb.JDB
}

type Blockchain struct {
	Blocks    int `json:"blocks"`
	Txs       int `json:"txs"`
	Addresses int `json:"addresses"`
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
