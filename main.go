package main

import (
	"net/http"

	"github.com/TV2-Bachelorproject/server/controller"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Hello)
	http.ListenAndServe(":3000", r)
}
