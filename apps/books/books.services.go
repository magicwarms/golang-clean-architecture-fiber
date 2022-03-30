package books

import (
	"errors"
	"startup-backend/apps/books/entity"
)

// BookService is an interface from which our api module can access our repository of all our models
type BookService interface {
	FetchBooks() (*[]entity.Books, error)
	InsertBook(book *BookDTO) error
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
func (s *bookService) FetchBooks() (*[]entity.Books, error) {
	return s.bookRepository.ListBook()
}

// InsertBook is a service layer that helps insert book data to database
func (s *bookService) InsertBook(book *BookDTO) error {
	// check book data by name to validate
	_, err := s.bookRepository.GetBookByName(book.Title)
	// if existed throw an error
	if err == nil {
		return errors.New("title already exists")
	}
	// start to insert the data to database through repository
	errSave := s.bookRepository.SaveBook(book)
	if errSave != nil {
		return errSave
	}
	return nil
}
