package handler

import (
	"net/http"
	"startup-backend/apps/books"
	"startup-backend/apps/books/model"
	"startup-backend/config"

	"github.com/gofiber/fiber/v2"
)

func NewBookHandler(book fiber.Router, bookService books.BookService) {
	book.Get("/list", GetAllBooks(bookService))
	book.Post("/store", AddNewBook(bookService))
	book.Delete("/remove", RemoveBook(bookService))
	book.Get("/:bookId", GetBook(bookService))
	book.Put("/update", UpdateBook(bookService))
}

// GetAllBooks is to get all books data
func GetAllBooks(bookService books.BookService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		getAllBooks, err := bookService.FindAll()
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(config.ErrorResponse(err))
		}
		return c.Status(http.StatusOK).JSON(config.AppResponse(getAllBooks))
	}
}

// AddNewBook is store book data into database
func AddNewBook(bookService books.BookService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var bookDTO model.BookModel
		if err := c.BodyParser(&bookDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(config.ErrorResponse(err))
		}

		err := bookService.Save(&bookDTO)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(config.ErrorResponse(err))
		}
		return c.Status(http.StatusCreated).JSON(config.AppResponse(nil))
	}
}

// RemoveBook is delete book data in database
func RemoveBook(bookService books.BookService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var bookDTO model.BookModel
		if err := c.BodyParser(&bookDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(config.ErrorResponse(err))
		}
		err := bookService.Delete(bookDTO.ID)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(config.ErrorResponse(err))
		}
		return c.Status(http.StatusOK).JSON(config.AppResponse(nil))
	}
}

// GetBook is to get spesific book data by book ID
func GetBook(bookService books.BookService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get bookId params
		var id string = c.Query("bookId")
		if id == "" {
			return c.Status(http.StatusUnprocessableEntity).JSON(config.AppResponse("book ID is mandatory"))
		}
		getBook, err := bookService.Get(id)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(config.ErrorResponse(err))
		}
		if getBook.ID == "" {
			getBook = nil
		}
		return c.Status(http.StatusOK).JSON(config.AppResponse(&getBook))
	}
}

// UpdateBook is update book data into database
func UpdateBook(bookService books.BookService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var bookDTO model.BookModel
		if err := c.BodyParser(&bookDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(config.ErrorResponse(err))
		}

		result, err := bookService.Update(&bookDTO)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(config.ErrorResponse(err))
		}
		return c.Status(http.StatusCreated).JSON(config.AppResponse(&result))
	}
}
