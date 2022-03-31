package books

import (
	"startup-backend/apps/books/entity"
)

// BookService is an interface from which our api module can access our repository of all our models
type BookService interface {
	FetchBooks() (*[]entity.BookEntity, error)
	InsertBook(title string) error
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
func (s *bookService) FetchBooks() (*[]entity.BookEntity, error) {
	getAllBooks, err := s.bookRepository.ListBook()
	if err != nil {
		return nil, err
	}
	return getAllBooks, nil
}

// InsertBook is a service layer that helps insert book data to database
func (s *bookService) InsertBook(title string) error {
	// check book data by name to validate
	// _, err := s.bookRepository.GetBookByName(title)
	// // if existed throw an error
	// if err != nil {
	// 	return errors.New("title already exists")
	// }
	// start to insert the data to database through repository
	errSave := s.bookRepository.SaveBook(&entity.BookEntity{Title: title})
	if errSave != nil {
		return errSave
	}
	return nil
}
