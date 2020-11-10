package controller

import (
	"encoding/json"
	"net/http"

	"github.com/matheusrf96/api/views"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code:     http.StatusOK,
				Response: "pong",
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
