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
	m.HandleFunc("/db/{dbName}", s.handleDbGet).Methods(http.MethodGet)
	m.HandleFunc("/db/{dbName}", s.handleDbPut).Methods(http.MethodPut)
	m.HandleFunc("/db/{dbName}", s.handleDbDelete).Methods(http.MethodDelete)

	// Handle the collection functions
	m.HandleFunc("/db/{dbName}/collection", s.handleCollectionPost).Methods(http.MethodPost)
	m.HandleFunc("/db/{dbName}/collection", s.handleCollectionsGet).Methods(http.MethodGet)
	m.HandleFunc("/db/{dbName}/collection/{collectionName}", s.handleCollectionGet).Methods(http.MethodGet)
	m.HandleFunc("/db/{dbName}/collection/{collectionName}", s.handleCollectionPut).Methods(http.MethodPut)
	m.HandleFunc("/db/{dbName}/collection/{collectionName}", s.handleCollectionDelete).Methods(http.MethodDelete)

	// Handle all of the document functions
	m.HandleFunc("/db/{dbName}/collection/{collectionName}/document", s.handleDocumentPost).Methods(http.MethodPost)
	m.HandleFunc("/db/{dbName}/collection/{collectionName}/document", s.handleDocumentsGet).Methods(http.MethodGet)
	m.HandleFunc("/db/{dbName}/collection/{collectionName}/document/{id}", s.handleDocumentGet).Methods(http.MethodGet)
	m.HandleFunc("/db/{dbName}/collection/{collectionName}/document/{id}", s.handleDocumentPut).Methods(http.MethodPut)
	m.HandleFunc("/db/{dbName}/collection/{collectionName}/document/{id}", s.handleDocumentDelete).Methods(http.MethodDelete)

	return m, nil
}
