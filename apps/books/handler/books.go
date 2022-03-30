package handler

import (
	"net/http"
	"startup-backend/apps/books"
	"startup-backend/config"

	"github.com/gofiber/fiber/v2"
)

func NewUserHandler(book fiber.Router, bookService books.BookService) {
	book.Get("/list", GetAllBooks(bookService))
	book.Post("/store", CreateBook(bookService))
	// book.Get("/:userId", getUser(bookService))
}

// GetAllBooks is to get all books data
func GetAllBooks(bookService books.BookService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		getAllBooks, _ := bookService.FetchBooks()
		return c.Status(http.StatusOK).JSON(config.AppResponse(http.StatusOK, "OK", getAllBooks))
	}
}

// CreateBook is store book data into database
func CreateBook(bookService books.BookService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		book := new(books.BookDTO)
		if err := c.BodyParser(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(config.AppResponse(fiber.StatusBadRequest, "BAD-REQUEST", nil))
		}
		err := bookService.InsertBook(book)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(config.AppResponse(http.StatusInternalServerError, "ERROR", nil))
		}
		return c.Status(http.StatusCreated).JSON(config.AppResponse(http.StatusCreated, "OK", nil))
	}
}
