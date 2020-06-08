package main

import (
	"fmt"

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
	// srv := &http.Server{
	// 	Addr:         addr,
	// 	Handler:      controllers.NewControllersRouter(),
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	srv := &fasthttp.Server{
		Handler: controllers.TestHandler,
	}

	log.Infof("Listening on http://%s/", addr)
	// log.Fatalln(srv.ListenAndServe())

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
