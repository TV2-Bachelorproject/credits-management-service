package main

import (
	"log"
	"net/http"

	"github.com/TV2-Bachelorproject/server/controller"
	"github.com/TV2-Bachelorproject/server/model/private"
	"github.com/TV2-Bachelorproject/server/model/public"
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func routes(r *mux.Router) {
	r.HandleFunc("/", controller.Hello).Methods("GET")
	r.HandleFunc("/people", controller.People).Methods("GET")
	r.HandleFunc("/people/{id:[0-9]+}", controller.Person).Methods("GET")
}

func main() {
	err := db.Setup()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	db.Migrate(
		&private.Person{},
		&public.Credit{},
		&public.CreditGroup{},
		&public.Program{},
		&public.Season{},
		&public.Serie{},
		&public.Serie{},
	)

	r := mux.NewRouter()
	routes(r)
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
