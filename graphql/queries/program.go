package queries

import (
	"errors"

	"github.com/TV2-Bachelorproject/server/model/public"
	"github.com/graphql-go/graphql"
)

/*
	Query object type with fields "program" has type ProgramType:
       - Name: Query
       - Fields: a map of fields by using graphql.Fields
   Setup type of field, args and resolver function:
       - Type: program
       - Args: id
       - Resolve: function to query data using params from [Args] and return value with current type
*/
var ProgramType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"program": &graphql.Field{
				Type:        public.ProgramType,
				Description: "Get program by ID",
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

			"programList": &graphql.Field{
				Type:        graphql.NewList(public.ProgramType),
				Description: "Get list of programs",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					programs := public.Programs{}.Find()
					return programs, nil
				},
			},
		},
	})

func GetProgramQuery() *graphql.Field {
	return &graphql.Field{
		Type:        public.ProgramType,
		Description: "Get program by ID",
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

				return program, errors.New("Program not found!")
			}
			return nil, errors.New("Program not found!")
		},
	}
}

func GetProgramsQuery() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(public.ProgramType),
		Description: "Get list of programs",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			programs := public.Programs{}.Find()
			return programs, nil
		},
	}
}
