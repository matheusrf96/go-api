package controller

import (
	"encoding/json"
	"net/http"

	"github.com/matheusrf96/api/model"
	"github.com/matheusrf96/api/views"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)

			err := model.CreateTodo(data.Name, data.Todo)

			if err != nil {
				w.Write([]byte("Some error ocurried :("))
				return
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}
	}
}
