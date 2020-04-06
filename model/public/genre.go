package public

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type Genre struct {
	gorm.Model
	Name     string    `json:"name"`
	Programs []Program `json:"programs" gorm:"many2many:genre_programs;"`
}

//GenresType - object type with fields: id, name
var GenresType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Genres",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.Int},
			"name": &graphql.Field{Type: graphql.String},
		},
	})

func (g Genre) Find(id uint) Genre {
	db.Model(g).Where("id = ?", id).First(&g)
	return g
}

type Genres []Genre

func (g Genres) Find() Genres {
	db.Model(&Genre{}).Find(&g)
	return g
}
