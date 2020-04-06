package public

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type Production struct {
	gorm.Model
	Country     string `json:"country"`
	Year        int    `json:"year"`
	ProducedBy  string `json:"producedBy"`
	ProducedFor string `json:"producedFor"`
	Editor      string `json:"editor"`
}

//ProductionType - object type with fields: id, country, year, producedBy, producedFor, editor
var ProductionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Production",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.Int},
			"country":     &graphql.Field{Type: graphql.String},
			"year":        &graphql.Field{Type: graphql.Int},
			"producedBy":  &graphql.Field{Type: graphql.String},
			"producedFor": &graphql.Field{Type: graphql.String},
			"editor":      &graphql.Field{Type: graphql.String},
		},
	})

func (p Production) Find(id uint) Production {
	db.Model(p).Where("id = ?", id).First(&p)
	return p
}

type Productions []Production

func (p Productions) Find() Productions {
	db.Model(&Production{}).Find(&p)
	return p
}
