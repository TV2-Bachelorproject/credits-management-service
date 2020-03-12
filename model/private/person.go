package private

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Name  string
	Email string
	Address
}

func (s Person) Find(id uint) Person {
	db.Model(s).Where("id = ?", id).First(&s)
	return s
}

type People []Person

func (s People) Find() People {
	db.Model(&Person{}).Find(&s)
	return s
}
