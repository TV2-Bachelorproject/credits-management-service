package public

import (
	"time"

	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/graphql-go/graphql"
)

type Person struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string     `json:"name"`
}

//PersonType - object type with fields: TODO
var PersonType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Person",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.Int},
			"name": &graphql.Field{Type: graphql.String},
		},
	})

func (p Person) Find(id uint) Person {
	db.Model(p).Where("id = ?", id).First(&p)
	return p
}

type People []Person

func (p People) Find() People {
	db.Model(&Person{}).Find(&p)
	return p
}
