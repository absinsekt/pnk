package main

// import (
// 	"fmt"
// 	"time"

// 	log "github.com/sirupsen/logrus"
// 	"github.com/valyala/fasthttp"

// 	"github.com/absinsekt/pnk/lib"
// 	"github.com/absinsekt/pnk/lib/core"
// 	"github.com/absinsekt/pnk/lib/db"
// 	"github.com/absinsekt/pnk/lib/templateset"

// 	"github.com/absinsekt/pnk/controllers/admin"
// 	"github.com/absinsekt/pnk/controllers/api"
// 	"github.com/absinsekt/pnk/controllers/www"
// )

// func main() {
// 	core.InitConfiguration(nil)
// 	core.InitLogger()

// 	log.Info("Pnk is starting")

// 	db.InitConnection()
// 	templateset.InitTemplateSet()

// 	addr := fmt.Sprintf("%s:%d", core.Config.HostAddress, core.Config.Port)

// 	router := lib.NewRouter([]lib.Mountable{
// 		&admin.Routes{},
// 		&www.Routes{},
// 		&api.Routes{},
// 	})

// 	srv := &fasthttp.Server{
// 		Name:              "pnk",
// 		Handler:           router.Handler,
// 		ReadTimeout:       15 * time.Second,
// 		WriteTimeout:      15 * time.Second,
// 		ReduceMemoryUsage: true,
// 	}

// 	log.Infof("Listening on http://%s/", addr)
// 	log.Fatalln(srv.ListenAndServe(addr))
// }
