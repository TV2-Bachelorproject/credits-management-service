package public

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
)

type Production struct {
	gorm.Model
	Country     string
	Year        int
	ProducedBy  string
	ProducedFor string
	Editor      string
}

func (p Production) Find(id uint) Production {
	db.Model(p).Where("id = ?", id).First(&p)
	return p
}

type Productions []Production

func (p Productions) Find() Productions {
	db.Model(&Production{}).Find(&p)
	return p
}
