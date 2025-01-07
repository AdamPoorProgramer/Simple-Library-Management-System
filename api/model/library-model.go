package model

import (
	"gorm.io/gorm"
	"time"
)

type CustomModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
type Book struct {
	CustomModel
	Title     string     `json:"title"`
	Author    string     `json:"author"`
	Publisher string     `json:"publisher"`
	Year      uint       `json:"year"`
	Genre     string     `json:"genre"`
	Category  []Category `gorm:"many2many:book_category;" json:"categories"`
}

func (book Book) TableName() string {
	return "book"
}

type Member struct {
	CustomModel
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	JoinDate    time.Time `json:"join_date"`
}

func (member Member) TableName() string {
	return "member"
}

type Borrowing struct {
	CustomModel
	BookId     uint       `json:"book_id"`
	MemberId   uint       `json:"member_id"`
	Member     Member     `gorm:"foreignKey:MemberId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member"`
	Book       Book       `gorm:"foreignKey:BookId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"book"`
	BorrowDate time.Time  `json:"borrow_date"`
	ReturnDate *time.Time `json:"return_date,omitempty"`
	Returned   bool       `json:"returned"`
}

func (borrowing Borrowing) TableName() string {
	return "borrowing"
}

type Category struct {
	CustomModel
	Name string `json:"name"`
	Book []Book `gorm:"many2many:book_category;" json:"books"`
}

func (category Category) TableName() string {
	return "category"
}
