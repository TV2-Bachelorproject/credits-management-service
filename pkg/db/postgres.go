// +build !local

package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func connect() (*gorm.DB, error) {
	return gorm.Open("postgres", "host=127.0.0.1 port=5432 user=root dbname=root password=root sslmode=disable")
}
