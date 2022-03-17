package books

import "startup-backend/apps/books/entity"

//BookService is an interface from which our api module can access our repository of all our models
type BookService interface {
	FetchBooks() (*[]entity.Books, error)
	// InsertBook(book *entities.Book) (*entities.Book, error)
	// UpdateBook(book *entities.Book) (*entities.Book, error)
	// RemoveBook(ID string) error
}

type bookService struct {
	bookRepository BookRepository
}

//NewService is used to create a single instance of the service
func NewService(r BookRepository) BookService {
	return &bookService{
		bookRepository: r,
	}
}

//FetchBooks is a service layer that helps fetch all books in Book table
func (s *bookService) FetchBooks() (*[]entity.Books, error) {
	return s.bookRepository.ReadBook()
}
