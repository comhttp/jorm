package app

import (
	"crypto/tls"
	"github.com/comhttp/jorm/app/jorm/coin"
	"net/http"
)

type JORM struct {
	Coins coin.Coins
	//Hosts         map[string]Host
	Server    *http.Server
	TLSconfig *tls.Config
}
