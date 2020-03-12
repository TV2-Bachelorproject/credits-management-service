package public

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
)

type Program struct {
	gorm.Model
	Title    string
	Credits  []Credit
	SeasonID uint
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
