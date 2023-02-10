package api

import "net/http"

func NewRouter() (*http.ServeMux, error) {
	// create a new blank server to avoid using the global serve mux
	m := http.NewServeMux()

	// register the routes
	m.HandleFunc("/db", handleDb)

	return m, nil
}
