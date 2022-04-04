package books

import (
	"errors"
	"startup-backend/apps/books/model"
)

// BookService is an interface from which our api module can access our repository of all our models
type BookService interface {
	FindAll() ([]model.BookModel, error)
	Save(book *model.BookModel) error
	// UpdateBook(book *entity.Books) (*entity.Books, error)
	// RemoveBook(ID string) error
}

type bookService struct {
	bookRepository BookRepository
}

// NewService is used to create a single instance of the service
func NewService(r BookRepository) BookService {
	return &bookService{
		bookRepository: r,
	}
}

// FetchBooks is a service layer that helps fetch all books in Book table
func (s *bookService) FindAll() ([]model.BookModel, error) {
	getAllBooks, err := s.bookRepository.ListBook()
	if err != nil {
		return nil, err
	}
	return getAllBooks, nil
}

// InsertBook is a service layer that helps insert book data to database
func (s *bookService) Save(book *model.BookModel) error {
	// check book data by name to validate
	result, errGetBookByName := s.bookRepository.GetBookByName(book.Title)
	if errGetBookByName != nil {
		return errGetBookByName
	}
	// if existed throw an error
	if result.Title != "" {
		return errors.New("title already exists")
	}
	// start to insert the data to database through repository
	errSave := s.bookRepository.SaveBook(book)
	if errSave != nil {
		return errSave
	}
	return nil
}
