package books

import (
	"startup-backend/apps/books/entity"
	"startup-backend/apps/books/model"

	"gorm.io/gorm"
)

// BookRepository interface allows us to access the CRUD Operations in postgresQL here.
type BookRepository interface {
	ListBook() (*[]entity.BookEntity, error)
	GetBookByName(bookTitle string) (*entity.BookEntity, error)
	SaveBook(*BookDTO) error
}

type BookDTO struct {
	Title string `json:"title"`
}

type bookRepository struct {
	db *gorm.DB
}

// NewRepo is the single instance repo that is being created.
func NewRepo(gormDB *gorm.DB) BookRepository {
	gormDB.AutoMigrate(&model.Books{})
	return &bookRepository{
		db: gormDB,
	}
}

// GetAllBooks is to get all books data
func (bookRepo *bookRepository) ListBook() (*[]entity.BookEntity, error) {
	var books []entity.BookEntity
	results := bookRepo.db.Find(&books)
	if results.Error != nil {
		return nil, results.Error
	}
	return &books, nil
}

// GetBookByName is to get only one book data by nmae
func (bookRepo *bookRepository) GetBookByName(bookTitle string) (*entity.BookEntity, error) {
	var book entity.BookEntity
	result := bookRepo.db.First(&book, "title = ?", bookTitle)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

// SaveBook is to save book data based on user input
func (bookRepo *bookRepository) SaveBook(book *BookDTO) error {
	bookModel := 
	result := bookRepo.db.Create(bookModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
