package api

import (
	"encoding/json"
	"go-db/pkg/errors"
	"go-db/pkg/model"
	"io"
	"log"
	"net/http"
)

func (s sessionHandler) handleCollectionPost(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle collections post")

	// get the database name
	u := _getVars(r)

	// read the body
	rb, err := io.ReadAll(r.Body)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// unmarshall
	nc := model.Collection{}
	if err := json.Unmarshal(rb, &nc); err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// write the collection to the corresponding database
	nc, err = s.client.AddCollection(nc, u.dbName)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// if no error return success to user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nc)
}

func (s sessionHandler) handleCollectionsGet(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle collections get")

	// get the database name
	u := _getVars(r)

	// get the collections from the database
	collections, err := s.client.GetCollections(u.dbName)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// return collections to user
	json.NewEncoder(w).Encode(collections)
}

func (s sessionHandler) handleCollectionGet(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle ")

	// get the name for the database and collection
	u := _getVars(r)

	// try and get specific collection from database
	collection, err := s.client.GetCollection(u.dbName, u.cName)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// return collection to user if no errors
	json.NewEncoder(w).Encode(collection)
}

func (s sessionHandler) handleCollectionPut(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle collections put")

	// get the names for the database and collection
	u := _getVars(r)

	// read the body
	rb, err := io.ReadAll(r.Body)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// unmarshall
	c := model.Collection{}
	if err := json.Unmarshal(rb, &c); err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// try to modify collection if it exists
	nc, err := s.client.PutCollection(u.dbName, u.cName, c)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// return success
	json.NewEncoder(w).Encode(nc)
}

func (s sessionHandler) handleCollectionDelete(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle collection delete")

	// get the name for the database and collection
	u := _getVars(r)

	// attempt to delete
	err := s.client.DeleteCollection(u.dbName, u.cName)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// return that item doesn't exist
	w.WriteHeader(http.StatusNoContent)
}
