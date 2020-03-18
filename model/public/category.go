package public

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string
}

func (c Category) Find(id uint) Category {
	db.Model(c).Where("id = ?", id).First(&c)
	return c
}

type Categories []Category

func (c Categories) Find() Categories {
	db.Model(c).Find(&c)
	return c
}
