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
