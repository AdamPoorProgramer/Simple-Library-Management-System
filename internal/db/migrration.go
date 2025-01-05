package db

import (
	"LIBRARY-API-SERVER/api/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) (err error) {
	return db.AutoMigrate(&model.Book{}, &model.Borrowing{}, &model.Member{}, &model.Category{})
}
