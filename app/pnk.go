package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/controllers"
	"github.com/absinsekt/pnk/models"
)

func main() {
	initLogger()

	logrus.Info("Pnk is starting")

	models.CheckConnection()

	addr := fmt.Sprintf("%s:%d", configuration.HostAddress, configuration.Port)

	srv := &fasthttp.Server{
		Name:              "pnk",
		Handler:           controllers.NewRouter(),
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		ReduceMemoryUsage: true,
	}

	logrus.Infof("Listening on http://%s/", addr)
	logrus.Fatalln(srv.ListenAndServe(addr))
}

func initLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	})

	if configuration.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
