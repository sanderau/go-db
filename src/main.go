package main

import (
	"go-db/api"
	"log"
	"net/http"
)

func main() {
	// get a new server from API
	serv, err := api.NewRouter()
	if err != nil {
		log.Fatal(err)
	}

	// listen and serve the api
	log.Println("== Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", serv))
}
