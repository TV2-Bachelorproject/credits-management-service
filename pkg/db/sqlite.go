// +build testing

package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func connect() (*gorm.DB, error) {
	return gorm.Open("sqlite3", "/tmp/tv2-bachelorproject-database.db")
}
