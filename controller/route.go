package controller

import (
	"net/http"
)

// Register routes on mux
func Register() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", hello())
	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/create", create())
	mux.HandleFunc("/reads", readAll())
	mux.HandleFunc("/read-by-name", readByName())

	return mux
}
