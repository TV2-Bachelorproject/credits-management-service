package main

import (
	"log"
	"net/http"

	"github.com/TV2-Bachelorproject/server/controller/auth"
	"github.com/TV2-Bachelorproject/server/controller/credits"
	"github.com/TV2-Bachelorproject/server/controller/people"
	"github.com/TV2-Bachelorproject/server/controller/programs"
	"github.com/TV2-Bachelorproject/server/controller/users"
	"github.com/TV2-Bachelorproject/server/graphql/queries"
	"github.com/TV2-Bachelorproject/server/middleware"
	"github.com/TV2-Bachelorproject/server/model"
	"github.com/TV2-Bachelorproject/server/model/user"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Schema for graphql.
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queries.ProgramType,
	//Mutation: mutations.UserType,
})

func routes(r *mux.Router) {
	u := mux.NewRouter()
	u.Use(middleware.Authenticated(user.Admin))
	r.Handle("/users", u)
	r.Handle("/users/{id:[0-9]+}", u)
	u.HandleFunc("/users", users.List).Methods("GET")
	u.HandleFunc("/users", users.Create).Methods("POST")
	u.HandleFunc("/users/{id:[0-9]+}", users.Show).Methods("GET")
	u.HandleFunc("/users/{id:[0-9]+}", users.Update).Methods("PUT")
	u.HandleFunc("/users/{id:[0-9]+}", users.Delete).Methods("DELETE")

	c := mux.NewRouter()
	c.Use(middleware.Authenticated(user.Admin, user.Producer))
	c.Handle("/credits", c)
	c.HandleFunc("/credits", credits.Create).Methods("POST")
	c.HandleFunc("/credits", credits.Delete).Methods("DELETE")

	ca := mux.NewRouter()
	ca.Use(middleware.Authenticated(user.Admin))
	ca.Handle("/credits", ca)
	ca.HandleFunc("/credits/accept", credits.Accept).Methods("POST")

	p := mux.NewRouter()
	p.Use(middleware.Authenticated(user.Admin, user.Producer))
	r.Handle("/people", p)
	r.Handle("/people/{id:[0-9]+}", p)
	p.HandleFunc("/people", people.List).Methods("GET")
	p.HandleFunc("/people", people.Create).Methods("POST")
	p.HandleFunc("/people/{id:[0-9]+}", people.Show).Methods("GET")
	p.HandleFunc("/people/{id:[0-9]+}", people.Update).Methods("PUT")
	p.HandleFunc("/people/{id:[0-9]+}", people.Delete).Methods("DELETE")

	//Routes for programs
	r.HandleFunc("/programs", programs.GetAll).Methods("GET")
	r.HandleFunc("/programs/{id:[0-9]+}", programs.Get).Methods("GET")

	r.HandleFunc("/auth/login", auth.Login).Methods("POST")
	r.HandleFunc("/auth/refresh", auth.Refresh).Methods("POST")

	//Route for Graphql TODO Needs authentication
	startGraphql(r)

}

func startGraphql(r *mux.Router) {
	//create graphql-go HTTP handler for schema
	h := handler.New(&handler.Config{
		Schema:     &Schema,
		Pretty:     true, // return pretty json
		GraphiQL:   true,
		Playground: true,
	})

	g := mux.NewRouter()
	g.Use(middleware.Validate)
	g.Handle("/graphql", h)

	// serve the GraphQL endpoint at "/graphql"
	r.Handle("/graphql", g)
}

func main() {
	model.Migrate()
	model.Seed()

	r := mux.NewRouter()

	//setup Routes
	routes(r)

	// and serve!
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
