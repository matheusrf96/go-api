package controller

import (
	"encoding/json"
	"net/http"

	"github.com/matheusrf96/api/model"
)

func reads() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data, err := model.ReadAll()

			if err != nil {
				w.Write([]byte(err.Error()))
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
