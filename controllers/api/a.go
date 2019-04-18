package api

import (
	"github.com/gorilla/mux"
)

// Mount mounts api subpath
func Mount(root *mux.Router) {
	r := root.
		StrictSlash(true).
		PathPrefix("/api").
		Subrouter()

	// add here more subpaths for business entities
	mountUsers(r)
}
