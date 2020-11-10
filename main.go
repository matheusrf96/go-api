package main

import (
	"fmt"
	"net/http"

	"github.com/matheusrf96/api/controller"
)

func main() {
	fmt.Println("Serving on 0.0.0.0:3000")

	mux := controller.Register()

	http.ListenAndServe("0.0.0.0:3000", mux)
}
