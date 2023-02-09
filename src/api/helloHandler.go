package api

import (
	"fmt"
	"log"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	// switch to ensure the requested type is supported
	switch r.Method {
	case http.MethodPost:
		handleHelloPost(w, r)
	default:
		log.Println("== Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleHelloPost(w http.ResponseWriter, r *http.Request) {
	// handle the hello post request
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "status - ok")
}
