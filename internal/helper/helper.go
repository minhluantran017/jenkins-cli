package helper

import (
	log "github.com/sirupsen/logrus"
)

func HandleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
