package entity

import (
	"time"
)

type BookEntity struct {
	ID        string    `valid:"required" json:"id"`
	Title     string    `valid:"required" json:"title"`
	CreatedAt time.Time `valid:"required" json:"createdAt"`
	UpdatedAt time.Time `valid:"required" json:"updatedAt"`
	DeletedAt time.Time `valid:"required" json:"deletedAt"`
}
