package api

import (
	"go-db/pkg/module"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() (*mux.Router, error) {
	// create a new blank server to avoid using the global serve mux
	m := mux.NewRouter()

	// create a new session client that will handle our data
	sc := module.NewSessionClient()
	s := sessionHandler{
		client: sc,
	}

	// Handle database functions
	m.HandleFunc("/db", s.handleDbsGet).Methods(http.MethodGet)
	m.HandleFunc("/db", s.handleDbPost).Methods(http.MethodPost)
	m.HandleFunc("/db/{name}", s.handleDbGet).Methods(http.MethodGet)
	m.HandleFunc("/db/{name}", s.handleDbPut).Methods(http.MethodPut)
	m.HandleFunc("/db/{name}", s.handleDbDelete).Methods(http.MethodDelete)

	return m, nil
}
