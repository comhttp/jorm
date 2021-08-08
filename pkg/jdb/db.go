package jdb

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/comhttp/jdbc"
)

//var JDB, _ = NewJDB(cfg.Path, nil)

type (
	JDB struct {
		col     string
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		path    string
		clients map[string]*jdbc.Client
	}
)

// New creates a new scribble database at the desired directory location, and
// returns a *Driver to then use for interacting with the database
func NewJDB(jbds map[string]string) *JDB {
	j := JDB{
		mutexes: make(map[string]*sync.Mutex),
		clients: make(map[string]*jdbc.Client),
	}
	for js, url := range jbds {
		j.cl(js, url)
	}
	return &j
}

func (j *JDB) cl(js, url string) {
	headers := http.Header{}
	auth := ""
	if auth != "" {
		headers.Add("Authorization", "Bearer "+auth)
	}
	client, err := jdbc.NewClient("http://"+url, jdbc.ClientOptions{Headers: headers})
	if err != nil {
	}
	j.clients[js] = client
	//defer j.clients[js].Close()
}

// Write locks the database and attempts to write the record to the database under
// the [collection] specified with the [resource] name given
func (j *JDB) Write(db, key string, v interface{}) error {
	// ensure there is a place to save record
	if db == "" {
		return fmt.Errorf("Missing collection - no place to save record!")
	}
	// ensure there is a resource (name) to save record as
	if key == "" {
		return fmt.Errorf("Missing resource - unable to save record (no name)!")
	}
	mutex := j.getOrCreateMutex(db)
	mutex.Lock()
	defer mutex.Unlock()
	// move final file into place
	return j.clients[db].SetJSON(key, v)
}

// Read a record from the database
func (j *JDB) Read(db, key string, v interface{}) error {
	// ensure there is a place to save record
	if db == "" {
		return fmt.Errorf("Missing collection - no place to save record!")
	}
	// ensure there is a resource (name) to save record as
	if key == "" {
		return fmt.Errorf("Missing resource - unable to save record (no name)!")
	}
	// unmarshal data
	//return json.Unmarshal(b, &v)
	return j.clients[db].GetJSON(key, &v)
}

// ReadAll records from a collection; this is returned as a slice of strings because
// there is no way of knowing what type the record is.
func (j *JDB) ReadAll(db, prefix string) ([]string, error) {
	var all []string
	al, err := j.clients[db].GetByPrefix(prefix)
	for a, _ := range al {
		all = append(all, a)
	}
	return all, err
}

// Delete locks that database and then attempts to remove the collection/resource
// specified by [path]
func (j *JDB) Delete(collection, resource string) error {
	return nil
}

// getOrCreateMutex creates a new collection specific mutex any time a collection
// is being modfied to avoid unsafe operations
func (j *JDB) getOrCreateMutex(collection string) *sync.Mutex {
	j.mutex.Lock()
	defer j.mutex.Unlock()
	m, ok := j.mutexes[collection]
	// if the mutex doesn't exist make it
	if !ok {
		m = &sync.Mutex{}
		j.mutexes[collection] = m
	}
	return m
}
