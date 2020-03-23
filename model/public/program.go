package public

import (
	"fmt"

	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
)

//Program struct
type Program struct {
	gorm.Model
	ProgramID           string `json:"programId"`
	Title               string
	Teaser              string
	Description         string
	Cast                string
	CategoryID          uint
	Category            Category `gorm:"foreignkey:CategoryID"`
	Genres              []Genre  `gorm:"many2many:genre_programs;"`
	SeasonID            uint
	Season              Season `gorm:"foreignkey:SeasonID"`
	SeasonEpisodeNumber int
	LinearEpisodeNumber int
	ProductionID        uint
	Production          Production `gorm:"foreignkey:ProductionID"`
	SerieID             uint
	Serie               Serie     `gorm:"foreignkey:SerieID"`
	AirtimeFrom         int       `gorm:"type:bigint"`
	AirtimeTo           int       `gorm:"type:bigint"`
	Credit              []Credits `gorm:"many2many:credit_groups;"`

}

//Find single program entry
func (p Program) Find(id uint) Program {
	//Preload preloads structs - Creates a SQL query pr. Preload. Should be fixed in Gorm V2.
	if err := db.Model(p).
		Preload("Production").
		Preload("Category").
		Preload("Genres").
		Preload("Serie").
		Preload("Season").
		Preload("Credit").
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
	db.Model(&Program{}).
		Preload("Production").
		Preload("Category").
		Preload("Genres").
		Preload("Serie").
		Preload("Season").
		Find(&p).GetErrors()
	return p
}
