package public

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type Serie struct {
	gorm.Model
	Title   string
	Seasons []Season
}

//SerieType - object type with fields: id, title
var SerieType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Serie",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.Int},
			"title": &graphql.Field{Type: graphql.String},
		},
	})

func (s Serie) Find(id uint) Serie {
	db.Model(s).Where("id = ?", id).First(&s)
	return s
}

type Series []Serie

func (s Series) Find() Series {
	db.Model(&Serie{}).Find(&s)
	return s
}
