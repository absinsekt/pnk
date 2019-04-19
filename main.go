package main

import (
	"fmt"
	"github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/controllers"
	"github.com/absinsekt/pnk/utils"
	"net/http"
	"time"
)

var (
	addr = fmt.Sprintf("%s:%d", configuration.HostAddress, configuration.Port)

	srv = &http.Server{
		Addr:         addr,
		Handler:      controllers.NewControllersRouter(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
)

func main() {
	utils.Logf("Pnk is listening on http://%s/", addr)
	utils.LogFatal(srv.ListenAndServe())
}
