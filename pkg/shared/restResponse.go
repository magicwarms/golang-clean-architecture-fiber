package shared

import "github.com/gofiber/fiber/v2"

func NewRestResponse(message string, data interface{}) *fiber.Map {
	return &fiber.Map{
		"message": message,
		"data":    data,
	}
}

func NewRestErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"error": err,
	}
}
