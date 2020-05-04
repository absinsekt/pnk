package controllers

import (
	"net/http"
	"net/http/httputil"

	"github.com/absinsekt/pnk/controllers/admin"
	"github.com/absinsekt/pnk/controllers/api"
	"github.com/absinsekt/pnk/controllers/www"
	"github.com/absinsekt/pnk/utils"
	"github.com/absinsekt/pnk/utils/core"

	"github.com/absinsekt/pnk/configuration"
	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

var templateSet *core.TemplateSet

// NewControllersRouter mounts all paths, e. g. apis and static page handlers
// and returns root Router
func NewControllersRouter() *mux.Router {
	var err error

	// TODO migrate to main pnk.go
	templateSet, err = core.NewTemplateSet(configuration.TemplatePath)
	utils.Check(err, true)

	root := mux.NewRouter().StrictSlash(true)

	// mount all subpaths
	www.MountIndex(root, templateSet)
	admin.MountAdmin(root, templateSet)
	api.MountAuth(root)

	if configuration.Debug {
		// proxy sockjs for hmr
		proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
			req.Host = "127.0.0.1:5001"
			req.URL.Scheme = "http"
			req.URL.Host = "127.0.0.1:5001"
		}}

		root.PathPrefix("/sockjs-node").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			proxy.ServeHTTP(w, r)
		})

		// serve static
		root.PathPrefix("/dist").Handler(http.StripPrefix("/dist", http.FileServer(http.Dir("../www/dist"))))
		root.PathPrefix("/").Handler(http.FileServer(http.Dir("../www/dist")))
	}

	// default handler
	root.PathPrefix("/").HandlerFunc(handleDefault)

	return root
}

func handleDefault(res http.ResponseWriter, req *http.Request) {
	log.Errorf("%s URI isn't described, falling back to default", req.RequestURI)
	res.WriteHeader(http.StatusTeapot)
}
