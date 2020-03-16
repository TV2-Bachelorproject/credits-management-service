package public

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Program struct {
	gorm.Model
	ProgramID           string `json:"programId" gorm:"primary_key"`
	Title               string
	OriginalTitle       string
	Teaser              string
	Description         string
	Cast                string
	Category            string
	Genres              pq.StringArray `gorm:"type:varchar(100)[]"`
	SeasonID            string         `gorm:"type:varchar(100)"`
	SeasonEpisodeNumber string
	ProductionID        int
	Production          Production `gorm:"foreignkey:production_id"`
	AirtimeFrom         int        `gorm:"type:bigint"`
	AirtimeTo           int        `gorm:"type:bigint"`
	Airtime             struct {
		From int `gorm:"-"`
		To   int `gorm:"-"`
	}
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
