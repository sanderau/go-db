package errors

import (
	"log"
	"net/http"
)

// constants for error handling
const DbNotFound = "database not found"
const CollectionExists = "collection already exists"
const CollectionNotFound = "collection not found"
const DocumentNotFound = "document not found"

// simple helper function to help reduce boilerplate code when returning errors
func WriteError(w http.ResponseWriter, r *http.Request, err error) {
	// log the error
	log.Println(err.Error())

	// find the best matching http response code for error
	// and write back error to user
	if err.Error() == DbNotFound {
		w.WriteHeader(http.StatusNotFound)
	} else if err.Error() == CollectionExists {
		w.WriteHeader(http.StatusConflict)
	} else if err.Error() == CollectionNotFound {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write([]byte(err.Error()))
}
