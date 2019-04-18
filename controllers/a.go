package controllers

import (
	"github.com/absinsekt/pnk/configuration"
	"github.com/absinsekt/pnk/utils"
	"github.com/flosch/pongo2"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	tmplBaseDir    = "templates"
	wwwTemplateSet *pongo2.TemplateSet
)

func renderTemplate(templateName string, context pongo2.Context, res http.ResponseWriter) {
	var (
		tmpl *pongo2.Template
		err  error
	)

	timerStart := time.Now()

	if wwwTemplateSet.Debug {
		tmpl, err = wwwTemplateSet.FromFile(templateName)
	} else {
		tmpl, err = wwwTemplateSet.FromCache(templateName)
	}

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

// NewControllersRouter mounts all paths, e. g. apis and static page handlers
// and returns root Router
func NewControllersRouter() *mux.Router {
	wwwFSLoader, err := pongo2.NewSandboxedFilesystemLoader(tmplBaseDir)
	utils.Check(err, true)

	wwwTemplateSet = pongo2.NewSet("www", wwwFSLoader)
	wwwTemplateSet.Debug = configuration.Debug

	root := mux.NewRouter()
	root.HandleFunc("/", handleIndex)

	// for debug only
	root.PathPrefix("/dist").Handler(http.StripPrefix("/dist", http.FileServer(http.Dir("static/dist"))))

	return root
}
