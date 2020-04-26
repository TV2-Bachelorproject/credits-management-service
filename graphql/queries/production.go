package queries

import (
	"github.com/TV2-Bachelorproject/server/model/public"
	"github.com/graphql-go/graphql"
)

/*
	Query object type with fields "production" has type ProductionType:
       - Name: Query
       - Fields: a map of fields by using graphql.Fields
   Setup type of field, args and resolver function:
       - Type: production
       - Args: id
       - Resolve: function to query data using params from [Args] and return value with current type
*/
var ProductionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"production": &graphql.Field{
				Type:        public.ProductionType,
				Description: "Get production by ID",
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
						production := public.Production{}.Find(uintID)
						return production, nil
					}
					return nil, nil
				},
			},

			"productionList": &graphql.Field{
				Type:        graphql.NewList(public.ProductionType),
				Description: "Get list of productions",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					productions := public.Productions{}.Find()
					return productions, nil
				},
			},
		},
	})

func GetProductionQuery() *graphql.Field {
	return &graphql.Field{
		Type:        public.ProductionType,
		Description: "Get production by ID",
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
				production := public.Production{}.Find(uintID)
				return production, nil
			}
			return nil, nil
		},
	}
}

func GetProductionsQuery() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(public.ProductionType),
		Description: "Get list of productions",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			productions := public.Productions{}.Find()
			return productions, nil
		},
	}
}
