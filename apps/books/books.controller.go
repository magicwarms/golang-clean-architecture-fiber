package books

import (
	"net/http"

	"startup-backend/config"

	"github.com/gofiber/fiber/v2"
)

// GetAllBooks is to get all books data
func GetAllBooks(bookService BookService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		getAllBooks, _ := bookService.FetchBooks()
		return c.Status(http.StatusOK).JSON(config.AppResponse(http.StatusOK, "OK", getAllBooks))
	}

}

// // GetBook is to get one book data
// func GetBook(c *fiber.Ctx) error {
// 	bookId := c.Query("id")
// 	getBook := repositories.GetBook(bookId)
// 	if getBook.Title == "" {
// 		return c.JSON(config.AppResponse{
// 			Code:    http.StatusOK,
// 			Message: "NOT-FOUND",
// 			Data:    nil,
// 		})
// 	}
// 	return c.JSON(config.AppResponse{
// 		Code:    http.StatusOK,
// 		Message: "OK",
// 		Data:    getBook,
// 	})
// }

// // CreateBook is create new book data
// func CreateBook(c *fiber.Ctx) error {
// 	book := new(models.Books)
// 	if err := c.BodyParser(book); err != nil {
// 		return c.JSON(config.AppResponse{
// 			Code:    http.StatusUnprocessableEntity,
// 			Message: "INVALID-PARAMS",
// 			Data:    nil,
// 		})
// 	}
// 	createBook := repositories.CreateBook(book)
// 	return c.JSON(config.AppResponse{
// 		Code:    http.StatusOK,
// 		Message: "OK",
// 		Data:    createBook,
// 	})
// }

// // DeleteBook is to delete book data
// func DeleteBook(c *fiber.Ctx) error {
// 	book := new(models.Books)
// 	if err := c.BodyParser(book); err != nil {
// 		return c.JSON(config.AppResponse{
// 			Code:    http.StatusUnprocessableEntity,
// 			Message: "INVALID-PARAMS",
// 			Data:    nil,
// 		})
// 	}
// 	getBook := repositories.GetBook(book.ID)
// 	if getBook.Title == "" {
// 		return c.JSON(config.AppResponse{
// 			Code:    http.StatusOK,
// 			Message: "NOT-FOUND",
// 			Data:    nil,
// 		})
// 	}
// 	go repositories.DeleteBook(book)
// 	return c.JSON(config.AppResponse{
// 		Code:    http.StatusOK,
// 		Message: "OK",
// 		Data:    nil,
// 	})
// }

// // UpdateBook is to update book data
// func UpdateBook(c *fiber.Ctx) error {
// 	book := new(models.Books)
// 	if err := c.BodyParser(book); err != nil {
// 		return c.JSON(config.AppResponse{
// 			Code:    http.StatusUnprocessableEntity,
// 			Message: "INVALID-PARAMS",
// 			Data:    nil,
// 		})
// 	}
// 	getBook := repositories.GetBook(book.ID)
// 	if getBook.Title == "" {
// 		return c.JSON(config.AppResponse{
// 			Code:    http.StatusOK,
// 			Message: "NOT-FOUND",
// 			Data:    nil,
// 		})
// 	}
// 	updateBook := repositories.UpdateBook(book)
// 	return c.JSON(config.AppResponse{
// 		Code:    http.StatusOK,
// 		Message: "OK",
// 		Data:    updateBook,
// 	})
// }
