package books

import (
	"errors"
	"startup-backend/apps/books/entity"
	"startup-backend/apps/books/model"

	"gorm.io/gorm"
)

// BookRepository interface allows us to access the CRUD Operations in postgresQL here.
type BookRepository interface {
	ListBook() (*[]entity.BookEntity, error)
	GetBookByName(title string) (*entity.BookEntity, error)
	SaveBook(*entity.BookEntity) error
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
func (bookRepo *bookRepository) ListBook() (*[]entity.BookEntity, error) {
	var books []entity.BookEntity
	results := bookRepo.db.Find(&books)
	if results.Error != nil {
		return nil, results.Error
	}
	return &books, nil
}

// GetBookByName is to get only one book data by nmae
func (bookRepo *bookRepository) GetBookByName(title string) (*entity.BookEntity, error) {
	var book entity.BookEntity
	result := bookRepo.db.Where("title = ?", title).First(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &entity.BookEntity{}, nil
	}
	return &book, nil
}

// SaveBook is to save book data based on user input
func (bookRepo *bookRepository) SaveBook(book *entity.BookEntity) error {
	bookModel := model.NewBookModel(book)
	if err := bookRepo.db.Save(bookModel).Error; err != nil {
		return err
	}
	return nil
}
