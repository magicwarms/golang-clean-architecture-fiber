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
		getAllBooks, err := bookService.FetchBooks()
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(config.AppResponse(http.StatusInternalServerError, "ERROR", err))
		}
		return c.Status(http.StatusOK).JSON(config.AppResponse(http.StatusOK, "OK", getAllBooks))
	}
}

// CreateBook is store book data into database
func CreateBook(bookService books.BookService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type BookDTO struct {
			Title string `json:"title"`
		}

		var bookDTO BookDTO
		if err := c.BodyParser(&bookDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(config.AppResponse(fiber.StatusBadRequest, "BAD-REQUEST", err))
		}

		err := bookService.InsertBook(bookDTO.Title)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(config.AppResponse(http.StatusInternalServerError, "ERROR", err))
		}
		return c.Status(http.StatusCreated).JSON(config.AppResponse(http.StatusCreated, "OK", nil))
	}
}
