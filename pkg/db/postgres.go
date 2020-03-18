// +build !sqlite

package db

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"

	// import the postgres dialect for gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func connect() (*gorm.DB, error) {
	var (
		err            error
		host, port     = "localhost", int64(5432)
		user, pass, db = "root", "root", "root"
	)

	if value, ok := os.LookupEnv("DB_HOST"); ok {
		host = value
	}

	if value, ok := os.LookupEnv("DB_PORT"); ok {
		port, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			panic(err)
		}
	}

	if value, ok := os.LookupEnv("DB_USER"); ok {
		user = value
	}

	if value, ok := os.LookupEnv("DB_PASSWORD"); ok {
		pass = value
	}

	if value, ok := os.LookupEnv("DB_DATABASE"); ok {
		db = value
	}

	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			host, port, user, db, pass,
		),
	)
}
