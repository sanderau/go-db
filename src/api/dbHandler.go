package api

import (
	"encoding/json"
	"go-db/pkg/errors"
	"go-db/pkg/model"
	"io"
	"log"
	"net/http"
)

// db post handler to create a new database
func (s sessionHandler) handleDbPost(w http.ResponseWriter, r *http.Request) {
	// handle the DB post request
	log.Println("== handle DB post")

	// read the body
	rb, err := io.ReadAll(r.Body)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// unmarshall
	ndb := model.Database{}
	if err := json.Unmarshal(rb, &ndb); err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// add to session
	ndb, err = s.client.AddDatabase(ndb)
	if err != nil {
		errors.WriteError(w, r, err)
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
		errors.WriteError(w, r, err)
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
	u := _getVars(r)

	// try and retrieve the database from the session
	db, err := s.client.GetDatabase(u.dbName)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// return the database back to the user
	json.NewEncoder(w).Encode(db)
}

// delete a database from the session
func (s sessionHandler) handleDbDelete(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle db delete")

	// get the name of the db to delete
	u := _getVars(r)

	// send the request to delete it
	err := s.client.DeleteDatabase(u.dbName)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// return that it was deleted
	w.WriteHeader(http.StatusNoContent)
}

// modify an existing item
func (s sessionHandler) handleDbPut(w http.ResponseWriter, r *http.Request) {
	log.Println("== handle db put")

	// get the name of the database from the URL
	u := _getVars(r)

	// read the body
	rb, err := io.ReadAll(r.Body)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// unmarshall
	ndb := model.Database{}
	if err := json.Unmarshal(rb, &ndb); err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// try modifying an existing
	db, err := s.client.PutDatabase(u.dbName, ndb.Name)
	if err != nil {
		errors.WriteError(w, r, err)
		return
	}

	// return new db if succesfully renamed
	json.NewEncoder(w).Encode(db)
}
