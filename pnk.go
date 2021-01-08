package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/absinsekt/pnk/lib"
	"github.com/absinsekt/pnk/lib/configuration"
	"github.com/absinsekt/pnk/models"

	"github.com/absinsekt/pnk/controllers/admin"
	"github.com/absinsekt/pnk/controllers/api"
	"github.com/absinsekt/pnk/controllers/www"
)

func main() {
	initLogger()

	logrus.Info("Pnk is starting")

	models.CheckConnection()

	addr := fmt.Sprintf("%s:%d", configuration.HostAddress, configuration.Port)

	router := lib.NewRouter([]lib.Mountable{
		&admin.Routes{},
		&www.Routes{},
		&api.Routes{},
	})

	srv := &fasthttp.Server{
		Name:              "pnk",
		Handler:           router.Handler,
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
