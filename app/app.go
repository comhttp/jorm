package app

import (
	"crypto/tls"
	coins2 "github.com/comhttp/jorm/coins"
	jdb2 "github.com/comhttp/jorm/jdb"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
)

type JORM struct {
	Coins coins2.Coins
	//Hosts         map[string]Host
	WWW         *http.Server
	WS          *http.Server
	TLSconfig   *tls.Config
	CertManager autocert.Manager
	JDB         *jdb2.JDB
}
