package jdb

import (
	"fmt"
	"strings"
	"sync"

	"github.com/rs/zerolog/log"
)

//func (j *JDB)CollectionQueries(collection string) *CollectionQueries {
//	return &CollectionQueries{
//		j:j,
//		col:collection,
//	}
//}

// Write locks the database and attempts to write the record to the database under
// the [collection] specified with the [resource] name given
func (j *JDB) Write(collection, key string, v interface{}) error {
	// ensure there is a place to save record
	if collection == "" {
		return fmt.Errorf("Missing collection - no place to save record!")
	}
	// ensure there is a resource (name) to save record as
	if key == "" {
		return fmt.Errorf("Missing resource - unable to save record (no name)!")
	}
	mutex := j.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()
	// move final file into place
	return j.client.SetJSON(collection+j.delimiter+key, v)
}

// Write locks the database and attempts to write the record to the database under
// the [collection] specified with the [resource] name given
func (j *JDB) WriteAll(data map[string]interface{}) error {
	mutex := j.getOrCreateMutex("all")
	mutex.Lock()
	defer mutex.Unlock()
	// move final file into place
	return j.client.SetJSONs(data)
}

// Read a record from the database
func (j *JDB) Read(collection, key string, v interface{}) error {
	//// ensure there is a place to save record
	//if db == "" {
	//	return fmt.Errorf("Missing collection - no place to save record!")
	//}
	// ensure there is a resource (name) to save record as
	if key == "" {
		return fmt.Errorf("Missing resource - unable to save record (no name)!")
	}
	// unmarshal data
	//return json.Unmarshal(b, &v)
	//fmt.Println("00001111",key)
	if key == "status" {
		//spew.Dump(j.client)
		//fmt.Println("00001111clientclientclientclient",j.client)
	}

	return j.client.GetJSON(collection+j.delimiter+key, &v)
}

// ReadAll records from a collection; this is returned as a slice of strings because
// there is no way of knowing what type the record is.
func (j *JDB) ReadAll(collection string) ([]string, error) {
	var all []string
	al, err := j.client.GetByPrefix(collection + j.delimiter)
	for a, _ := range al {
		all = append(all, strings.TrimPrefix(a, collection+j.delimiter))
	}
	return all, err
}

func (j *JDB) ReadAllPerPages(collection string, per, page int) (p []map[string]interface{}) {
	records, err := j.ReadAll(collection)
	log.Err(err)
	recordCount := len(records)
	startRecord := recordCount - per*page
	minusBlockStart := int(startRecord + per)
	for _, record := range records {
		fmt.Println("END cryptocompare COINS:::::::::::::::::::::::::::::: ", record)
		r := make(map[string]interface{})
		err := j.Read(collection, record, &r)
		log.Err(err)
		//fmt.Println("<RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRS>:::::::::::::::::::::::::::::: ",r)
		if ibh := minusBlockStart; ibh >= startRecord {
			p = append(p, r)
			ibh--
		}
		//for ibh := minusBlockStart; ibh >= startRecord; {
		//p = append(p, record)
		////p = append(p, GetBlockShort(j, coin, strconv.Itoa(ibh)))
		//ibh--
	}
	//sort.SliceStable(p, func(i, j int) bool {
	//	return int(p[i]["order"].(int64)) < int(p[j]["order"].(int64))
	//})
	return p
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

//func (j *JDBS)CollectionQueries(db,col string) interface{} {
//	return &Collection{
//		j:   j.B[db],
//		col: col,
//	}
//}
