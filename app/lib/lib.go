package lib

import (
	log "github.com/sirupsen/logrus"
)

// Check checks multiple err != nil mantras & logs if something went wrong
func Check(e error, isFatal bool) {
	if e != nil {
		if isFatal {
			log.Fatalln(e.Error())
		} else {
			log.Warnln(e.Error())
		}
	}
}
