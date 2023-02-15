package errors

import (
	"log"
	"net/http"
)

// constants for error handling
const DbNotFound = "database not found"
const DbExists = "database already exists"
const CollectionExists = "collection already exists"
const CollectionNotFound = "collection not found"
const DocumentNotFound = "document not found"
const CollectionsHasKids = "cannot delete collection while it has documents"
const DatabaseHasKids = "cannot delete database while it has collections"

// simple helper function to help reduce boilerplate code when returning errors
func WriteError(w http.ResponseWriter, r *http.Request, err error) {
	// log the error
	log.Println(err.Error())

	// find the best matching http response code for error
	// and write back error to user
	if err.Error() == DbNotFound || err.Error() == CollectionNotFound {
		w.WriteHeader(http.StatusNotFound)
	} else if err.Error() == DbExists || err.Error() == CollectionExists || err.Error() == CollectionsHasKids || err.Error() == DatabaseHasKids {
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write([]byte(err.Error()))
}
