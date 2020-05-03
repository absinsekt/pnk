package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/controllers"
	"github.com/absinsekt/pnk/models"
)

func main() {
	initLogger()
	log.Info("Pnk is starting")

	models.CheckConnection()

	addr := fmt.Sprintf("%s:%d", configuration.HostAddress, configuration.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      controllers.NewControllersRouter(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Infof("Listening on http://%s/", addr)
	log.Fatalln(srv.ListenAndServe())
}

func initLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	})

	if configuration.Debug {
		log.SetLevel(log.DebugLevel)
	}
}
