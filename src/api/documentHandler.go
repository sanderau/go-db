package api

import (
	"encoding/json"
	"go-db/pkg/errors"
	"go-db/pkg/model"
	"io"
	"log"
	"net/http"
)

func (s sessionHandler) handleDocumentPost(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle document post")

	// get the names for the database and collection
	u := _getVars(r)

	// read the body
	rb, err := io.ReadAll(r.Body)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// unmarshall body into document struct
	doc := model.Document{}
	if err := json.Unmarshal(rb, &doc); err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// put document into corresponding database=>collection
	doc, err = s.client.PostDocument(u.dbName, u.cName, doc)
	if err != nil {
		errors.WriteError(w, r, err)
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
	json.NewEncoder(w).Encode(docs)
}

func (s sessionHandler) handleDocumentsGetBySearch(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle documents get by search")

	// get the names for the database and collection
	u := _getVars(r)

	// read the body
	rb, err := io.ReadAll(r.Body)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// unmarshall body into search struct
	search := model.Search{}
	if err := json.Unmarshal(rb, &search); err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// get documents by a search keyword
	docs, err := s.client.GetDocumentsBySearch(u.dbName, u.cName, search)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// return the documents
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
	json.NewEncoder(w).Encode(doc)
}

func (s sessionHandler) handleDocumentPut(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle document put")

	// get the variables from the url
	u := _getVars(r)

	// read the body
	rb, err := io.ReadAll(r.Body)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// unmarshall body into document struct
	doc := model.Document{}
	if err := json.Unmarshal(rb, &doc); err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// call the backend function and handle any errors
	newDoc, err := s.client.PutDocument(u.dbName, u.cName, u.id, doc)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// return the new doc and return success
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
