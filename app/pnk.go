package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/controllers"
	"github.com/absinsekt/pnk/models"
)

func main() {
	initLogger()

	log.Info("Pnk is starting")

	models.CheckConnection()

	addr := fmt.Sprintf("%s:%d", configuration.HostAddress, configuration.Port)

	srv := &fasthttp.Server{
		Name:              "pnk",
		Handler:           controllers.NewRouter(),
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		ReduceMemoryUsage: true,
	}

	log.Infof("Listening on http://%s/", addr)
	log.Fatalln(srv.ListenAndServe(addr))
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
