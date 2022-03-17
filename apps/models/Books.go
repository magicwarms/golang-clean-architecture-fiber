package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Books model
type Books struct {
	ID        string         `gorm:"default:uuid_generate_v4();primaryKey" json:"id"`
	Title     string         `gorm:"index" json:"title"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

//DEFINE HOOKS

func (book *Books) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Before create data", book)
	return
}

func (book *Books) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("After create data", book)
	return
}

func (book *Books) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("Before update data", book)
	return
}

func (book *Books) AfterUpdate(tx *gorm.DB) (err error) {
	fmt.Println("After update data", book)
	return
}

func (book *Books) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("Before delete data", book)
	return
}

func (book *Books) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("After delete data", book)
	return
}
