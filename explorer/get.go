package explorer

import (
	"github.com/comhttp/jorm/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"strconv"
)

func GetExplorer(j *jdb.JDB) Explorer {
	e := Explorer{}
	err := j.Read("info", "explorer", &e)
	utl.ErrorLog(err)
	return e
}

func GetIndex(j *jdb.JDB, c, t string) map[int]string {
	index := map[int]string{}
	err := j.Read(c, t, &index)
	utl.ErrorLog(err)
	return index
}

func GetBlock(j *jdb.JDB, c, id string) map[string]interface{} {
	blockHash := ""
	block := make(map[string]interface{})
	_, err := strconv.Atoi(id)
	if err != nil {
		blockHash = id
	} else {
		blockHash = ""
		err = j.Read(c, "block_"+id, &blockHash)
	}
	err = j.Read(c, "block_"+blockHash, &block)
	utl.ErrorLog(err)
	return block
}

func GetTx(j *jdb.JDB, c, id string) map[string]interface{} {
	tx := make(map[string]interface{})
	err := j.Read(c, "tx_"+id, &tx)
	utl.ErrorLog(err)
	return tx
}
func GetAddr(j *jdb.JDB, c, id string) map[string]interface{} {
	addr := make(map[string]interface{})
	err := j.Read(c, "addr_"+id, &addr)
	utl.ErrorLog(err)
	return addr
}
