package main

import (
	"log"
	"net/http"

	"github.com/TV2-Bachelorproject/server/controller/people"
	"github.com/TV2-Bachelorproject/server/model"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func routes(r *mux.Router) {
	r.HandleFunc("/people", people.List).Methods("GET")
	r.HandleFunc("/people", people.Create).Methods("POST")
	r.HandleFunc("/people/{id:[0-9]+}", people.Show).Methods("GET")
	r.HandleFunc("/people/{id:[0-9]+}", people.Update).Methods("PUT")
	r.HandleFunc("/people/{id:[0-9]+}", people.Delete).Methods("DELETE")
}

func main() {
	model.Migrate()

	r := mux.NewRouter()
	routes(r)
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
