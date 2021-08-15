package jdb

import (
	"sync"

	"github.com/comhttp/jdbc"
)

//var JDB, _ = NewJDB(cfg.Path, nil)

type (
	JDBS struct {
		B         map[string]*JDB
		delimiter string
	}

	JDB struct {
		col       string
		mutex     sync.Mutex
		mutexes   map[string]*sync.Mutex
		path      string
		client    *jdbc.Client
		delimiter string
	}
)
