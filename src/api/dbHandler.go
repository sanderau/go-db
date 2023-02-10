package api

import (
	"encoding/json"
	"fmt"
	"go-db/pkg/db"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func handleDb(w http.ResponseWriter, r *http.Request) {
	// switch to ensure the requested type is supported
	switch r.Method {
	case http.MethodPost:
		handleDbPost(w, r)
	default:
		log.Println("== Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleDbPost(w http.ResponseWriter, r *http.Request) {
	// handle the DB post request
	log.Println("== handle DB post")

	// read body from req
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("== Error reading the body from the request")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Error reading the body from request")
	}

	// unmarshall json into struct
	udb := db.DB{}
	if err := json.Unmarshal(req, &udb); err != nil {
		log.Println("== Error unmarshalling the body into a json")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "")
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "status - ok")
}
