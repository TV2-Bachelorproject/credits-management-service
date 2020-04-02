package main

import (
	"log"
	"net/http"

	"github.com/TV2-Bachelorproject/server/controller/auth"
	"github.com/TV2-Bachelorproject/server/controller/people"
	"github.com/TV2-Bachelorproject/server/controller/programs"
	"github.com/TV2-Bachelorproject/server/controller/users"
	"github.com/TV2-Bachelorproject/server/middleware"
	"github.com/TV2-Bachelorproject/server/model"
	"github.com/TV2-Bachelorproject/server/model/public"
	"github.com/TV2-Bachelorproject/server/model/user"
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// ProgramType is the GraphQL schema/typedef for the program type.
var ProgramType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Program",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.Int},
			"programId":   &graphql.Field{Type: graphql.String},
			"title":       &graphql.Field{Type: graphql.String},
			"teaser":      &graphql.Field{Type: graphql.String},
			"description": &graphql.Field{Type: graphql.String},
			"categoryId":  &graphql.Field{Type: graphql.Int},
			//"category":            &graphql.Field{Type: graphql.String},
			//"genres":              &graphql.Field{Type: graphql.String},
			"seasonId": &graphql.Field{Type: graphql.Int},
			//"season":              &graphql.Field{Type: graphql.String},
			"seasonEpisodeNumber": &graphql.Field{Type: graphql.Int},
			"linearEpisodeNumber": &graphql.Field{Type: graphql.Int},
			"productionId":        &graphql.Field{Type: graphql.Int},
			//"production":          &graphql.Field{Type: graphql.String},
			"serieId": &graphql.Field{Type: graphql.Int},
			//"serie":               &graphql.Field{Type: graphql.String},
			"airTimeFrom": &graphql.Field{Type: graphql.Int},
			"airTimeTo":   &graphql.Field{Type: graphql.Int},
			//"credit":              &graphql.Field{Type: graphql.String},
		},
	})

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"program": &graphql.Field{
				Type:        ProgramType,
				Description: "Get program by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)

					if ok {
						var uintID = uint(id)
						//Find the program
						program := public.Program{}.Find(uintID)
						return program, nil
					}
					return nil, nil
				},
			},
		},
	})

// Schema for graphql. TODO Add mutation
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
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

	//Route for Graphql
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

	// serve the GraphQL endpoint at "/graphql"
	r.Handle("/graphql", h)
}

func main() {
	model.Migrate()

	u1, err := user.New("admin", "admin@example.com", "123456", user.Admin)

	if err != nil {
		log.Fatal(err)
	}

	u2, err := user.New("producer", "producer@example.com", "123456", user.Producer)

	if err != nil {
		log.Fatal(err)
	}

	db.Create(&u1)
	db.Create(&u2)

	r := mux.NewRouter()

	//setup Routes
	routes(r)

	// and serve!
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
