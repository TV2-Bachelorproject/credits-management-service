package public

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
)

type Genre struct {
	gorm.Model
	Name string
}

func (g Genre) Find(id uint) Genre {
	db.Model(g).Where("id = ?", id).First(&g)
	return g
}

type Genres []Genre

func (g Genres) Find() Genres {
	db.Model(&Genre{}).Find(&g)
	return g
}
