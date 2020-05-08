package www

import (
	"net/http"

	"github.com/absinsekt/pnk/controllers/middlewares"

	"github.com/absinsekt/pnk/utils/templateset"
	"github.com/gorilla/mux"
)

// Mount attach all entry points of file
func Mount(r *mux.Router, t *templateset.TemplateSet) {
	sub := r.PathPrefix("/").Subrouter()

	sub.Use(middlewares.CSRFMiddleware)

	sub.Path("/").Methods("GET").HandlerFunc(getHandlerIndex(t))
}

func getHandlerIndex(templateSet *templateset.TemplateSet) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		templateSet.Render("index.html", res, req, nil)
	}
}
