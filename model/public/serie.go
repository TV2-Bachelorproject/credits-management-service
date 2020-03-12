package public

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
)

type Serie struct {
	gorm.Model
	Title   string
	Seasons []Season
}

func (s Serie) Find(id uint) Serie {
	db.Model(s).Where("id = ?", id).First(&s)
	return s
}

type Series []Serie

func (s Series) Find() Series {
	db.Model(&Serie{}).Find(&s)
	return s
}
