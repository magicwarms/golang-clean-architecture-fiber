package books

import (
	"startup-backend/apps/books/model"

	"gorm.io/gorm"
)

// BookRepository interface allows us to access the CRUD Operations in postgresQL here.
type BookRepository interface {
	ListBook() ([]*model.BookModel, error)
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
		db: gormDB.Table("books"),
	}
}

// GetAllBooks is to get all books data
func (bookRepo *bookRepository) ListBook() ([]*model.BookModel, error) {
	var books []*model.BookModel
	results := bookRepo.db.Find(&books)
	if results.Error != nil {
		return nil, results.Error
	}
	return books, nil
}

// GetBookByName is to get only one book data by nmae
func (bookRepo *bookRepository) GetBookByName(title string) (model.BookModel, error) {
	var book model.BookModel
	result := bookRepo.db.Where("title = ?", title).Limit(1).Find(&book)
	if result.Error != nil {
		return model.BookModel{}, result.Error
	}
	if result.RowsAffected > 0 {
		return book, nil
	}
	return model.BookModel{}, nil
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
