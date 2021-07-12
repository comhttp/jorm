package explorer

import (
	"github.com/comhttp/jorm/jdb"
	"github.com/comhttp/jorm/pkg/utl"
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
