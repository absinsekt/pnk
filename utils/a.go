package utils

import (
	"log"
)

// Log logs to preferred channel using preferred logger
func Log(msg interface{}) {
	log.Println(msg)
}

// Logf logs formatted messages to preferred channel using preferred logger
func Logf(format string, msgs ...interface{}) {
	log.Printf(format+"\n", msgs...)
}

// LogFatal logs to preferred channel using preferred logger ans exits with error code
func LogFatal(msg interface{}) {
	log.Fatalln(msg)
}

// Check checks multiple err != nil mantras & logs if something went wrong
func Check(e error, isFatal bool) {
	if e != nil {
		if isFatal {
			LogFatal(e.Error())
		} else {
			Log(e.Error())
		}
	}
}
