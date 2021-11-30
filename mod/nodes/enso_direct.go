package nodes

import (
	"fmt"

	"github.com/comhttp/jorm/pkg/cfg"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/gorilla/mux"
)

func ENSOroutesDirect(path string, rpcLogin cfg.RPClogin, r *mux.Router) *mux.Router {
	//info := Queries(j, "info")
	//info := Queries(j.JDBclient("explorer"),"info")
	//r.StrictSlash(true)
	//n := s.PathPrefix("/n").Subrouter()
	//n.HandleFunc("/{coin}/nodes", explorersCollection.CoinNodesHandler).Methods("GET")
	//n.HandleFunc("/{coin}/{nodeip}", explorersCollection.nodeHandler).Methods("GET")
	bsRPC := bitNodesRPC{
		path:     path,
		username: rpcLogin.Username,
		password: rpcLogin.Password,
		// bitNodes: bs,
	}

	// coinBitNodes := BitNodes{}
	// c, _ := cfg.NewCFG(path, nil)
	// err := c.Read("nodes", "parallelcoin", &coinBitNodes)
	// utl.ErrorLog(err)
	// fmt.Println("coinBitNodescoinBitNodescoinBitNodescoinBitNodes:  ", coinBitNodes)

	b := r.PathPrefix("/direct").Subrouter()
	// b.HandleFunc("/{coin}/status", ViewStatus).Methods("GET")
	b.HandleFunc("/{coin}/blocks/{per}/{page}", bsRPC.directViewBlocks).Methods("GET")
	// b.HandleFunc("/{coin}/lastblock", bsRPC.directLastBlock).Methods("GET")
	// b.HandleFunc("/{coin}/block/{id}", bsRPC.directViewBlock).Methods("GET")
	// b.HandleFunc("/{coin}/tx/{txid}", bsRPC.directViewTx).Methods("GET")
	// b.HandleFunc("/{coin}/mempool", bsRPC.directViewRawMemPool).Methods("GET")
	// b.HandleFunc("/{coin}/mining", bsRPC.directViewMiningInfo).Methods("GET")
	// b.HandleFunc("/{coin}/info", bsRPC.directViewInfo).Methods("GET")
	// b.HandleFunc("/{coin}/peers", bsRPC.directViewPeers).Methods("GET")
	//b.HandleFunc("/{coin}/market", explorersCollection.ViewMarket).Methods("GET")
	return r
}

func RPCSRC(bitNodes BitNodes, username, password string) (b *BitNode) {
	for _, bn := range bitNodes {

		fmt.Println("bnbnbnbn:", bn)

		// fmt.Println("path:   ", b.path)
		// fmt.Println("coin::  ", v["coin"])
		fmt.Println("username:  ", username)
		fmt.Println("password:  ", password)
		fmt.Println("IPIPIP :    ", bn.IP)
		fmt.Println("PortPortPort :    ", bn.Port)
		fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
		// fmt.Println("bitNodescoin:  ", bitNodes(b.path, v["coin"]))

		b.Jrc = utl.NewClient(username, password, bn.IP, bn.Port)

		fmt.Println("JrcJrcJrcJrcJrcJrc :    ", b.Jrc)

		// fmt.Println("bitNodeSRC:", b)
		// if b.Jrc != nil {
		// 	break
		// }
		// return
	}

	return
}

func bitNodes(path, coin string) (bitNodes BitNodes) {
	c, _ := cfg.NewCFG(path, nil)
	err := c.Read("nodes", coin, &bitNodes)
	utl.ErrorLog(err)
	return
}
