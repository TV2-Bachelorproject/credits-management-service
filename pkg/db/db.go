package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	conn          *gorm.DB
	NewRecord     func(interface{}) bool
	Create        func(interface{}) *gorm.DB
	Save          func(interface{}) *gorm.DB
	First         func(interface{}, ...interface{}) *gorm.DB
	Take          func(interface{}, ...interface{}) *gorm.DB
	Last          func(interface{}, ...interface{}) *gorm.DB
	Model         func(interface{}) *gorm.DB
	Delete        func(interface{}, ...interface{}) *gorm.DB
	Where         func(interface{}, ...interface{}) *gorm.DB
	Select        func(interface{}, ...interface{}) *gorm.DB
	Order         func(interface{}, ...bool) *gorm.DB
	Table         func(string) *gorm.DB
	Limit         func(interface{}) *gorm.DB
	Offset        func(interface{}) *gorm.DB
	Not           func(interface{}, ...interface{}) *gorm.DB
	FirstOrInit   func(interface{}, ...interface{}) *gorm.DB
	FirstOrCreate func(interface{}, ...interface{}) *gorm.DB
	Set           func(string, interface{}) *gorm.DB
	Raw           func(string, ...interface{}) *gorm.DB
)

func Setup() error {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=root dbname=root password=root sslmode=disable")

	if err != nil {
		return err
	}

	conn = db
	NewRecord = db.NewRecord
	Create = db.Create
	Save = db.Save
	First = db.First
	Take = db.Take
	Last = db.Last
	Model = db.Model
	Delete = db.Delete
	Where = db.Where
	Select = db.Select
	Order = db.Order
	Table = db.Table
	Limit = db.Limit
	Offset = db.Offset
	Not = db.Not
	FirstOrInit = db.FirstOrInit
	FirstOrCreate = db.FirstOrCreate
	Set = db.Set
	Raw = db.Raw

	return nil
}

func Migrate(tables ...interface{}) {
	if conn == nil {
		return
	}

	conn.AutoMigrate(tables...)
}

func Close() error {
	if conn == nil {
		return nil
	}

	return conn.Close()
}
