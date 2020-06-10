package admin

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	"github.com/absinsekt/pnk/controllers/middlewares"
	"github.com/absinsekt/pnk/models/user"
	"github.com/absinsekt/pnk/lib/templateset"
)

// Mount attach all entry points of file
func Mount(r *mux.Router, t *templateset.TemplateSet) {
	sub := r.PathPrefix("/admin").Subrouter()

	// must be first
	sub.Use(middlewares.CSRFMiddleware)
	sub.Use(middlewares.GetAuthMiddleware(t, true))

	sub.Path("/").Methods("GET").HandlerFunc(getHandlerIndex(t))
}

func getHandlerIndex(templateSet *templateset.TemplateSet) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		var usr *user.SessionData

		if _usr := context.Get(req, "user"); _usr != nil {
			usr = _usr.(*user.SessionData)
		}

		templateSet.Render("admin_index.html", res, req, map[string]interface{}{"user": usr})
	}
}
