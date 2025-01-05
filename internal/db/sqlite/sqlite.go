package sqlite

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteOrPanic() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
	}
	return db
}
