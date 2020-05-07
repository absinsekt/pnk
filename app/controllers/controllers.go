package controllers

import (
	"net/http"
	"net/http/httputil"

	"github.com/absinsekt/pnk/controllers/admin"
	"github.com/absinsekt/pnk/controllers/api"
	"github.com/absinsekt/pnk/controllers/www"
	"github.com/absinsekt/pnk/utils"
	"github.com/absinsekt/pnk/utils/responses"
	"github.com/absinsekt/pnk/utils/templateset"

	"github.com/absinsekt/pnk/configuration"
	"github.com/gorilla/mux"
)

var templateSet *templateset.TemplateSet

// NewControllersRouter mounts all paths, e. g. apis and static page handlers
// and returns root Router
func NewControllersRouter() *mux.Router {
	var err error

	templateSet, err = templateset.NewTemplateSet(configuration.TemplatePath)
	utils.Check(err, true)

	root := mux.NewRouter().StrictSlash(true)

	// mount all subpaths
	www.Mount(root, templateSet)
	admin.Mount(root, templateSet)
	api.Mount(root)

	if configuration.Debug {
		// proxy sockjs for hmr
		proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
			req.Host = "127.0.0.1:5001"
			req.URL.Scheme = "http"
			req.URL.Host = "127.0.0.1:5001"
		}}

		root.PathPrefix("/sockjs-node").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			proxy.ServeHTTP(res, req)
		})

		// serve static
		root.PathPrefix("/dist").Handler(http.StripPrefix("/dist", http.FileServer(http.Dir("../www/dist"))))
	}

	// default handler
	root.PathPrefix("/").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		responses.ErrorResponse(res, req, http.StatusNotFound, templateSet)
	})

	return root
}
