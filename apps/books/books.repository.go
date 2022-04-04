package books

import (
	"errors"
	"startup-backend/apps/books/model"

	"gorm.io/gorm"
)

// BookRepository interface allows us to access the CRUD Operations in postgresQL here.
type BookRepository interface {
	ListBook() (*[]model.BookModel, error)
	GetBookByName(title string) (model.BookModel, error)
	SaveBook(book *model.BookModel) error
}

type bookRepository struct {
	db *gorm.DB
}

// NewRepo is the single instance repo that is being created.
func NewRepo(gormDB *gorm.DB) BookRepository {
	gormDB.AutoMigrate(&model.BookModel{})
	return &bookRepository{
		db: gormDB,
	}
}

// GetAllBooks is to get all books data
func (bookRepo *bookRepository) ListBook() (*[]model.BookModel, error) {
	var books []model.BookModel
	results := bookRepo.db.Find(&books)
	if results.Error != nil {
		return nil, results.Error
	}
	return &books, nil
}

// GetBookByName is to get only one book data by nmae
func (bookRepo *bookRepository) GetBookByName(title string) (model.BookModel, error) {
	var book model.BookModel
	result := bookRepo.db.Where("title = ?", title).Take(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.BookModel{}, result.Error
	}
	if result.Error != nil {
		return model.BookModel{}, result.Error
	}
	return book, nil
}

// SaveBook is to save book data based on user input
func (bookRepo *bookRepository) SaveBook(book *model.BookModel) error {
	bookModel := model.BookModel{
		Title: book.Title,
	}
	if err := bookRepo.db.Create(&bookModel).Error; err != nil {
		return err
	}
	return nil
}
