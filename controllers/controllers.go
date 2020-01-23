package controllers

import (
	"net/http"
	"net/http/httputil"

	"github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/utils"
	"github.com/gorilla/mux"
)

var templateSet *utils.TemplateSet

// NewControllersRouter mounts all paths, e. g. apis and static page handlers
// and returns root Router
func NewControllersRouter() *mux.Router {
	var err error

	templateSet, err = utils.NewTemplateSet(configuration.TemplatePath)
	utils.Check(err, true)

	root := mux.NewRouter()
	root.HandleFunc("/", handleIndex)

	if configuration.Debug {
		// for debug only
		root.PathPrefix("/dist").Handler(http.StripPrefix("/dist", http.FileServer(http.Dir("static/dist"))))

		// for debug only
		proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
			req.Host = "127.0.0.1:9001"
			req.URL.Scheme = "http"
			req.URL.Host = "127.0.0.1:9001"
		}}

		root.PathPrefix("/sockjs-node").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			proxy.ServeHTTP(w, r)
		})
	}

	return root
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	templateSet.Render("index.html", res, nil)
}
