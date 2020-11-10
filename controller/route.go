package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", hello())
	mux.HandleFunc("/ping", ping())

	return mux
}
