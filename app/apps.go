package app

import (
	"github.com/comhttp/jorm/mod/cloudflare"
	"github.com/comhttp/jorm/mod/coin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (j *JORM) srvJORM(service, path, singleCoin string) {
	switch service {
	case "proxy":
		log.Print("reverse proxy")
		h := &baseHandle{}
		http.Handle("/", h)
		j.WWW.Handler = h
		j.WWW.Addr = ":" + j.config.Port["proxy"]
		log.Fatal().Err(j.WWW.ListenAndServe())
	case "jorm":
		log.Print("jorm")
		j.setExplorers()
		j.JormSRV()
		j.WWW.Handler = j.JORMhandlers()
		j.WWW.Addr = ":" + j.config.Port["jorm"]
		log.Fatal().Err(j.WWW.ListenAndServe())
	case "enso":
		log.Print("enso")
		j.setExplorers()
		j.WWW.Handler = j.ENSOhandlers()
		j.WWW.Addr = ":" + j.config.Port["enso"]
		log.Fatal().Err(j.WWW.ListenAndServe())
	case "our":
		log.Print("our")
		j.WWW.Handler = j.OURhandlers()
		j.WWW.Addr = ":" + j.config.Port["our"]
		log.Fatal().Err(j.WWW.ListenAndServe())
	//ourSRV()
	case "comhttp":
		log.Print("comhttp")
		j.WWW.Handler = j.COMHTTPhandlers()
		j.WWW.Addr = ":" + j.config.Port["comhttp"]
		log.Fatal().Err(j.WWW.ListenAndServe())
	case "admin":
		log.Print("admin")
		j.WWW.Handler = j.ADMINhandlers()
		j.WWW.Addr = ":" + j.config.Port["admin"]
		log.Fatal().Err(j.WWW.ListenAndServe())
	case "explorer":
		//log.Print("explorer " + *coin)
		if path == "" {
			j.ExplorerSRV(singleCoin)
			j.WWW.Addr = ":" + j.config.Port[singleCoin]
		} else {
			log.Print("Missing coin for explorer!")
		}
		log.Fatal().Err(j.WWW.ListenAndServe())
	case "cloudflare":
		log.Print("cloudflare")
		cloudflare.CloudFlare(j.config, &coin.CoinsShort{})
		return
	}
	return
}
