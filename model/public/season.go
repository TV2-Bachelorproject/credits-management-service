package public

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
)

type Season struct {
	gorm.Model
	Title    string
	Programs []Program
	SerieID  uint
}

func Seasons() *gorm.DB {
	return db.Model(&Season{})
}
