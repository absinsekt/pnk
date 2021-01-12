package core

import (
	log "github.com/sirupsen/logrus"
)

// InitLogger sets debug level and logger preferences
func InitLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	})

	if Config.Debug {
		log.SetLevel(log.DebugLevel)
	}
}
