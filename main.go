package main

import (
	"log"
	"net/http"

	"github.com/TV2-Bachelorproject/server/controller/auth"
	"github.com/TV2-Bachelorproject/server/controller/people"
	"github.com/TV2-Bachelorproject/server/controller/programs"
	"github.com/TV2-Bachelorproject/server/controller/users"
	"github.com/TV2-Bachelorproject/server/graphql/mutations"
	"github.com/TV2-Bachelorproject/server/graphql/queries"
	"github.com/TV2-Bachelorproject/server/middleware"
	"github.com/TV2-Bachelorproject/server/model"
	"github.com/TV2-Bachelorproject/server/model/user"
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

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

	//Route for Graphql TODO Needs authentication
	startGraphql(r)

}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST,OPTIONS")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func startGraphql(r *mux.Router) {
	// Schema for graphql.
	var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: queries.GetRootFields(),
		}),
		Mutation: mutations.UserType,
	})

	//create graphql-go HTTP handler for schema
	h := handler.New(&handler.Config{
		Schema:     &Schema,
		Pretty:     true, // return pretty json
		GraphiQL:   true,
		Playground: true,
	})

	// serve the GraphQL endpoint at "/graphql"
	r.Handle("/graphql", CorsMiddleware(h))
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
	if err := http.ListenAndServe(":3000", CorsMiddleware(r)); err != nil {
		log.Fatal(err)
	}
}
