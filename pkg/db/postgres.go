// +build !sqlite

package db

import (
	"fmt"

	"github.com/TV2-Bachelorproject/server/pkg/config"
	"github.com/jinzhu/gorm"

	// import the postgres dialect for gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func connect() (*gorm.DB, error) {
	conf := config.Get().DB
	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			conf.Host, conf.Port, conf.User, conf.Database, conf.Password,
		),
	)
}
