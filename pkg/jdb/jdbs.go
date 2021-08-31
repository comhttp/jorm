package jdb

import (
	"github.com/comhttp/jorm/pkg/utl"
	"net/http"
	"sync"

	"github.com/comhttp/jdbc"
)

var (
	headers   = http.Header{}
	delimiter = "_"
)

func NewJDB(url string) (*JDB, error) {
	auth := ""
	if auth != "" {
		headers.Add("Authorization", "Bearer "+auth)
	}
	client, err := jdbc.NewClient("http://"+url, jdbc.ClientOptions{Headers: headers})
	utl.ErrorLog(err)
	return &JDB{
		mutexes:   make(map[string]*sync.Mutex),
		client:    client,
		delimiter: delimiter,
	}, err
}
