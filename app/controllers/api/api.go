package api

import (
	"github.com/gorilla/mux"
)

// Mount attach all entry points of file
func Mount(r *mux.Router) {
	sub := r.PathPrefix("/api").Subrouter()
	sub.Path("/login/").Methods("POST").HandlerFunc(handleLogin)
}
