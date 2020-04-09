package public

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

//CategoryType - object type with fields: id, name
var CategoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Category",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.Int},
			"name": &graphql.Field{Type: graphql.String},
		},
	})

func (c Category) Find(id uint) Category {
	db.Model(c).Where("id = ?", id).First(&c)
	return c
}

type Categories []Category

func (c Categories) Find() Categories {
	db.Model(c).Find(&c)
	return c
}
