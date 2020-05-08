package api

import (
	"github.com/absinsekt/pnk/controllers/middlewares"
	"github.com/gorilla/mux"
)

// Mount attach all entry points of file
func Mount(r *mux.Router) {
	sub := r.PathPrefix("/api").Subrouter()

	sub.Use(middlewares.CSRFMiddleware)

	sub.Path("/login/").Methods("POST").HandlerFunc(handleLogin)
}
