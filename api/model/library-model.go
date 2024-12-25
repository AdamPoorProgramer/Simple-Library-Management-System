package model

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Title     string
	Author    string
	Publisher string
	Year      uint
	Genre     string
}

func (book Book) TableName() string {
	return "book"
}

type Member struct {
	gorm.Model
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	JoinDate    time.Time
}

func (member Member) TableName() string {
	return "member"
}

type Borrowing struct {
	gorm.Model
	Member     Member
	Book       Book
	BorrowDate time.Time
	ReturnDate time.Time
	Returned   bool
}

func (borrowing Borrowing) TableName() string {
	return "borrowing"
}

type Category struct {
	gorm.Model
	Name string
}

func (category Category) TableName() string {
	return "category"
}

type BookCategory struct {
	gorm.Model
	Book     Book
	Category Category
}

func (bookCategory BookCategory) TableName() string {
	return "book_category"
}
