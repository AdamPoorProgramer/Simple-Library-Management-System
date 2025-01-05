package handler

import "gorm.io/gorm"

type BookCategoryHandler struct {
	DB *gorm.DB
}

func NewBookCategoryHandler(db *gorm.DB) *BookCategoryHandler {
	return &BookCategoryHandler{DB: db}
}
