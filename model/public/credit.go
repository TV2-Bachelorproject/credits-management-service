package public

import (
	"time"

	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/graphql-go/graphql"
)

type CreditGroup struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string
}

type Credit struct {
	ID            uint `gorm:"primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time  `sql:"index"`
	Persons       []Person    `gorm:"many2many:credit_persons;"`
	CreditGroupID uint        `json:"-"`
	CreditGroup   CreditGroup `json:"creditGroup" gorm:"foreignkey:CreditGroupID;"`
	ProgramID     uint        `json:"-"`
	SeasonID      uint        `json:"-"`
	SerieID       uint        `json:"-"`
	Accepted      bool        `json:"accepted"`
}

var CreditGroupType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CreditGroup",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.Int},
			"name": &graphql.Field{Type: graphql.String},
		},
	})

//CreditType - object type with fields: id, title, rawSeasonID,serieId,serie
var CreditType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Credit",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.Int},
			"persons":     &graphql.Field{Type: &graphql.List{OfType: PersonType}},
			"creditGroup": &graphql.Field{Type: CreditGroupType},
			"accepted":    &graphql.Field{Type: graphql.Boolean},
		},
	})

func (s Credit) Find(id uint) Credit {
	db.Model(s).Where("id = ?", id).First(&s)
	return s
}

type Credits []Credit

func (s Credits) Find() Credits {
	db.Model(&Credit{}).Find(&s)
	return s
}

func (c Credits) ForProgram(id uint) Credits {
	db.Model(&Credit{}).Where("program_id = ?", id).Preload("Persons").Preload("CreditGroup").
		Find(&c)
	return c
}
