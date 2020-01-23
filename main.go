package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/controllers"
)

func main() {
	initLogger()

	addr := fmt.Sprintf("%s:%d", configuration.HostAddress, configuration.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      controllers.NewControllersRouter(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Infof("Pnk is listening on http://%s/", addr)
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
