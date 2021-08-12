package utl

import (
	"log"
)

func ErrorLog(err error) {
	if err != nil {
		log.Println("Error", err)
	}
}
