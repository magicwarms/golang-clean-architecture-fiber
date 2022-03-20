package books

import (
	"fmt"
	"startup-backend/apps/books/entity"

	"gorm.io/gorm"
)

//BookRepository interface allows us to access the CRUD Operations in postgresQL here.
type BookRepository interface {
	ReadBook() (*[]entity.Books, error)
}

type bookRepository struct {
	table *gorm.DB
}

//NewRepo is the single instance repo that is being created.
func NewRepo(gormDB *gorm.DB) BookRepository {
	// gormDB.AutoMigrate()
	return &bookRepository{
		table: gormDB.Table("books"),
	}
}

// GetAllBooks is to get all books data
func (bookRepo *bookRepository) ReadBook() (*[]entity.Books, error) {
	var books []entity.Books
	results := bookRepo.table.Find(&books)
	if results.Error != nil {
		fmt.Println(results.Error)
		return nil, results.Error
	}
	return &books, nil
}
