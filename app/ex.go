package app

import (
	"github.com/comhttp/jorm/mod/nodes"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
)

//type Explorer struct {
//	Status map[string]*BlockchainStatus `json:"status"`
//}

type Explorer struct {
	Coin     string
	BitNodes nodes.BitNodes
	status   *BlockchainStatus
}
type Explorers map[string]*Explorer

func NewExplorer(j *jdb.JDB, coin string) *Explorer {
	//log.SetLevel(parseLogLevel(loglevel))
	bitNodes := nodes.BitNodes{}
	//if err := app.CFG.Read("nodes", coin, &bitNodes); err != nil {
	//	log.Println("Error", err)
	//}
	e := &Explorer{
		Coin:     coin,
		BitNodes: bitNodes,
		//JDB:      jdb.NewJDB(app.C.JDBservers),
	}
	err := j.Read(coin, "status", &e.status)
	if e.status == nil {
		e.status = &BlockchainStatus{
			Blocks:    0,
			Txs:       0,
			Addresses: 0,
		}
	}
	utl.ErrorLog(err)
	return e
}
