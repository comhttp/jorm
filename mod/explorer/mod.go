package explorer

import (
	"github.com/comhttp/jorm/mod/nodes"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
)

//type Explorer struct {
//	Status map[string]*BlockchainStatus `json:"status"`
//}
type ExplorerQueries struct {
	j      *jdb.JDB
	status *BlockchainStatus
	col    string
}
type Explorer struct {
	Coin     string
	BitNodes nodes.BitNodes
	Status   *BlockchainStatus
}
type Explorers map[string]*Explorer

type BlockchainStatus struct {
	Blocks    int `json:"blocks"`
	Txs       int `json:"txs"`
	Addresses int `json:"addresses"`
}

func (eq *ExplorerQueries) NewExplorer(coin string) *Explorer {
	//log.SetLevel(parseLogLevel(loglevel))
	bitNodes := nodes.BitNodes{}
	//if err := app.CFG.Read("nodes", coin, &bitNodes); err != nil {
	//	log.Print("Error", err)
	//}
	e := &Explorer{
		Coin:     coin,
		BitNodes: bitNodes,
	}
	//eq := Queries(j,"info")
	status, err := eq.GetStatus(coin)
	utl.ErrorLog(err)
	e.Status = status
	if e.Status == nil {
		e.Status = &BlockchainStatus{
			Blocks:    0,
			Txs:       0,
			Addresses: 0,
		}
	}
	return e
}
