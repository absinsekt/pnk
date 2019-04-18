package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func mountUsers(node *mux.Router) {
	n := node.PathPrefix("/users/")

	n.Methods("GET").HandlerFunc(getUsers)
}

func getUsers(res http.ResponseWriter, req *http.Request) {
	log.Println("get users")
}
