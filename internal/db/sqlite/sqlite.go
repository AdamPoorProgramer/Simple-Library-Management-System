package sqlite

import (
	"LIBRARY-API-SERVER/configs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteOrPanic(cnf configs.Sqlite) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(cnf.Path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
