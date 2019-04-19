package controllers

import (
	"github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/utils"
	"github.com/flosch/pongo2"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httputil"
	"time"
)

var (
	tmplBaseDir    = "templates"
	wwwTemplateSet *pongo2.TemplateSet
)

// NewControllersRouter mounts all paths, e. g. apis and static page handlers
// and returns root Router
func NewControllersRouter() *mux.Router {
	wwwFSLoader, err := pongo2.NewSandboxedFilesystemLoader(tmplBaseDir)
	utils.Check(err, true)

	wwwTemplateSet = pongo2.NewSet("www", wwwFSLoader)
	wwwTemplateSet.Debug = configuration.Debug

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

func renderTemplate(templateName string, context pongo2.Context, res http.ResponseWriter) {
	var (
		tmpl *pongo2.Template
		err  error
	)

	timerStart := time.Now()

	// if wwwTemplateSet.Debug {
	// 	tmpl, err = wwwTemplateSet.FromFile(templateName)
	// } else {
	tmpl, err = wwwTemplateSet.FromCache(templateName)
	// }

	if err != nil {
		// disable if no debug
		_, _ = res.Write([]byte(err.Error()))
	} else {
		if wwwTemplateSet.Debug {
			utils.Logf(
				"Template %s/%s rendered in %.2fms",
				tmplBaseDir,
				templateName,
				time.Now().Sub(timerStart).Seconds()*1000)
		}

		tmpl.ExecuteWriter(context, res)
	}
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	renderTemplate("index.html", pongo2.Context{}, res)
	// renderTemplate("index.html", pongo2.Context{
	// 	"users": &users,
	// }, res)
}
