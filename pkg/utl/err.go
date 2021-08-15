package utl

import (
	"github.com/rs/zerolog/log"
)

func ErrorLog(err error) {
	if err != nil {
		log.Print("Error", err)
	}
}
