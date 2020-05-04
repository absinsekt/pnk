package admin

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	"github.com/absinsekt/pnk/controllers/middlewares"
	"github.com/absinsekt/pnk/models/user"
	"github.com/absinsekt/pnk/utils/core"
)

// MountAdmin attach all entry points of file
func MountAdmin(r *mux.Router, t *core.TemplateSet) {
	sub := r.Path("/admin/").Subrouter()

	sub.Use(middlewares.AuthMiddleware)

	sub.
		Path("/").
		Methods("GET").
		HandlerFunc(getHandlerIndex(t))
}

func getHandlerIndex(templateSet *core.TemplateSet) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		var usr *user.SessionData

		if _usr := context.Get(req, "user"); _usr != nil {
			usr = _usr.(*user.SessionData)
		}

		templateSet.Render("admin_index.html", res, map[string]interface{}{"user": usr})
	}
}
