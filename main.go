package main

import (
	"fmt"
	"net/http"

	"github.com/matheusrf96/api/controller"
	"github.com/matheusrf96/api/model"
)

func main() {
	fmt.Println("Serving on 0.0.0.0:3000")

	mux := controller.Register()

	db := model.Connect()
	defer db.Close()

	http.ListenAndServe("0.0.0.0:3000", mux)
}
