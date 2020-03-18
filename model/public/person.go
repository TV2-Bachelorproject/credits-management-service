package public

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Name    string
	Credits []Credit `gorm:"many2many:person_credits;"`
}

func (p Person) Find(id uint) Person {
	db.Model(p).Where("id = ?", id).First(&p)
	return p
}

type People []Person

func (p People) Find() People {
	db.Model(&Person{}).Find(&p)
	return p
}
