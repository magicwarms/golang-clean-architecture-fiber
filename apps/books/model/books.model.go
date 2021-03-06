package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Book Constructs your Book model under entities.
type BookModel struct {
	ID        string         `gorm:"default:uuid_generate_v4();primaryKey" json:"id"`
	Title     string         `gorm:"index" json:"title"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Set tablename (GORM)
func (BookModel) TableName() string {
	return "books"
}

//DEFINE HOOKS
func (book *BookModel) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Before create data", book)
	return
}

func (book *BookModel) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("After create data", book)
	return
}

func (book *BookModel) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("Before update data", book)
	return
}

func (book *BookModel) AfterUpdate(tx *gorm.DB) (err error) {
	fmt.Println("After update data", book)
	return
}

func (book *BookModel) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("Before delete data", book)
	return
}

func (book *BookModel) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("After delete data", book)
	return
}
