package infrastructure

import (
	"startup-backend/apps/books"
	"startup-backend/apps/books/handler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Dispatch(DBConnection *gorm.DB, apiV1 fiber.Router) {
	bookRepo := books.NewRepo(DBConnection)
	bookService := books.NewService(bookRepo)
	handler.NewBookHandler(apiV1.Group("/books"), bookService)
}
