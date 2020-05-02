package public

import (
	"fmt"
	"log"
	"time"

	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/graphql-go/graphql"
)

//Program struct
type Program struct {
	ID                  uint `gorm:"primary_key"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time `sql:"index"`
	ProgramID           string     `json:"programId`
	Title               string     `json:"title"`
	Teaser              string     `json:"teaser"`
	Description         string     `json:"description"`
	Cast                string     `json:"cast"`
	CategoryID          uint       `json:"categoryId"`
	Category            Category   `json:"category" gorm:"foreignkey:CategoryID"`
	Genres              []Genre    `json:"genres" gorm:"many2many:genre_programs;"`
	SeasonID            uint       `json:"seasonId`
	Season              Season     `json:"season" gorm:"foreignkey:SeasonID"`
	SeasonEpisodeNumber int        `json:"seasonEpisodeNumber"`
	LinearEpisodeNumber int        `json:"linearEpisodeNumber"`
	ProductionID        uint       `json:"productionId"`
	Production          Production `json:"production" gorm:"foreignkey:ProductionID"`
	SerieID             uint       `json:"serieId"`
	Serie               Serie      `json:"serie" gorm:"foreignkey:SerieID"`
	AirtimeFrom         int        `json:"airTimeFrom" gorm:"type:bigint"`
	AirtimeTo           int        `json:"airTimeTo" gorm:"type:bigint"`
	Credits             []Credit   `json:"credits"`
}

// ProgramType is the GraphQL schema/typedef for the program type.
var ProgramType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Program",
		Fields: graphql.Fields{
			"id":                  &graphql.Field{Type: graphql.Int},
			"programId":           &graphql.Field{Type: graphql.String},
			"title":               &graphql.Field{Type: graphql.String},
			"teaser":              &graphql.Field{Type: graphql.String},
			"description":         &graphql.Field{Type: graphql.String},
			"cast":                &graphql.Field{Type: graphql.String},
			"categoryId":          &graphql.Field{Type: graphql.Int},
			"category":            &graphql.Field{Type: CategoryType},
			"genres":              &graphql.Field{Type: &graphql.List{OfType: GenresType}},
			"seasonId":            &graphql.Field{Type: graphql.Int},
			"season":              &graphql.Field{Type: SeasonType},
			"seasonEpisodeNumber": &graphql.Field{Type: graphql.Int},
			"linearEpisodeNumber": &graphql.Field{Type: graphql.Int},
			"productionId":        &graphql.Field{Type: graphql.Int},
			"production":          &graphql.Field{Type: ProductionType},
			"serieId":             &graphql.Field{Type: graphql.Int},
			"serie":               &graphql.Field{Type: SerieType},
			"airTimeFrom":         &graphql.Field{Type: graphql.String},
			"airTimeTo":           &graphql.Field{Type: graphql.String},
			"credits":             &graphql.Field{Type: &graphql.List{OfType: CreditType}},
		},
	})

//Find single program entry
func (p Program) Find(id uint) Program {
	//Preload preloads structs - Creates a SQL query pr. Preload. Should be fixed in Gorm V2.
	if err := db.Model(p).
		Preload("Production").
		Preload("Category").
		Preload("Genres").
		Preload("Serie").
		Preload("Season").
		Preload("Credits").
		Where("id = ?", id).
		First(&p).Error; err != nil {
		fmt.Println(err)
	}
	return p
}

//Programs struct
type Programs []Program

//Find all programs
func (p Programs) Find() Programs {
	//Preload preloads structs - Creates a SQL query pr. Preload. Should be fixed in Gorm V2.
	errors := db.Model(&Program{}).
		Preload("Production").
		Preload("Category").
		Preload("Genres").
		Preload("Serie").
		Preload("Season").
		Preload("Credits").
		Find(&p).GetErrors()

	for _, err := range errors {
		log.Fatal(err)
	}

	return p
}
