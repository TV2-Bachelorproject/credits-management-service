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

func People() *gorm.DB {
	return db.Model(&Person{})
}
