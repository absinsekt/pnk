package core

import (
	"os"
	"strconv"

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

// GetEnv returns ENV variable value or fallback if not set
func GetEnv(key string, fallback interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		switch fallback.(type) {
		case int:
			if v, err := strconv.Atoi(value); err == nil {
				return v
			}
		case bool:
			if v, err := strconv.ParseBool(value); err == nil {
				return v
			}
		default:
			return value
		}
	}

	return fallback
}
