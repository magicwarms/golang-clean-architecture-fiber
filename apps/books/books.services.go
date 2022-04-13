package books

import (
	"errors"
	"startup-backend/apps/books/model"
)

// BookService is an interface from which our api module can access our repository of all our models
type BookService interface {
	FindAll() (*[]model.BookModel, error)
	Save(book *model.BookModel) error
	Update(book *model.BookModel) (*model.BookModel, error)
	Delete(ID string) error
	Get(ID string) (*model.BookModel, error)
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

// FindAll is a service layer that helps fetch all books in Book table
func (s *bookService) FindAll() (*[]model.BookModel, error) {
	getAllBooks, err := s.bookRepository.ListBook()
	if err != nil {
		return nil, err
	}
	return getAllBooks, nil
}

// Save is a service layer that helps insert book data to database
func (s *bookService) Save(book *model.BookModel) error {
	// check book data by name to validate
	result, _ := s.bookRepository.GetBookByName(book.Title)
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

// Save is a service layer that helps insert book data to database
func (s *bookService) Delete(id string) error {
	// check book data by name to validate
	result, _ := s.bookRepository.GetBookById(id)
	if result.ID == "" {
		return errors.New("ID not found")
	}
	// start to insert the data to database through repository
	errSave := s.bookRepository.DeleteBook(id)
	if errSave != nil {
		return errSave
	}
	return nil
}

// getBook is a service layer that helps get book data
func (s *bookService) Get(id string) (*model.BookModel, error) {
	getBook, err := s.bookRepository.GetBookById(id)
	if err != nil {
		return &model.BookModel{}, err
	}
	return &getBook, nil
}

// Update is a service layer that helps update book data to database
func (s *bookService) Update(book *model.BookModel) (*model.BookModel, error) {
	// check book data by name to validate
	result, _ := s.bookRepository.GetBookById(book.ID)
	if result.ID == "" {
		return &model.BookModel{}, errors.New("ID not found")
	}
	// start to insert the data to database through repository
	errUpdate := s.bookRepository.UpdateBook(book)
	if errUpdate != nil {
		return &model.BookModel{}, errUpdate
	}
	return book, nil
}
