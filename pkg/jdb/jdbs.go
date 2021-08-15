package jdb

import (
	"net/http"
	"sync"

	"github.com/comhttp/jdbc"
)

//// New creates a new scribble database at the desired directory location, and
// returns a *Driver to then use for interacting with the database
func NewJDBS(jbds map[string]string) *JDBS {
	j := &JDBS{
		B: make(map[string]*JDB),
	}
	headers := http.Header{}
	auth := ""
	if auth != "" {
		headers.Add("Authorization", "Bearer "+auth)
	}
	for d, url := range jbds {
		client, err := jdbc.NewClient("http://"+url, jdbc.ClientOptions{Headers: headers})
		if err != nil {
		}
		//j.B[d] = &JDB{
		//	mutexes: make(map[string]*sync.Mutex),
		//	client:  client,
		//	delimiter: "_",
		//}
		j.B[d] = NewJDB(client, "_")
		//defer j.B[d].client.Close()
	}
	return j
}

func NewJDB(client *jdbc.Client, delimiter string) *JDB {
	return &JDB{
		mutexes:   make(map[string]*sync.Mutex),
		client:    client,
		delimiter: delimiter,
	}
}
