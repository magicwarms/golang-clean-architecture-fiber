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
	GetBookById(id string) (model.BookModel, error)
	SaveBook(book *model.BookModel) error
	UpdateBook(book *model.BookModel) error
	DeleteBook(id string) error
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
	results := bookRepo.db.Order("created_at").Find(&books)
	if results.Error != nil {
		return nil, results.Error
	}
	return &books, nil
}

// GetBookByName is to get only one book data by name
func (bookRepo *bookRepository) GetBookByName(title string) (model.BookModel, error) {
	var book model.BookModel
	result := bookRepo.db.Where("title = ?", title).Take(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.BookModel{}, nil
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

// GetBookById is to get only one book data by ID
func (bookRepo *bookRepository) GetBookById(id string) (model.BookModel, error) {
	var book model.BookModel
	result := bookRepo.db.Where("id = ?", id).Take(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return model.BookModel{}, nil
	}
	if result.Error != nil {
		return model.BookModel{}, result.Error
	}
	return book, nil
}

// DeleteBook is to save book data based on user input
func (bookRepo *bookRepository) DeleteBook(id string) error {
	bookModel := model.BookModel{
		ID: id,
	}
	if err := bookRepo.db.Delete(&bookModel).Error; err != nil {
		return err
	}
	return nil
}

// UpdateBook is to update book data based on user input
func (bookRepo *bookRepository) UpdateBook(book *model.BookModel) error {
	if err := bookRepo.db.Model(&book).Where("id = ?", book.ID).Update("title", book.Title).Error; err != nil {
		return err
	}
	return nil
}
