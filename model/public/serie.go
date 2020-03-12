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

func Series() *gorm.DB {
	return db.Model(&Serie{})
}
