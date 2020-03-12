package main

import (
	"log"
	"net/http"

	"github.com/TV2-Bachelorproject/server/controller"
	"github.com/TV2-Bachelorproject/server/model"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func routes(r *mux.Router) {
	r.HandleFunc("/", controller.Hello).Methods("GET")
	r.HandleFunc("/people", controller.People).Methods("GET")
	r.HandleFunc("/people/{id:[0-9]+}", controller.Person).Methods("GET")
}

func main() {
	model.Migrate()

	r := mux.NewRouter()
	routes(r)
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
