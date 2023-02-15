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

func (s sessionHandler) handleCollectionPost(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle collections post")

	// read the body
	rb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERR: Error reading the request body " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// unmarshall
	nc := model.Collection{}
	if err := json.Unmarshal(rb, &nc); err != nil {
		log.Println("ERR: Error unmarshalling the json" + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// get the database name
	dbName := mux.Vars(r)["dbName"]

	// write the collection to the corresponding database
	nc, err = s.client.AddCollection(nc, dbName)
	if err != nil {
		if err.Error() == errors.DbNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		} else if err.Error() == errors.CollectionExists {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}

	// if no error return success to user
	json.NewEncoder(w).Encode(nc)
}

func (s sessionHandler) handleCollectionsGet(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle collections get")

	// get the database name
	dbName := mux.Vars(r)["dbName"]

	// get the collections from the database
	collections, err := s.client.GetCollections(dbName)
	if err != nil {
		if err.Error() == errors.DbNotFound {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(err.Error()))
		return
	}

	// return collections to user
	json.NewEncoder(w).Encode(collections)
}

func (s sessionHandler) handleCollectionGet(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle ")

	// get the name for the database and collection
	dbName := mux.Vars(r)["dbName"]
	collectionName := mux.Vars(r)["collectionName"]

	// try and get specific collection from database
	collection, err := s.client.GetCollection(dbName, collectionName)
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

	// return collection to user if no errors
	json.NewEncoder(w).Encode(collection)
}

func (s sessionHandler) handleCollectionPut(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle collections put")

	// get the names for the database and collection
	dbName := mux.Vars(r)["dbName"]
	oldCollectionName := mux.Vars(r)["collectionName"]

	// read the body
	rb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERR: Error reading the request body " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// unmarshall
	c := model.Collection{}
	if err := json.Unmarshal(rb, &c); err != nil {
		log.Println("ERR: Error unmarshalling the json" + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// try to modify collection if it exists
	nc, err := s.client.PutCollection(dbName, oldCollectionName, c)
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

	// return success
	json.NewEncoder(w).Encode(nc)
}

func (s sessionHandler) handleCollectionDelete(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle collection delete")

	// get the name for the database and collection
	dbName := mux.Vars(r)["dbName"]
	collectionName := mux.Vars(r)["collectionName"]

	// attempt to delete
	s.client.DeleteCollection(dbName, collectionName)

	// return that item doesn't exist
	w.WriteHeader(http.StatusNoContent)
}
