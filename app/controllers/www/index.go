package www

import (
	"net/http"

	"github.com/absinsekt/pnk/utils/templateset"
	"github.com/gorilla/mux"
)

// MountIndex attach all entry points of file
func MountIndex(r *mux.Router, t *templateset.TemplateSet) {
	r.Path("/").
		Methods("GET").
		HandlerFunc(getHandlerIndex(t))
}

func getHandlerIndex(templateSet *templateset.TemplateSet) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		templateSet.Render("index.html", res, nil)
	}
}
