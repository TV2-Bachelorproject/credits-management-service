package public

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
)

type CreditGroup struct {
	gorm.Model
	Name string
}

type Credit struct {
	gorm.Model
	Persons       []Person    `gorm:"many2many:credit_persons;"`
	CreditGroupID uint        `json:"-"`
	CreditGroup   CreditGroup `json:"creditGroup"`
	ProgramID     uint        `json:"-"`
	Program       Program     `json:"program" gorm:"foreignkey:ProgramID"`
	SeasonID      uint        `json:"-"`
	Season        Season      `json:"season" gorm:"foreignkey:SeasonID"`
	SerieID       uint        `json:"-"`
	Serie         Serie       `json:"serie" gorm:"foreignkey:SerieID"`
	Accepted      bool        `json:"accepted"`
}

func (s Credit) Find(id uint) Credit {
	db.Model(s).Where("id = ?", id).First(&s)
	return s
}

type Credits []Credit

func (s Credits) Find() Credits {
	db.Model(&Credit{}).Find(&s)
	return s
}
