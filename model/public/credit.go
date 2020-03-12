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

func Credits() *gorm.DB {
	return db.Model(&Credit{})
}
