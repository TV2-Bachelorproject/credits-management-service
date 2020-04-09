package public

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type Season struct {
	gorm.Model
	Title       string    `json:"title"`
	RawSeasonID string    `json:"rawSeasonID" gorm:"type:varchar(100)"`
	Programs    []Program `json:"programs"`
	SerieID     uint      `json:"serieId"`
	Serie       Serie     `json:"serie"`
}

//SeasonType - object type with fields: id, title, rawSeasonID,serieId,serie
var SeasonType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Season",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.Int},
			"title":       &graphql.Field{Type: graphql.String},
			"rawSeasonID": &graphql.Field{Type: graphql.String},
			//"programs":    &graphql.Field{Type: &graphql.List{OfType: ProgramType}}, // Don't know if this should be possible?
			"serieId": &graphql.Field{Type: graphql.Int},
			"serie":   &graphql.Field{Type: SerieType},
		},
	})

func (s Season) Find(id uint) Season {
	db.Model(s).Where("id = ?", id).First(&s)
	return s
}

type Seasons []Season

func (s Seasons) Find() Seasons {
	db.Model(&Season{}).Find(&s)
	return s
}
