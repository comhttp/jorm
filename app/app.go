package app

import (
	"net/http"
)

type JORM struct {
	//Hosts         map[string]Host
	Server *http.Server
}
