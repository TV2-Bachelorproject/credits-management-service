package public

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Name    string
	Credits []Credit
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
