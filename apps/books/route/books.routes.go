package route

import (
	"startup-backend/apps/books"

	"github.com/gofiber/fiber/v2"
)

// BookRouter is all routes in Books package
func BookRouter(book fiber.Router, bookService books.BookService) {
	book.Get("books/list", books.GetAllBooks(bookService))
}
