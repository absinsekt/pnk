package www

import (
	"net/http"

	"github.com/absinsekt/pnk/utils/core"
	"github.com/gorilla/mux"
)

// MountIndex attach all entry points of file
func MountIndex(r *mux.Router, t *core.TemplateSet) {
	r.Path("/").
		Methods("GET").
		HandlerFunc(getHandlerIndex(t))
}

func getHandlerIndex(templateSet *core.TemplateSet) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		templateSet.Render("index.html", res, nil)
	}
}
