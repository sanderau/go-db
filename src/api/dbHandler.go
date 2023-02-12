package api

import (
	"encoding/json"
	"go-db/pkg/model"
	"go-db/pkg/module"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type sessionHandler struct {
	client *module.SessionClient
}

// db post handler to create a new database
func (s sessionHandler) handleDbPost(w http.ResponseWriter, r *http.Request) {
	// handle the DB post request
	log.Println("== handle DB post")

	// read the body
	rb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERR: Error reading the request body " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// unmarshall
	ndb := model.Database{}
	if err := json.Unmarshal(rb, &ndb); err != nil {
		log.Println("ERR: Error unmarshalling the json" + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// add to session
	ndb, err = s.client.AddDatabase(ndb)
	if err != nil {
		log.Println("ERR: Error adding database to session" + err.Error())
		w.WriteHeader(http.StatusConflict)
		io.WriteString(w, err.Error())
		return
	}

	// write back to user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ndb)
}

// get all instances of databases inside the session
func (s sessionHandler) handleDbsGet(w http.ResponseWriter, r *http.Request) {
	// handle the DB Get request
	log.Println("== handle DBs Get")

	// get the db
	udb, err := s.client.GetDatabases()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "error getting databases")
		return
	}

	// if no error return db
	json.NewEncoder(w).Encode(udb)
}

// get just one instance of a database in the session
func (s sessionHandler) handleDbGet(w http.ResponseWriter, r *http.Request) {
	// handle the DB get request
	log.Println("== handle DB get")

	// get the name of the database from the URL
	dbName := mux.Vars(r)["dbName"]

	// try and retrieve the database from the session
	db, err := s.client.GetDatabase(dbName)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	// return the database back to the user
	json.NewEncoder(w).Encode(db)
}

// delete a database from the session
func (s sessionHandler) handleDbDelete(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle db delete")

	// get the name of the db to delete
	dbName := mux.Vars(r)["dbName"]

	// send the request to delete it
	s.client.DeleteDatabase(dbName)

	// return that it was deleted
	w.WriteHeader(http.StatusNoContent)
}

// modify an existing item
func (s sessionHandler) handleDbPut(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle db put")

	// get the name of the database from the URL
	dbName := mux.Vars(r)["dbName"]

	// read the body
	rb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERR: Error reading the request body " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// unmarshall
	ndb := model.Database{}
	if err := json.Unmarshal(rb, &ndb); err != nil {
		log.Println("ERR: Error unmarshalling the json" + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// try modifying an existing
	db, err := s.client.PutDatabase(dbName, ndb.Name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	}

	// return new db if succesfully renamed
	json.NewEncoder(w).Encode(db)
}
