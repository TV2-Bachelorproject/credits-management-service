package mutations

import (
	"github.com/TV2-Bachelorproject/server/model/user"
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/graphql-go/graphql"
)

/*
	Mutation object type with fields "createUser" has type UserType:
       - Name: Mutation
       - Fields: a map of fields by using graphql.Fields
   Setup type of field, args and resolver function:
       - Type: user
       - Args: name,email,password,type
       - Resolve: function to create user
*/
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type:        user.UserType,
			Description: "Create new User",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"type": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				u := user.User{}

				//create user
				u, _ = user.New(
					params.Args["name"].(string),
					params.Args["email"].(string),
					params.Args["password"].(string),
					user.Type(params.Args["type"].(int)))

				db.Create(&u)
				return u, nil

			},
		},
	},
})
