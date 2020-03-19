package main

import (
	"log"
	"net/http"

	"github.com/TV2-Bachelorproject/server/controller/auth"
	"github.com/TV2-Bachelorproject/server/controller/people"
	"github.com/TV2-Bachelorproject/server/middleware"
	"github.com/TV2-Bachelorproject/server/model"
	"github.com/TV2-Bachelorproject/server/model/user"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func routes(r *mux.Router) {
	p := mux.NewRouter()
	r.Handle("/people", p)
	p.Use(middleware.Authenticated(user.Admin, user.Producer))
	p.HandleFunc("/people", people.List).Methods("GET")
	p.HandleFunc("/people", people.Create).Methods("POST")
	p.HandleFunc("/people/{id:[0-9]+}", people.Show).Methods("GET")
	p.HandleFunc("/people/{id:[0-9]+}", people.Update).Methods("PUT")
	p.HandleFunc("/people/{id:[0-9]+}", people.Delete).Methods("DELETE")

	r.HandleFunc("/auth/login", auth.Login).Methods("POST")
	r.HandleFunc("/auth/refresh", auth.Refresh).Methods("POST")
}

func main() {
	model.Migrate()

	r := mux.NewRouter()
	routes(r)
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
