package jdb

import (
	"sync"

	"github.com/comhttp/jdbc"
)

type (
	JDB struct {
		mutex     sync.Mutex
		mutexes   map[string]*sync.Mutex
		path      string
		client    *jdbc.Client
		delimiter string
	}
	Collection struct {
		j   *JDB
		col string
	}
)
