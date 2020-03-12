package private

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Name  string
	Email string
	Address
}

func People() *gorm.DB {
	return db.Model(&Person{})
}
