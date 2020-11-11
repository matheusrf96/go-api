package controller

import (
	"net/http"

	"github.com/matheusrf96/api/model"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			err := model.CreateTodo()

			if err != nil {
				w.Write([]byte("Some error ocurried :("))
				return
			}
		}
	}
}
