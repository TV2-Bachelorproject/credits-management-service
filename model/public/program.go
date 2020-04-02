package public

import (
	"fmt"

	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
)

//Program struct
type Program struct {
	gorm.Model
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
	Credit              []Credits  `json:"credit" gorm:"many2many:credit_groups;"`
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
	errors := db.Model(&Program{}).
		Preload("Production").
		Preload("Category").
		Preload("Genres").
		Preload("Serie").
		Preload("Season").
		Find(&p).GetErrors()

	fmt.Println(len(errors))

	for _, err := range errors {
		fmt.Println(err)
	}

	return p
}
