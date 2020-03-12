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

func Programs() *gorm.DB {
	return db.Model(&Program{})
}
