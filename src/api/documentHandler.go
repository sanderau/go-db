package api

import (
	"encoding/json"
	"go-db/pkg/errors"
	"go-db/pkg/model"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type urlVars struct {
	dbName string
	cName  string
	id     string
}

// helper function to reduce boiler plate code when retrieving variables from the URL
func _getVars(r *http.Request) urlVars {
	var u urlVars

	u.dbName = mux.Vars(r)["dbName"]
	u.cName = mux.Vars(r)["collectionName"]
	u.id = mux.Vars(r)["id"]

	return u
}

func (s sessionHandler) handleDocumentPost(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle document post")

	// get the names for the database and collection
	dbName := mux.Vars(r)["dbName"]
	cName := mux.Vars(r)["collectionName"]

	// read the body
	rb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERR: Error reading the request body " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// unmarshall body into document struct
	doc := model.Document{}
	if err := json.Unmarshal(rb, &doc); err != nil {
		log.Println("ERR: Error unmarshalling the json" + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// put document into corresponding database=>collection
	doc, err = s.client.PostDocument(dbName, cName, doc)
	if err != nil {
		if err.Error() == errors.DbNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else if err.Error() == errors.CollectionNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(err.Error()))
		return
	}

	// write the docuemnt back to user, and the status created code
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(doc)
}

func (s sessionHandler) handleDocumentsGet(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle documents get")

	// get the names for the database and collection
	u := _getVars(r)

	// get the documents from the session client
	docs, err := s.client.GetDocuments(u.dbName, u.cName)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// return the documents
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(docs)
}

func (s sessionHandler) handleDocumentGet(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle document get")

	// get the url variabls
	u := _getVars(r)

	// attempt to get the collection by id
	doc, err := s.client.GetDocument(u.dbName, u.cName, u.id)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// return the document to user
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(doc)
}

func (s sessionHandler) handleDocumentPut(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle document put")

	// get the variables from the url
	u := _getVars(r)

	// read the body
	rb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERR: Error reading the request body " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// unmarshall body into document struct
	doc := model.Document{}
	if err := json.Unmarshal(rb, &doc); err != nil {
		log.Println("ERR: Error unmarshalling the json" + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// call the backend function and handle any errors
	newDoc, err := s.client.PutDocument(u.dbName, u.cName, u.id, doc)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// return the new doc and return success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newDoc)
}

func (s sessionHandler) handleDocumentDelete(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle document delete")

	// get the vars from the url
	u := _getVars(r)

	// try deleting the document
	err := s.client.DeleteDocument(u.dbName, u.cName, u.id)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// return there is no content
	w.WriteHeader(http.StatusNoContent)
}
