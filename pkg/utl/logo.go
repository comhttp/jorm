package utl

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
)

type Logo struct {
	Slug      string `json:"slug"`
	Extension string `json:"ext"`
	Data      string `json:"data"`
}

func (l *Logo) SetLogo(url string) {
	if url != "" && url != "missing_large.png" {
		res, err := http.Get(url)
		if err != nil {
			log.Print("Problem Insert", err)
		}
		defer res.Body.Close()
		logo, err := ioutil.ReadAll(res.Body)
		img256, _ := imageResize(logo, options{Width: 256, Height: 256})

		l.Data = hex.EncodeToString(img256)
	}
	return
}
