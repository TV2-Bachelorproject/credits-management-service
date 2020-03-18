package public

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
)

type Season struct {
	gorm.Model
	Title       string
	RawSeasonID string `gorm:"type:varchar(100)"`
	Programs    []Program
	SerieID     uint
}

func (s Season) Find(id uint) Season {
	db.Model(s).Where("id = ?", id).First(&s)
	return s
}

type Seasons []Season

func (s Seasons) Find() Seasons {
	db.Model(&Season{}).Find(&s)
	return s
}
