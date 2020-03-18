package public

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
)

type Program struct {
	gorm.Model
	ProgramID           string `json:"programId"`
	Title               string
	OriginalTitle       string
	EpisodeTite         string
	Teaser              string
	Description         string
	Cast                string
	CategoryID          uint
	Category            Category `gorm:"foreignkey:category_id"`
	Genres              []Genre  `gorm:"many2many:genre_programs;"`
	SeasonID            uint
	Season              Season `gorm:"foreignkey:season_id"`
	SeasonEpisodeNumber string
	ProductionID        uint
	Production          Production `gorm:"foreignkey:production_id"`
	AirtimeFrom         int        `gorm:"type:bigint"`
	AirtimeTo           int        `gorm:"type:bigint"`
}

func (p Program) Find(id uint) Program {
	db.Model(p).Where("id = ?", id).First(&p)
	return p
}

type Programs []Program

func (p Programs) Find() Programs {
	db.Model(&Program{}).Find(&p)
	return p
}
