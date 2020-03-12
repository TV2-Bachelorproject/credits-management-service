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
	Persons     []Person
	CreditGroup CreditGroup
	ProgramID   uint
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
